//Load require libs
const express = require('express');
var request = require('request');
var os = require("os");

//configure port
let PORT = process.env.PORT || "3000";
PORT = parseInt(PORT)

const APIVERSION = "v1"
const SOURCE_URL = "http://api.pnd.gs/v1/sources/"
//Create an instance of Express
app = express();

//Define routes
//GET /
app.get('/', (req, resp) => {
    resp.status(200).json({ 
        "status": "ok",
        "hostname": os.hostname()
    })
})

var router = express.Router()

    //route GET /api/v1/time
    router.get('/time', function (req, res) {
        res.status(200)
        res.type("application/json")
        res.json({ "time": new Date().toISOString().replace('T', ' ') })
    })

    //route GET /api/v1/source
    router.get('/source', (req, res) => {
	let config = {
	"url": SOURCE_URL,
	'User-Agent': 'Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.96 Safari/537.36'
	}

        request(config, (error, response, body) => {
            if (error) {
                console.log('error:', error)
                res.status(500).end()
                return
            }
            var data = {
                "latest": [],
                "popular": []
            }
            body = JSON.parse(body)
            for (let index = 0; index < body.length; index++) {
                const item = body[index];
                data["latest"].push( item["endpoints"]["latest"] )
                data["popular"].push( item["endpoints"]["popular"] )  
            }

            res.status(200)
            res.type("application/json")
            res.json( data )
        });

        
    })

// Registering /api/v1 as route for /time & /source
app.use(`/api/${APIVERSION}`, router)
//Start the server
app.listen(PORT, () => {
    console.log("Application started at %s on port %d", new Date(), PORT);
})
