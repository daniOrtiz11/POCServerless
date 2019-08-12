/*
Func to set the input addr into p2p network json
*/
var AWS = require('aws-sdk');
var s3 = new AWS.S3();

exports.handler = async (event) => {
  var coinsFromGo = JSON.parse(event);
  console.log(coinsFromGo)
  if (coinsFromGo != null) {
    var bucket = "serverlesspocbucket";
    var file = "data.json";
    var getParams = {
      Bucket: bucket,
      Key: file
    };
    return await s3.getObject(getParams).promise()
      .then((res) => {
        console.log(res)
        console.log(res.Body)
        var objOld = JSON.parse(res.Body);
        let date = new Date();
        var timestamp = date.getTime();
        let coinsToSave = 0;
        var obj = new Object();
        var arrJson = new Array();
        if (objOld != null && objOld.length > 0) {
          coinsToSave = objOld["coins"];
        }
        coinsToSave = + coinsFromGo;
        obj["coins"] = coinsToSave;
        obj["timestamp"] = timestamp;
        arrJson.push(obj);
        var uploadParams = {
          Bucket: bucket,
          Body: JSON.stringify(arrJson),
          Key: file
        };
        s3.upload(uploadParams, function (err, data) {
          //handle error
          if (err) {
            return "err";
          }
          //success
          if (data) {
            return "ok";
          }
        });
        const response = {
          statusCode: 200,
          body: 0
        };
        return response;
      })
  }
};

