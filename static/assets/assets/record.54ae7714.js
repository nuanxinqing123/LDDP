import{o as a}from"./index.3d5f452d.js";function e(e){return a.get(`/v2/api/user/login/division/data?page=${e}`)}function t(e){return a.get(`/v2/api/user/recharge/division/data?page=${e}`)}function i(e){return a.get(`/v2/api/user/consumption/division/data?page=${e}`)}function n(){return a.get("/v2/api/user/order/data")}function r(e){return a.get(`/v2/api/user/order/division/data?page=${e}`)}function s(e){return a.get(`/v2/api/user/order/refund?order=${e}`)}function o(e){return a.get(`/v3/api/order/division/data?page=${e}`)}function u(e){return a.get(`/v3/api/order/search?type=${e.fc}&state=${e.state}&s=${e.data}`)}function d(e){return a.get(`/v3/api/consumption/division/data?page=${e}`)}function g(e){return a.get(`/v3/api/consumption/search?s=${e.data}`)}function p(e){return a.get(`/v3/api/recharge/division/data?page=${e}`)}function c(e){return a.get(`/v3/api/recharge/search?type=${e.fc}&s=${e.data}`)}function v(e){return a.get(`/v3/api/login/division/data?page=${e}`)}function f(e){return a.get(`/v3/api/login/search?s=${e.data}`)}export{r as a,i as b,t as c,e as d,o as e,d as f,n as g,g as h,p as i,c as j,v as k,f as l,u as o,s as r};