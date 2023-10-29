const apiAdapter = require("../../apiAdapter")

const { URL_SERVICE_TRANSACTION } = process.env

const api = apiAdapter(URL_SERVICE_TRANSACTION)

module.exports = async (req, res) => {

    try {
        const user = await api.get("/api/v1/courses")
        return res.json(user.data)
    }catch (error) {

        if(error.code == "ECONNREFUSED"){
            return res.status(500).json({status : "error", message : "service tidak teresedia"})
        }

        const {status , data} = error.response;
        return res.status(status).json(data)

    }

}