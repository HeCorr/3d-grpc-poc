const grpc = require('@grpc/grpc-js')
const protoLoader = require('@grpc/proto-loader')

const PROTO_PATH = __dirname + (process.env.ENV == 'docker' ? '/../render_server.proto' : '/../../proto/render_server.proto')
const packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {
        keepCase: true,
        longs: String,
        enums: String,
        defaults: true,
        oneofs: true
    }
)
const renderer = grpc.loadPackageDefinition(packageDefinition).render_server
const client = new renderer.Renderer(process.env.ENV == 'docker' ? '3d-grpc-server:5656' : 'localhost:5656', grpc.credentials.createInsecure())

function requestRender(objects) {
    return new Promise((resolve, reject) => {
        client.RequestRender({
            objects: objects
        }, (err, resp) => {
            if (err) return reject(err)
            if (!resp) return reject("empty response")
            resolve([resp.imageBytes, resp.renderTime])
        })
    })
}

module.exports = { requestRender }
