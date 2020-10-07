import React, {ChangeEventHandler, FC} from "react";
import {Input} from "antd";

interface iHeader{
    addItem:(content:string)=>void;
}

const Header:FC<iHeader>=({
  addItem})=>{
    let inputValue=""

        return(
            <div>
                <Input addonBefore={"todoList"}
                       onPressEnter={(e)=>{
                           addItem(inputValue)
                       }}
                       onChange={(e)=>{
                           inputValue = e.target.value
                       }}>

                </Input>
            </div>
        )

}

export default Header;