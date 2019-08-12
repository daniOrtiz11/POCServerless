/*
Func to get the current wallet state
*/
var AWS = require('aws-sdk');
// get reference to S3 client 
var s3 = new AWS.S3();
 
exports.handler = async (event) => {
    console.log(event)
    var bucket = "serverlesspocbucket";
    var file = "data.json";
    var getParams = {
    Bucket: bucket, 
    Key: file 
    };
    return await s3.getObject(getParams).promise()
    .then((res) => {
        console.log("into await")
        var obj = JSON.parse(res.Body); 
        var coinsWallet = 0;
        for (let i = 0; i < obj.length; i++){
            coinsWallet =+ obj[i].Amount;
        }
        const response = {
            statusCode: 200,
            body: coinsWallet,
        };
        console.log("returning coins: " + coinsWallet)
        return response;
    })
    .catch((err) => {
        return err;
    });
};

