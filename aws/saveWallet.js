var AWS = require('aws-sdk');
var s3 = new AWS.S3();

exports.handler = async (event) => {
  var coinsFromGo = JSON.parse(event);
  if (coinsFromGo != null) {
    var bucket = "serverlesspocbucket";
    var file = "data.json";
    var getParams = {
      Bucket: bucket,
      Key: file 
    };
    return await s3.getObject(getParams).promise()
      .then((res) => {
        var objOld = JSON.parse(res.Body);
        let date = new Date();
        var timestamp = date.getTime();
        let coinsToSave = 0;
        var obj = new Object();
        var arrJson = new Array();
        if (objOld != null && objOld.length > 0) {
          coinsToSave = objOld[0].coins;
        }
        coinsToSave += coinsFromGo;
        obj["coins"] = coinsToSave;
        obj["timestamp"] = timestamp;
        arrJson.push(obj);
        var uploadParams = {
          Bucket: bucket,
          Body: JSON.stringify(arrJson),
          Key: file
        };
        return s3.upload(uploadParams).promise()
          .then((res) => {
            const response = {
              statusCode: 200,
              body: 1
            };
            return response;
          }).catch((err) => {
            console.log(err);
            const response = {
              statusCode: 400,
              body: -1,
            };
            return response;
          });
      });
  }
};

