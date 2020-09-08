package com;

import com.alibaba.fastjson.JSONObject;
import io.milvus.client.*;
import org.testng.Assert;
import org.testng.annotations.*;

import java.util.List;
import java.util.Map;

public class TestCollection {
    int segmentRowCount = 5000;
    int dimension = 128;

    // case-01
    @Test(dataProvider = "ConnectInstance", dataProviderClass = MainClass.class)
    public void testCreateCollection(MilvusClient client, String collectionName) {
        CollectionMapping collectionSchema = new CollectionMapping.Builder(collectionName)
                .withFields(Utils.genDefaultFields(dimension,false))
                .withParamsInJson(String.format("{\"segment_row_count\": %s}",segmentRowCount))
                .build();
        Response res = client.createCollection(collectionSchema);
        assert(res.ok());
        Assert.assertEquals(res.ok(), true);
    }

    // case-02
    @Test(dataProvider = "DisConnectInstance", dataProviderClass = MainClass.class)
    public void testCreateCollectionDisconnect(MilvusClient client, String collectionName) {
        CollectionMapping collectionSchema = new CollectionMapping.Builder(collectionName)
                .withFields(Utils.genDefaultFields(dimension,false))
                .withParamsInJson(String.format("{\"segment_row_count\": %s}",segmentRowCount))
                .build();
        Response res = client.createCollection(collectionSchema);
        assert(!res.ok());
    }

    // case-03
    @Test(dataProvider = "ConnectInstance", dataProviderClass = MainClass.class)
    public void testCreateCollectionRepeatably(MilvusClient client, String collectionName) {
        CollectionMapping collectionSchema = new CollectionMapping.Builder(collectionName)
                .withFields(Utils.genDefaultFields(dimension,false))
                .withParamsInJson(String.format("{\"segment_row_count\": %s}",segmentRowCount))
                .build();
        Response res = client.createCollection(collectionSchema);
        Assert.assertEquals(res.ok(), true);
        Response resNew = client.createCollection(collectionSchema);
        Assert.assertEquals(resNew.ok(), false);
    }

    // case-04
    @Test(dataProvider = "ConnectInstance", dataProviderClass = MainClass.class)
    public void testCreateCollectionWrongParams(MilvusClient client, String collectionName) {
        Integer dim = 0;
        CollectionMapping collectionSchema = new CollectionMapping.Builder(collectionName)
                .withFields(Utils.genDefaultFields(dim,false))
                .withParamsInJson(String.format("{\"segment_row_count\": %s}",segmentRowCount))
                .build();
        Response res = client.createCollection(collectionSchema);
        System.out.println(res.toString());
        Assert.assertEquals(res.ok(), false);
    }

    // case-05
    @Test(dataProvider = "ConnectInstance", dataProviderClass = MainClass.class)
    public void testShowCollections(MilvusClient client, String collectionName) {
        Integer collectionNum = 10;
        ListCollectionsResponse res = null;
        for (int i = 0; i < collectionNum; ++i) {
            String collectionNameNew = collectionName+"_"+Integer.toString(i);
            CollectionMapping collectionSchema = new CollectionMapping.Builder(collectionNameNew)
                    .withFields(Utils.genDefaultFields(dimension,false))
                    .withParamsInJson(String.format("{\"segment_row_count\": %s}",segmentRowCount))
                    .build();
            client.createCollection(collectionSchema);
            List<String> collectionNames = client.listCollections().getCollectionNames();
            Assert.assertTrue(collectionNames.contains(collectionNameNew));
        }
    }

    // case-06
    @Test(dataProvider = "DisConnectInstance", dataProviderClass = MainClass.class)
    public void testShowCollectionsWithoutConnect(MilvusClient client, String collectionName) {
        ListCollectionsResponse res = client.listCollections();
        assert(!res.getResponse().ok());
    }

    // case-07
    @Test(dataProvider = "Collection", dataProviderClass = MainClass.class)
    public void testDropCollection(MilvusClient client, String collectionName) throws InterruptedException {
        Response res = client.dropCollection(collectionName);
        assert(res.ok());
        Thread.currentThread().sleep(1000);
        List<String> collectionNames = client.listCollections().getCollectionNames();
        Assert.assertFalse(collectionNames.contains(collectionName));
    }

    // case-08
    @Test(dataProvider = "Collection", dataProviderClass = MainClass.class)
    public void testDropCollectionNotExisted(MilvusClient client, String collectionName) {
        Response res = client.dropCollection(collectionName+"_");
        assert(!res.ok());
        List<String> collectionNames = client.listCollections().getCollectionNames();
        Assert.assertTrue(collectionNames.contains(collectionName));
    }

    // case-09
    @Test(dataProvider = "DisConnectInstance", dataProviderClass = MainClass.class)
    public void testDropCollectionWithoutConnect(MilvusClient client, String collectionName) {
        Response res = client.dropCollection(collectionName);
        assert(!res.ok());
    }

    // case-10
    // TODO
    @Test(dataProvider = "Collection", dataProviderClass = MainClass.class)
    public void testDescribeCollection(MilvusClient client, String collectionName) {
        GetCollectionInfoResponse res = client.getCollectionInfo(collectionName);
        assert(res.getResponse().ok());
        CollectionMapping collectionSchema = res.getCollectionMapping().get();
        List<Map<String,Object>> fields = (List<Map<String, Object>>) collectionSchema.getFields();
        int dim = 0;
        for(Map<String,Object> field: fields){
            if ("float_vector".equals(field.get("field"))) {
                JSONObject jsonObject = JSONObject.parseObject(field.get("params").toString());
                String dimParams = jsonObject.getString("params");
                dim = Utils.getParam(dimParams,"dim");
            }
            continue;
        }
        String segmentParams = collectionSchema.getParamsInJson();
        Assert.assertEquals(dim, dimension);
        Assert.assertEquals(collectionSchema.getCollectionName(), collectionName);
        Assert.assertEquals(Utils.getParam(segmentParams,"segment_row_count"), segmentRowCount);
    }

    // case-11
    @Test(dataProvider = "DisConnectInstance", dataProviderClass = MainClass.class)
    public void testDescribeCollectionWithoutConnect(MilvusClient client, String collectionName) {
        GetCollectionInfoResponse res = client.getCollectionInfo(collectionName);
        assert(!res.getResponse().ok());
    }

    // case-12
    @Test(dataProvider = "Collection", dataProviderClass = MainClass.class)
    public void testHasCollectionNotExisted(MilvusClient client, String collectionName) {
        HasCollectionResponse res = client.hasCollection(collectionName+"_");
        assert(res.getResponse().ok());
        Assert.assertFalse(res.hasCollection());
    }

    // case-13
    @Test(dataProvider = "DisConnectInstance", dataProviderClass = MainClass.class)
    public void testHasCollectionWithoutConnect(MilvusClient client, String collectionName) {
        HasCollectionResponse res = client.hasCollection(collectionName);
        assert(!res.getResponse().ok());
    }

    // case-14
    @Test(dataProvider = "Collection", dataProviderClass = MainClass.class)
    public void testHasCollection(MilvusClient client, String collectionName) {
        HasCollectionResponse res = client.hasCollection(collectionName);
        assert(res.getResponse().ok());
        Assert.assertTrue(res.hasCollection());
    }
}