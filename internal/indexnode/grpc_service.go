package indexnode

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/zilliztech/milvus-distributed/internal/proto/commonpb"
	"github.com/zilliztech/milvus-distributed/internal/proto/indexpb"
)

const (
	reqTimeoutInterval = time.Second * 10
)

func (b *Builder) BuildIndex(ctx context.Context, request *indexpb.BuildIndexRequest) (*indexpb.BuildIndexResponse, error) {
	t := NewIndexAddTask()
	t.req = request
	t.idAllocator = b.idAllocator
	t.buildQueue = b.sched.IndexBuildQueue
	t.table = b.metaTable
	t.kv = b.kv
	var cancel func()
	t.ctx, cancel = context.WithTimeout(ctx, reqTimeoutInterval)
	defer cancel()

	fn := func() error {
		select {
		case <-ctx.Done():
			return errors.New("insert timeout")
		default:
			return b.sched.IndexAddQueue.Enqueue(t)
		}
	}
	ret := &indexpb.BuildIndexResponse{
		Status: &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_SUCCESS,
		},
	}

	err := fn()
	if err != nil {
		ret.Status.ErrorCode = commonpb.ErrorCode_UNEXPECTED_ERROR
		ret.Status.Reason = err.Error()
		return ret, nil
	}

	err = t.WaitToFinish()
	if err != nil {
		ret.Status.ErrorCode = commonpb.ErrorCode_UNEXPECTED_ERROR
		ret.Status.Reason = err.Error()
		return ret, nil
	}
	ret.IndexID = t.indexID
	return ret, nil
}

func (b *Builder) GetIndexStates(ctx context.Context, request *indexpb.IndexStatesRequest) (*indexpb.IndexStatesResponse, error) {
	var indexStates []*indexpb.IndexInfo
	for _, indexID := range request.IndexID {
		indexState, err := b.metaTable.GetIndexStates(indexID)
		if err != nil {
			log.Println("GetIndexStates error, err=", err)
		}
		indexStates = append(indexStates, indexState)
	}
	ret := &indexpb.IndexStatesResponse{
		Status: &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_SUCCESS,
			Reason:    "",
		},
		States: indexStates,
	}

	return ret, nil
}

func (b *Builder) GetIndexFilePaths(ctx context.Context, request *indexpb.IndexFilePathsRequest) (*indexpb.IndexFilePathsResponse, error) {
	ret := &indexpb.IndexFilePathsResponse{
		Status: &commonpb.Status{ErrorCode: commonpb.ErrorCode_SUCCESS},
	}
	var filePathInfos []*indexpb.IndexFilePathInfo
	for _, indexID := range request.IndexIDs {
		filePathInfo := &indexpb.IndexFilePathInfo{
			Status:  &commonpb.Status{ErrorCode: commonpb.ErrorCode_SUCCESS},
			IndexID: indexID,
		}

		filePaths, err := b.metaTable.GetIndexFilePaths(indexID)
		if err != nil {
			filePathInfo.Status.ErrorCode = commonpb.ErrorCode_UNEXPECTED_ERROR
			filePathInfo.Status.Reason = err.Error()
		}
		filePathInfo.IndexFilePaths = filePaths
		filePathInfos = append(filePathInfos, filePathInfo)
	}
	ret.FilePaths = filePathInfos
	return ret, nil
}