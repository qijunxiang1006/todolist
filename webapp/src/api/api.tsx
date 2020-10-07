import axios from 'axios';

const todoPath = "/list"
const uuidPath = "/uuid"
const host="http://"+window.location.host;

const request=axios.create({
    baseURL: host,
    headers: {
        'Content-Type': 'application/json;',
    },
});

class API {

    public getUUID(){
        return request.get(uuidPath)
    }

    public getList(uuid:string){
        return request.get(todoPath+"?uuid="+uuid)
    }
    public postList(uuid:string,content:string){
        return request.post(todoPath+"?uuid="+uuid,JSON.stringify(
            {
                value:content
            }
        ))
    }

    public putList(uuid:string,id:number,status:boolean){
        return request.put(todoPath+"?uuid="+uuid,JSON.stringify(
            {
                id:id,
                status:status
            }
        ))
    }
    public deleteList(uuid:string,id:number){
        return request.delete(todoPath+"?uuid="+uuid+"&&id="+id)
    }

}

export default API