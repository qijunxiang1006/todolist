import React, {FC} from "react";
import { List, } from 'antd';
// import reqwest from 'reqwest';

import InfiniteScroll from 'react-infinite-scroller';
import ListItem from "../listitem";

interface iList {
    list:[]|any;
    close:(id:number)=>void;
    changeStatus:(id:number,status:boolean)=>void;
}

const ToDoList:FC<iList>=({
      list,
      close,
      changeStatus
  })=> {
   const handleInfiniteOnLoad = () => {
        if ( list!==null&&list.length > 14) {
            return;
        }

    };

    return (
        <div className="demo-infinite-container">
            <InfiniteScroll
                initialLoad={false}
                pageStart={0}
                loadMore={handleInfiniteOnLoad}
                hasMore={true}
                useWindow={false}
            >
                <List
                    dataSource={list}
                    renderItem={(item,index )=> (
                        <List.Item>
                            <ListItem id={list[index].id} content={list[index].value} status={list[index].status}
                                      close={close} changeStatus={changeStatus}/>
                        </List.Item>
                    )}
                >
                </List>
            </InfiniteScroll>
        </div>
    )
}

export default ToDoList;