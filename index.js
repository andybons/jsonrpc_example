var rpc = require('json-rpc2');

var client = rpc.Client.$create(1234, 'localhost');

client.connectSocket(function(err, conn) {
  if (err) {
    console.log(err);
    return;
  }

  conn.call('Calculator.Add', {X: 1, Y: 2}, function(err, result) {
    if (err) {
      console.log(err);
      conn.end();
      return;
    }
    console.log('1 + 2 = ' + result);
    conn.end();
  });
});


