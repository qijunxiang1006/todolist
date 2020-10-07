import React, {FC} from "react";
import { Checkbox } from 'antd';
import { Button } from 'antd';
import { SearchOutlined } from '@ant-design/icons';

interface iItem {
    id:number;
    content:string;
    status:boolean;
    close:(id:number)=>void;
    changeStatus:(id:number,status:boolean)=>void;
}

const ListItem:FC<iItem>=({
                          id,
                          content,
                          status,
                          close,
                          changeStatus
                      })=>(
    <div>
        <Checkbox onChange={(e)=>{
            changeStatus(id,e.target.checked)
        }}
        checked={status}>{content}</Checkbox>
        <Button type="primary" shape="circle" onClick={(e)=>{
           close(id)
        }}>
            close
        </Button>
    </div>
)

export default ListItem;