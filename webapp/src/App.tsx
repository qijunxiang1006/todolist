import React from 'react';
import './App.css';
import Header from "./header";
import ToDoList from "./list";
import API from "./api/api";

interface listItem {
    id: number;
    value: string;
    status:boolean
}
const api=new API()
class App extends React.Component<any, any>{
    state={
        list:[] as any[],
        uuid:""
    }
    async componentDidMount() {
        if (localStorage.getItem("uuid") == null) {
            await api.getUUID().then(res => {
                localStorage.setItem("uuid", res.data.data)
                this.setState({uuid:res.data.data})
                console.log(this.state)
            })
        }else{
            let uuid = localStorage.getItem("uuid")
            if(uuid!=null){
                await api.getList(uuid).then(
                    res=>{
                        if (res.data.code===0){
                            console.log(this.state)
                            this.setState(res.data.data)
                            console.log(this.state)
                        }
                    }
                )
            }
        }
    }

    addItem = (content: string) => {
        api.postList(this.state.uuid, content).then(res=>{
            if (res.data.code===0) {
                const id = res.data.data
                let list = this.state.list
                const obj: listItem = {
                    id: id,
                    value: content,
                    status: false
                }
                list.push(obj)
                this.setState({list: list})
            }
        })
    };
    close = async (id:number)=>{
        await api.deleteList(this.state.uuid, id).then(res=>{
            if (res.data.code === 0){
                let list = this.state.list
                for (let i = 0; i < list.length; i++) {
                    if (list[i].id === id) {
                        list.splice(i, 1)
                        break
                    }
                }
                this.setState({list:list})
            }
        })
    }
    changeStatus = (id:number,status:boolean)=>{
        api.putList(this.state.uuid,id,status).then(res=>{
            if (res.data.code === 0){
                let list = this.state.list
                for (let i = 0; i < list.length; i++) {
                    if (list[i].id === id) {
                        list[i].status=status;
                        break
                    }
                }
                this.setState({list:list})
            }
        })
    }

    render() {
        return (
            <React.Fragment>
                <Header addItem={this.addItem}/>
                <ToDoList list={this.state.list} close={this.close} changeStatus={this.changeStatus}/>:<div/>
            </React.Fragment>
        );
    }
}


export default App;