var PROTO_PATH = __dirname + '/../server/src/proto/render_server.proto';

var grpc = require('@grpc/grpc-js');
var protoLoader = require('@grpc/proto-loader');
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

var client = new renderer.Renderer('localhost:1313', grpc.credentials.createInsecure());

function requestRender() {
    return new Promise((resolve, reject) => {
        client.RequestRender({
            objects: [
                {
                    shape: "Cube",
                    position: {
                        x: 0,
                        y: 0,
                        z: 0,
                    },
                    scale: {
                        x: 1,
                        y: 1,
                        z: 1,
                    }
                }
            ]
        }, (err, resp) => {
            if (err) reject(err);
            resolve(resp.imageBytes)
        })
    })
}

module.exports = { requestRender };
