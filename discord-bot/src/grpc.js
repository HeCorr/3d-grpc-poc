var grpc = require('@grpc/grpc-js');
var protoLoader = require('@grpc/proto-loader');

// TODO: fix not working in Docker
var PROTO_PATH = __dirname + '/../../proto/render_server.proto';
var packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {
        keepCase: true,
        longs: String,
        enums: String,
        defaults: true,
        oneofs: true
    });
var renderer = grpc.loadPackageDefinition(packageDefinition).render_server;

var client = new renderer.Renderer('localhost:5656', grpc.credentials.createInsecure());

function requestRender(objects) {
    return new Promise((resolve, reject) => {
        client.RequestRender({
            objects: objects
        }, (err, resp) => {
            if (err) return reject(err);
            if (!resp) return reject("empty response");
            resolve([resp.imageBytes, resp.renderTime])
        })
    })
}

module.exports = { requestRender };
