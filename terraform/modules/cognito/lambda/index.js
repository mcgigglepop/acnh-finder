const AWS = require('aws-sdk');
const ddb = new AWS.DynamoDB.DocumentClient();

exports.handler = async (event, context, callback) => {
  const userId = event.request.userAttributes.sub;
  const defaultHemisphere = "unset"; // or pull from a form if you're fancy later

  if (!userId) {
    console.error("No user sub provided.");
    return callback("Error: No user sub provided.");
  }

  const params = {
    TableName: "UserProfiles",
    Item: {
      user_id: userId,
      hemisphere: defaultHemisphere,
    },
  };

  try {
    await ddb.put(params).promise();
    console.log(`User ${userId} profile created.`);
    callback(null, event);
  } catch (err) {
    console.error("Error writing to DynamoDB", err);
    callback(err);
  }
};
