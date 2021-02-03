package masterservice

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	ms "github.com/zilliztech/milvus-distributed/internal/msgstream"
	"github.com/zilliztech/milvus-distributed/internal/msgstream/pulsarms"
	msutil "github.com/zilliztech/milvus-distributed/internal/msgstream/util"
	"github.com/zilliztech/milvus-distributed/internal/proto/commonpb"
	"github.com/zilliztech/milvus-distributed/internal/proto/datapb"
	"github.com/zilliztech/milvus-distributed/internal/proto/indexpb"
	"github.com/zilliztech/milvus-distributed/internal/proto/internalpb2"
	"github.com/zilliztech/milvus-distributed/internal/proto/milvuspb"
	"github.com/zilliztech/milvus-distributed/internal/proto/proxypb"
	"github.com/zilliztech/milvus-distributed/internal/proto/schemapb"
	"github.com/zilliztech/milvus-distributed/internal/util/typeutil"
)

type proxyMock struct {
	randVal   int
	collArray []string
	mutex     sync.Mutex
}

func (p *proxyMock) GetTimeTickChannel() (string, error) {
	return fmt.Sprintf("proxy-time-tick-%d", p.randVal), nil
}
func (p *proxyMock) InvalidateCollectionMetaCache(request *proxypb.InvalidateCollMetaCacheRequest) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.collArray = append(p.collArray, request.CollectionName)
	return nil
}
func (p *proxyMock) GetCollArray() []string {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	ret := make([]string, 0, len(p.collArray))
	ret = append(ret, p.collArray...)
	return ret
}

type dataMock struct {
	randVal int
}

func (d *dataMock) GetInsertBinlogPaths(req *datapb.InsertBinlogPathRequest) (*datapb.InsertBinlogPathsResponse, error) {
	rst := &datapb.InsertBinlogPathsResponse{
		FieldIDs: []int64{},
		Paths:    []*internalpb2.StringList{},
		Status: &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_SUCCESS,
			Reason:    "",
		},
	}
	for i := 0; i < 200; i++ {
		rst.FieldIDs = append(rst.FieldIDs, int64(i))
		path := &internalpb2.StringList{
			Values: []string{fmt.Sprintf("file0-%d", i), fmt.Sprintf("file1-%d", i), fmt.Sprintf("file2-%d", i)},
			Status: &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_SUCCESS,
				Reason:    "",
			},
		}
		rst.Paths = append(rst.Paths, path)
	}
	return rst, nil
}

func (d *dataMock) GetSegmentInfoChannel() (string, error) {
	return fmt.Sprintf("segment-info-channel-%d", d.randVal), nil
}

type indexMock struct {
	fileArray []string
	idxArray  []int64
	mutex     sync.Mutex
}

func (idx *indexMock) BuildIndex(req *indexpb.BuildIndexRequest) (*indexpb.BuildIndexResponse, error) {
	idx.mutex.Lock()
	defer idx.mutex.Unlock()
	idx.fileArray = append(idx.fileArray, req.DataPaths...)
	idx.idxArray = append(idx.idxArray, rand.Int63())
	return &indexpb.BuildIndexResponse{
		Status: &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_SUCCESS,
			Reason:    "",
		},
		IndexBuildID: idx.idxArray[len(idx.idxArray)-1],
	}, nil
}

func (idx *indexMock) getFileArray() []string {
	idx.mutex.Lock()
	defer idx.mutex.Unlock()

	ret := make([]string, 0, len(idx.fileArray))
	ret = append(ret, idx.fileArray...)
	return ret
}

func TestMasterService(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	core, err := NewCore(ctx)
	assert.Nil(t, err)
	randVal := rand.Int()

	Params.TimeTickChannel = fmt.Sprintf("master-time-tick-%d", randVal)
	Params.DdChannel = fmt.Sprintf("master-dd-%d", randVal)
	Params.StatisticsChannel = fmt.Sprintf("master-statistics-%d", randVal)
	Params.MetaRootPath = fmt.Sprintf("/%d/%s", randVal, Params.MetaRootPath)
	Params.KvRootPath = fmt.Sprintf("/%d/%s", randVal, Params.KvRootPath)
	Params.MsgChannelSubName = fmt.Sprintf("subname-%d", randVal)

	pm := &proxyMock{
		randVal:   randVal,
		collArray: make([]string, 0, 16),
		mutex:     sync.Mutex{},
	}
	err = core.SetProxyService(pm)
	assert.Nil(t, err)

	dm := &dataMock{randVal: randVal}
	err = core.SetDataService(dm)
	assert.Nil(t, err)

	im := &indexMock{
		fileArray: []string{},
		idxArray:  []int64{},
		mutex:     sync.Mutex{},
	}
	err = core.SetIndexService(im)
	assert.Nil(t, err)

	err = core.Init()
	assert.Nil(t, err)

	err = core.Start()
	assert.Nil(t, err)

	proxyTimeTickStream := pulsarms.NewPulsarMsgStream(ctx, 1024)
	proxyTimeTickStream.SetPulsarClient(Params.PulsarAddress)
	proxyTimeTickStream.CreatePulsarProducers([]string{Params.ProxyTimeTickChannel})

	dataServiceSegmentStream := pulsarms.NewPulsarMsgStream(ctx, 1024)
	dataServiceSegmentStream.SetPulsarClient(Params.PulsarAddress)
	dataServiceSegmentStream.CreatePulsarProducers([]string{Params.DataServiceSegmentChannel})

	timeTickStream := pulsarms.NewPulsarMsgStream(ctx, 1024)
	timeTickStream.SetPulsarClient(Params.PulsarAddress)
	timeTickStream.CreatePulsarConsumers([]string{Params.TimeTickChannel}, Params.MsgChannelSubName, msutil.NewUnmarshalDispatcher(), 1024)
	timeTickStream.Start()

	ddStream := pulsarms.NewPulsarMsgStream(ctx, 1024)
	ddStream.SetPulsarClient(Params.PulsarAddress)
	ddStream.CreatePulsarConsumers([]string{Params.DdChannel}, Params.MsgChannelSubName, msutil.NewUnmarshalDispatcher(), 1024)
	ddStream.Start()

	time.Sleep(time.Second)

	t.Run("time tick", func(t *testing.T) {
		var timeTick typeutil.Timestamp = 100
		msgPack := ms.MsgPack{}
		baseMsg := ms.BaseMsg{
			MsgCtx:         nil,
			BeginTimestamp: timeTick,
			EndTimestamp:   timeTick,
			HashValues:     []uint32{0},
		}
		timeTickResult := internalpb2.TimeTickMsg{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_kTimeTick,
				MsgID:     0,
				Timestamp: timeTick,
				SourceID:  0,
			},
		}
		timeTickMsg := &ms.TimeTickMsg{
			BaseMsg:     baseMsg,
			TimeTickMsg: timeTickResult,
		}
		msgPack.Msgs = append(msgPack.Msgs, timeTickMsg)
		err := proxyTimeTickStream.Broadcast(&msgPack)
		assert.Nil(t, err)

		ttmsg, ok := <-timeTickStream.Chan()
		assert.True(t, ok)
		assert.Equal(t, len(ttmsg.Msgs), 1)
		ttm, ok := (ttmsg.Msgs[0]).(*ms.TimeTickMsg)
		assert.True(t, ok)
		assert.Equal(t, ttm.Base.Timestamp, timeTick)

		ddmsg, ok := <-ddStream.Chan()
		assert.True(t, ok)
		assert.Equal(t, len(ddmsg.Msgs), 1)
		ddm, ok := (ddmsg.Msgs[0]).(*ms.TimeTickMsg)
		assert.True(t, ok)
		assert.Equal(t, ddm.Base.Timestamp, timeTick)
	})

	t.Run("create collection", func(t *testing.T) {
		schema := schemapb.CollectionSchema{
			Name:        "testColl",
			Description: "testColl",
			AutoID:      true,
			Fields: []*schemapb.FieldSchema{
				{
					FieldID:      100,
					Name:         "vector",
					IsPrimaryKey: false,
					Description:  "vector",
					DataType:     schemapb.DataType_VECTOR_FLOAT,
					TypeParams:   nil,
					IndexParams: []*commonpb.KeyValuePair{
						{
							Key:   "ik1",
							Value: "iv1",
						},
					},
				},
			},
		}
		sbf, err := proto.Marshal(&schema)
		assert.Nil(t, err)

		req := &milvuspb.CreateCollectionRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_kCreateCollection,
				MsgID:     100,
				Timestamp: 100,
				SourceID:  100,
			},
			DbName:         "testDb",
			CollectionName: "testColl",
			Schema:         sbf,
		}
		status, err := core.CreateCollection(req)
		assert.Nil(t, err)
		assert.Equal(t, status.ErrorCode, commonpb.ErrorCode_SUCCESS)

		msg, ok := <-ddStream.Chan()
		assert.True(t, ok)
		assert.Equal(t, len(msg.Msgs), 1)
		createMsg, ok := (msg.Msgs[0]).(*ms.CreateCollectionMsg)
		assert.True(t, ok)
		createMeta, err := core.MetaTable.GetCollectionByName("testColl")
		assert.Nil(t, err)
		assert.Equal(t, createMsg.CollectionID, createMeta.ID)

		req.Base.MsgID = 101
		req.Base.Timestamp = 101
		req.Base.SourceID = 101
		status, err = core.CreateCollection(req)
		assert.Nil(t, err)
		assert.Equal(t, status.ErrorCode, commonpb.ErrorCode_UNEXPECTED_ERROR)

		req.Base.MsgID = 102
		req.Base.Timestamp = 102
		req.Base.SourceID = 102
		req.CollectionName = "testColl-again"
		status, err = core.CreateCollection(req)
		assert.Nil(t, err)
		assert.Equal(t, status.ErrorCode, commonpb.ErrorCode_UNEXPECTED_ERROR)

		schema.Name = req.CollectionName
		sbf, err = proto.Marshal(&schema)
		assert.Nil(t, err)
		req.Schema = sbf
		req.Base.MsgID = 103
		req.Base.Timestamp = 103
		req.Base.SourceID = 103
		status, err = core.CreateCollection(req)
		assert.Nil(t, err)
		assert.Equal(t, status.ErrorCode, commonpb.ErrorCode_SUCCESS)

		msg, ok = <-ddStream.Chan()
		assert.True(t, ok)
		createMsg, ok = (msg.Msgs[0]).(*ms.CreateCollectionMsg)
		assert.True(t, ok)
		createMeta, err = core.MetaTable.GetCollectionByName("testColl-again")
		assert.Nil(t, err)
		assert.Equal(t, createMsg.CollectionID, createMeta.ID)
	})

	t.Run("has collection", func(t *testing.T) {
		req := &milvuspb.HasCollectionRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_kHasCollection,
				MsgID:     110,
				Timestamp: 110,
				SourceID:  110,
			},
			DbName:         "testDb",
			CollectionName: "testColl",
		}
		rsp, err := core.HasCollection(req)
		assert.Nil(t, err)
		assert.Equal(t, rsp.Status.ErrorCode, commonpb.ErrorCode_SUCCESS)
		assert.Equal(t, rsp.Value, true)

		req = &milvuspb.HasCollectionRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_kHasCollection,
				MsgID:     111,
				Timestamp: 111,
				SourceID:  111,
			},
			DbName:         "testDb",
			CollectionName: "testColl2",
		}
		rsp, err = core.HasCollection(req)
		assert.Nil(t, err)
		assert.Equal(t, rsp.Status.ErrorCode, commonpb.ErrorCode_SUCCESS)
		assert.Equal(t, rsp.Value, false)

		// test time stamp go back
		req = &milvuspb.HasCollectionRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_kHasCollection,
				MsgID:     111,
				Timestamp: 111,
				SourceID:  111,
			},
			DbName:         "testDb",
			CollectionName: "testColl",
		}
		rsp, err = core.HasCollection(req)
		assert.Nil(t, err)
		assert.Equal(t, rsp.Status.ErrorCode, commonpb.ErrorCode_SUCCESS)
		assert.Equal(t, rsp.Value, true)
	})

	t.Run("describe collection", func(t *testing.T) {
		collMeta, err := core.MetaTable.GetCollectionByName("testColl")
		assert.Nil(t, err)
		req := &milvuspb.DescribeCollectionRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_kDescribeCollection,
				MsgID:     120,
				Timestamp: 120,
				SourceID:  120,
			},
			DbName:         "testDb",
			CollectionName: "testColl",
		}
		rsp, err := core.DescribeCollection(req)
		assert.Nil(t, err)
		assert.Equal(t, rsp.Status.ErrorCode, commonpb.ErrorCode_SUCCESS)
		assert.Equal(t, rsp.Schema.Name, "testColl")
		assert.Equal(t, rsp.CollectionID, collMeta.ID)

	})

	t.Run("show collection", func(t *testing.T) {
		req := &milvuspb.ShowCollectionRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_kShowCollections,
				MsgID:     130,
				Timestamp: 130,
				SourceID:  130,
			},
			DbName: "testDb",
		}
		rsp, err := core.ShowCollections(req)
		assert.Nil(t, err)
		assert.Equal(t, rsp.Status.ErrorCode, commonpb.ErrorCode_SUCCESS)
		assert.ElementsMatch(t, rsp.CollectionNames, []string{"testColl", "testColl-again"})
		assert.Equal(t, len(rsp.CollectionNames), 2)
	})

	t.Run("create partition", func(t *testing.T) {
		req := &milvuspb.CreatePartitionRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_kCreatePartition,
				MsgID:     140,
				Timestamp: 140,
				SourceID:  140,
			},
			DbName:         "testDb",
			CollectionName: "testColl",
			PartitionName:  "testPartition",
		}
		status, err := core.CreatePartition(req)
		assert.Nil(t, err)
		assert.Equal(t, status.ErrorCode, commonpb.ErrorCode_SUCCESS)
		collMeta, err := core.MetaTable.GetCollectionByName("testColl")
		assert.Nil(t, err)
		assert.Equal(t, len(collMeta.PartitionIDs), 2)
		partMeta, err := core.MetaTable.GetPartitionByID(collMeta.PartitionIDs[1])
		assert.Nil(t, err)
		assert.Equal(t, partMeta.PartitionName, "testPartition")

		msg, ok := <-ddStream.Chan()
		assert.True(t, ok)
		assert.Equal(t, len(msg.Msgs), 1)
		partMsg, ok := (msg.Msgs[0]).(*ms.CreatePartitionMsg)
		assert.True(t, ok)
		assert.Equal(t, partMsg.CollectionID, collMeta.ID)
		assert.Equal(t, partMsg.PartitionID, partMeta.PartitionID)
	})

	t.Run("has partition", func(t *testing.T) {
		req := &milvuspb.HasPartitionRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_kHasPartition,
				MsgID:     150,
				Timestamp: 150,
				SourceID:  150,
			},
			DbName:         "testDb",
			CollectionName: "testColl",
			PartitionName:  "testPartition",
		}
		rsp, err := core.HasPartition(req)
		assert.Nil(t, err)
		assert.Equal(t, rsp.Status.ErrorCode, commonpb.ErrorCode_SUCCESS)
		assert.Equal(t, rsp.Value, true)
	})

	t.Run("show partition", func(t *testing.T) {
		coll, err := core.MetaTable.GetCollectionByName("testColl")
		assert.Nil(t, err)
		req := &milvuspb.ShowPartitionRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_kShowPartitions,
				MsgID:     160,
				Timestamp: 160,
				SourceID:  160,
			},
			DbName:         "testDb",
			CollectionName: "testColl",
			CollectionID:   coll.ID,
		}
		rsp, err := core.ShowPartitions(req)
		assert.Nil(t, err)
		assert.Equal(t, rsp.Status.ErrorCode, commonpb.ErrorCode_SUCCESS)
		assert.Equal(t, len(rsp.PartitionNames), 2)
		assert.Equal(t, len(rsp.PartitionIDs), 2)
	})

	t.Run("show segment", func(t *testing.T) {
		coll, err := core.MetaTable.GetCollectionByName("testColl")
		assert.Nil(t, err)
		partID := coll.PartitionIDs[1]
		part, err := core.MetaTable.GetPartitionByID(partID)
		assert.Nil(t, err)
		assert.Zero(t, len(part.SegmentIDs))

		seg := &datapb.SegmentInfo{
			SegmentID:    1000,
			CollectionID: coll.ID,
			PartitionID:  part.PartitionID,
		}

		msgPack := ms.MsgPack{}
		baseMsg := ms.BaseMsg{
			BeginTimestamp: 0,
			EndTimestamp:   0,
			HashValues:     []uint32{0},
		}
		segMsg := &ms.SegmentInfoMsg{
			BaseMsg: baseMsg,
			SegmentMsg: datapb.SegmentMsg{
				Base: &commonpb.MsgBase{
					MsgType:   commonpb.MsgType_kSegmentInfo,
					MsgID:     0,
					Timestamp: 0,
					SourceID:  0,
				},
				Segment: seg,
			},
		}
		msgPack.Msgs = append(msgPack.Msgs, segMsg)
		err = dataServiceSegmentStream.Broadcast(&msgPack)
		assert.Nil(t, err)
		time.Sleep(time.Second)

		part, err = core.MetaTable.GetPartitionByID(partID)
		assert.Nil(t, err)
		assert.Equal(t, len(part.SegmentIDs), 1)

		req := &milvuspb.ShowSegmentRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_kShowSegment,
				MsgID:     170,
				Timestamp: 170,
				SourceID:  170,
			},
			CollectionID: coll.ID,
			PartitionID:  partID,
		}
		rsp, err := core.ShowSegments(req)
		assert.Nil(t, err)
		assert.Equal(t, rsp.Status.ErrorCode, commonpb.ErrorCode_SUCCESS)
		assert.Equal(t, rsp.SegmentIDs[0], int64(1000))
		assert.Equal(t, len(rsp.SegmentIDs), 1)
	})

	t.Run("create index", func(t *testing.T) {
		req := &milvuspb.CreateIndexRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_kCreateIndex,
				MsgID:     180,
				Timestamp: 180,
				SourceID:  180,
			},
			DbName:         "",
			CollectionName: "testColl",
			FieldName:      "vector",
			ExtraParams: []*commonpb.KeyValuePair{
				{
					Key:   "ik2",
					Value: "iv2",
				},
			},
		}
		collMeta, err := core.MetaTable.GetCollectionByName("testColl")
		assert.Nil(t, err)
		assert.Equal(t, len(collMeta.IndexParams), 1)

		rsp, err := core.CreateIndex(req)
		assert.Nil(t, err)
		assert.Equal(t, rsp.ErrorCode, commonpb.ErrorCode_SUCCESS)
		time.Sleep(time.Second)
		files := im.getFileArray()
		assert.Equal(t, 3, len(files))
		assert.ElementsMatch(t, files, []string{"file0-100", "file1-100", "file2-100"})
		collMeta, err = core.MetaTable.GetCollectionByName("testColl")
		assert.Nil(t, err)
		assert.Equal(t, len(collMeta.IndexParams), 2)

		req.FieldName = "no field"
		rsp, err = core.CreateIndex(req)
		assert.Nil(t, err)
		assert.NotEqual(t, rsp.ErrorCode, commonpb.ErrorCode_SUCCESS)
	})

	t.Run("describe segment", func(t *testing.T) {
		coll, err := core.MetaTable.GetCollectionByName("testColl")
		assert.Nil(t, err)

		req := &milvuspb.DescribeSegmentRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_kDescribeSegment,
				MsgID:     190,
				Timestamp: 190,
				SourceID:  190,
			},
			CollectionID: coll.ID,
			SegmentID:    1000,
		}
		rsp, err := core.DescribeSegment(req)
		assert.Nil(t, err)
		assert.Equal(t, rsp.Status.ErrorCode, commonpb.ErrorCode_SUCCESS)
		t.Logf("index id = %d", rsp.IndexID)
	})

	t.Run("describe index", func(t *testing.T) {
		req := &milvuspb.DescribeIndexRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_kDescribeIndex,
				MsgID:     200,
				Timestamp: 200,
				SourceID:  200,
			},
			DbName:         "",
			CollectionName: "testColl",
			FieldName:      "vector",
			IndexName:      "",
		}
		rsp, err := core.DescribeIndex(req)
		assert.Nil(t, err)
		assert.Equal(t, rsp.Status.ErrorCode, commonpb.ErrorCode_SUCCESS)
		assert.Equal(t, len(rsp.IndexDescriptions), 1)
		assert.Equal(t, rsp.IndexDescriptions[0].IndexName, Params.DefaultIndexName)
	})

	t.Run("flush segment", func(t *testing.T) {
		coll, err := core.MetaTable.GetCollectionByName("testColl")
		assert.Nil(t, err)
		partID := coll.PartitionIDs[1]
		part, err := core.MetaTable.GetPartitionByID(partID)
		assert.Nil(t, err)
		assert.Equal(t, len(part.SegmentIDs), 1)

		seg := &datapb.SegmentInfo{
			SegmentID:    1001,
			CollectionID: coll.ID,
			PartitionID:  part.PartitionID,
		}

		msgPack := ms.MsgPack{}
		baseMsg := ms.BaseMsg{
			BeginTimestamp: 0,
			EndTimestamp:   0,
			HashValues:     []uint32{0},
		}
		segMsg := &ms.SegmentInfoMsg{
			BaseMsg: baseMsg,
			SegmentMsg: datapb.SegmentMsg{
				Base: &commonpb.MsgBase{
					MsgType:   commonpb.MsgType_kSegmentInfo,
					MsgID:     0,
					Timestamp: 0,
					SourceID:  0,
				},
				Segment: seg,
			},
		}
		msgPack.Msgs = append(msgPack.Msgs, segMsg)
		err = dataServiceSegmentStream.Broadcast(&msgPack)
		assert.Nil(t, err)
		time.Sleep(time.Second)

		part, err = core.MetaTable.GetPartitionByID(partID)
		assert.Nil(t, err)
		assert.Equal(t, len(part.SegmentIDs), 2)

		flushMsg := &ms.FlushCompletedMsg{
			BaseMsg: baseMsg,
			SegmentFlushCompletedMsg: internalpb2.SegmentFlushCompletedMsg{
				Base: &commonpb.MsgBase{
					MsgType:   commonpb.MsgType_kSegmentFlushDone,
					MsgID:     0,
					Timestamp: 0,
					SourceID:  0,
				},
				SegmentID: 1001,
			},
		}
		msgPack.Msgs = []ms.TsMsg{flushMsg}
		err = dataServiceSegmentStream.Broadcast(&msgPack)
		assert.Nil(t, err)
		time.Sleep(time.Second)

		req := &milvuspb.DescribeIndexRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_kDescribeIndex,
				MsgID:     210,
				Timestamp: 210,
				SourceID:  210,
			},
			DbName:         "",
			CollectionName: "testColl",
			FieldName:      "vector",
			IndexName:      "",
		}
		rsp, err := core.DescribeIndex(req)
		assert.Nil(t, err)
		assert.Equal(t, rsp.Status.ErrorCode, commonpb.ErrorCode_SUCCESS)
		assert.Equal(t, len(rsp.IndexDescriptions), 3)
		assert.Equal(t, rsp.IndexDescriptions[0].IndexName, Params.DefaultIndexName)
		assert.Equal(t, rsp.IndexDescriptions[1].IndexName, "index_field_100_0")
		assert.Equal(t, rsp.IndexDescriptions[2].IndexName, "index_field_100_1")
	})

	t.Run("drop partition", func(t *testing.T) {
		req := &milvuspb.DropPartitionRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_kDropPartition,
				MsgID:     220,
				Timestamp: 220,
				SourceID:  220,
			},
			DbName:         "testDb",
			CollectionName: "testColl",
			PartitionName:  "testPartition",
		}
		collMeta, err := core.MetaTable.GetCollectionByName("testColl")
		assert.Nil(t, err)
		dropPartID := collMeta.PartitionIDs[1]
		status, err := core.DropPartition(req)
		assert.Nil(t, err)
		assert.Equal(t, status.ErrorCode, commonpb.ErrorCode_SUCCESS)
		collMeta, err = core.MetaTable.GetCollectionByName("testColl")
		assert.Nil(t, err)
		assert.Equal(t, len(collMeta.PartitionIDs), 1)
		partMeta, err := core.MetaTable.GetPartitionByID(collMeta.PartitionIDs[0])
		assert.Nil(t, err)
		assert.Equal(t, partMeta.PartitionName, Params.DefaultPartitionName)

		msg, ok := <-ddStream.Chan()
		assert.True(t, ok)
		assert.Equal(t, len(msg.Msgs), 1)
		dmsg, ok := (msg.Msgs[0]).(*ms.DropPartitionMsg)
		assert.True(t, ok)
		assert.Equal(t, dmsg.CollectionID, collMeta.ID)
		assert.Equal(t, dmsg.PartitionID, dropPartID)
	})

	t.Run("drop collection", func(t *testing.T) {
		req := &milvuspb.DropCollectionRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_kDropCollection,
				MsgID:     230,
				Timestamp: 230,
				SourceID:  230,
			},
			DbName:         "testDb",
			CollectionName: "testColl",
		}
		collMeta, err := core.MetaTable.GetCollectionByName("testColl")
		assert.Nil(t, err)
		status, err := core.DropCollection(req)
		assert.Nil(t, err)
		assert.Equal(t, status.ErrorCode, commonpb.ErrorCode_SUCCESS)

		msg, ok := <-ddStream.Chan()
		assert.True(t, ok)
		assert.Equal(t, len(msg.Msgs), 1)
		dmsg, ok := (msg.Msgs[0]).(*ms.DropCollectionMsg)
		assert.True(t, ok)
		assert.Equal(t, dmsg.CollectionID, collMeta.ID)
		collArray := pm.GetCollArray()
		assert.Equal(t, len(collArray), 1)
		assert.Equal(t, collArray[0], "testColl")

		req = &milvuspb.DropCollectionRequest{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_kDropCollection,
				MsgID:     231,
				Timestamp: 231,
				SourceID:  231,
			},
			DbName:         "testDb",
			CollectionName: "testColl",
		}
		status, err = core.DropCollection(req)
		assert.Nil(t, err)
		assert.Equal(t, status.ErrorCode, commonpb.ErrorCode_UNEXPECTED_ERROR)
		time.Sleep(time.Second)
		assert.Zero(t, len(ddStream.Chan()))
		collArray = pm.GetCollArray()
		assert.Equal(t, len(collArray), 1)
		assert.Equal(t, collArray[0], "testColl")
	})

	err = core.Stop()
	assert.Nil(t, err)

}