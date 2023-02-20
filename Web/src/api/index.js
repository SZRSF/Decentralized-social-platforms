//当前这个模块：API进行统一管理
import requests from "@/api/ajax";

// 三级联动接口
// 发请求:axios发请求返回的结果Promise

// 登录
export const dspUserLogin = (data) => requests({url: '/login', data,method: 'post'});

// 获取所有家的列表
export const dspFamilyList = () => requests({'url':'/family',method:'get'});

// 获取作品内容
export const dspWorksList = () => requests({'url':'/posts/',method:'get'});
