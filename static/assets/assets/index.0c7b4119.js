import{d as h,r as g,A as a,aC as i,aF as t,aG as o,bB as M,b5 as O,bC as Y,b7 as I,bD as L,e as U,B as n,aI as D,aH as $,aL as _,E as s,aM as d,aT as H,b4 as N,aU as P,bK as S,bL as G,bM as T,aZ as V,aJ as J}from"./arco.d4bcacf5.js";import{_ as x}from"./index.408bb41d.js";/* empty css               *//* empty css               *//* empty css               *//* empty css                */import{g as K,a as Z,r as j}from"./record.b322ae83.js";/* empty css              *//* empty css               *//* empty css               *//* empty css               *//* empty css                *//* empty css              */import{m as k}from"./moment.b4a440e7.js";import"./chart.b90db84f.js";import"./vue.5c43e2ab.js";import"./_commonjs-dynamic-modules.30ae7933.js";const q=h({__name:"data-panel",setup(A){const r=g({order_today:0,order_seven:0,order_thirty:0,order_total:0});return(async()=>{await K().then(u=>{r.order_today=u.data.order_today,r.order_seven=u.data.order_seven,r.order_thirty=u.data.order_thirty,r.order_total=u.data.order_total})})(),(u,y)=>{const c=M,f=O,m=Y,B=I,p=L;return a(),i(p,{cols:24,"row-gap":16,class:"panel"},{default:t(()=>[o(m,{class:"panel-col",span:{xs:12,sm:12,md:12,lg:12,xl:12,xxl:6}},{default:t(()=>[o(f,null,{default:t(()=>[o(c,{title:"\u4ECA\u65E5\u8BA2\u5355",value:r.order_today,"value-from":0,animation:"","show-group-separator":""},null,8,["value"])]),_:1})]),_:1}),o(m,{class:"panel-col",span:{xs:12,sm:12,md:12,lg:12,xl:12,xxl:6},style:{"border-right":"none"}},{default:t(()=>[o(f,null,{default:t(()=>[o(c,{title:"7\u65E5\u8BA2\u5355",value:r.order_seven,"value-from":0,animation:"","show-group-separator":""},null,8,["value"])]),_:1})]),_:1}),o(m,{class:"panel-col",span:{xs:12,sm:12,md:12,lg:12,xl:12,xxl:6}},{default:t(()=>[o(f,null,{default:t(()=>[o(c,{title:"30\u65E5\u8BA2\u5355",value:r.order_thirty,"value-from":0,animation:"","show-group-separator":""},null,8,["value"])]),_:1})]),_:1}),o(m,{class:"panel-col",span:{xs:12,sm:12,md:12,lg:12,xl:12,xxl:6},style:{"border-right":"none"}},{default:t(()=>[o(f,null,{default:t(()=>[o(c,{title:"\u603B\u8BA2\u5355",value:r.order_total,"value-from":0,animation:"","show-group-separator":""},null,8,["value"])]),_:1})]),_:1}),o(m,{span:24},{default:t(()=>[o(B,{class:"panel-border"})]),_:1})]),_:1})}}});const Q=x(q,[["__scopeId","data-v-a5134f71"]]);const R={style:{color:"#86909c"}},W={key:0},X={key:1},ee={key:2},te={key:3},ae={key:4},oe={key:5},re={key:0},se={key:1},ue=h({__name:"formData",setup(A){const r=U(!1),F=g({order_id:""}),u={current:1,pageSize:20},y=g({table:[]}),c=async p=>{await Z(p).then(v=>{u.current=v.data.page,y.table=v.data.page_data,y.table.forEach(l=>{l.CreatedAt=k(l.CreatedAt).format("YYYY-MM-DD HH:mm:ss"),l.UpdatedAt=k(l.UpdatedAt).format("YYYY-MM-DD HH:mm:ss")})})};c(u.current);const f=p=>{F.order_id=p,r.value=!0},m=()=>{j(F.order_id).then(p=>{p.code===2e3&&(H.success("\u5DF2\u7533\u8BF7\u9000\u5355"),c(u.current))}),r.value=!1},B=()=>{r.value=!1};return(p,v)=>{const l=N,b=P,C=S,E=G,w=T,z=V;return a(),n(D,null,[o(E,null,{default:t(()=>[(a(!0),n(D,null,$(y.table,e=>(a(),n("span",{key:e},[(a(),i(C,{key:e.ID,header:"\u521B\u5EFA\u65F6\u95F4\uFF1A"+e.CreatedAt},{extra:t(()=>[e.order_state===-1?(a(),i(l,{key:0,size:"mini",style:{color:"#165dff"}},{default:t(()=>[_("\u7B49\u5F85\u4E2D")]),_:1})):e.order_state===0?(a(),i(l,{key:1,size:"mini",style:{color:"#165dff"}},{default:t(()=>[_("\u8FDB\u884C\u4E2D")]),_:1})):e.order_state===3?(a(),i(l,{key:2,size:"mini",style:{color:"#165dff"}},{default:t(()=>[_("\u9000\u6B3E\u4E2D")]),_:1})):e.order_state===1?(a(),i(l,{key:3,size:"mini",style:{color:"#bedaff"}},{default:t(()=>[_("\u5DF2\u5B8C\u6210")]),_:1})):e.order_state===2?(a(),i(l,{key:4,size:"mini",style:{color:"#bedaff"}},{default:t(()=>[_("\u5DF2\u7EC8\u6B62")]),_:1})):(a(),i(l,{key:5,size:"mini",style:{color:"#bedaff"}},{default:t(()=>[_("\u5DF2\u9000\u6B3E")]),_:1}))]),default:t(()=>[s("div",R,[s("div",null,"\u521B\u5EFA\u65F6\u95F4\uFF1A"+d(e.CreatedAt),1),s("div",null,"\u66F4\u65B0\u65F6\u95F4\uFF1A"+d(e.UpdatedAt),1),s("div",null,"\u8BA2\u5355\u7C7B\u578B\uFF1A"+d(e.order_task_type),1),s("div",null,"\u6D88\u8D39\u70B9\u5238\uFF1A"+d(e.order_tickets)+"\u2002\u70B9\u5238",1),s("div",null,"\u4EFB\u52A1\u6570\u91CF\uFF1A"+d(e.order_number),1),s("div",null,"\u4EFB\u52A1\u53D8\u91CF\uFF1A"+d(e.order_variable),1),s("div",null,[_("\u8BA2\u5355\u72B6\u6001\uFF1A "),e.order_state===-1?(a(),n("span",W,"\u7B49\u5F85\u4E2D")):e.order_state===0?(a(),n("span",X,"\u8FDB\u884C\u4E2D")):e.order_state===1?(a(),n("span",ee,"\u5DF2\u5B8C\u6210")):e.order_state===2?(a(),n("span",te,"\u5DF2\u7EC8\u6B62")):e.order_state===3?(a(),n("span",ae,"\u9000\u6B3E\u4E2D")):(a(),n("span",oe,"\u5DF2\u9000\u6B3E"))]),s("div",null,[_("\u8BA2\u5355\u5B9E\u51B5\uFF1A "),e.order_status!==""?(a(),n("span",re,d(e.order_status),1)):(a(),n("span",se,"\u2002Loading"))]),s("div",null,"\u8BA2\u5355\u5907\u6CE8\uFF1A"+d(e.order_state_reason),1),s("div",null,"\u4EFB\u52A1\u5907\u6CE8\uFF1A"+d(e.order_remarks),1)]),s("div",null,[e.order_state===-1||e.order_state===0?(a(),i(b,{key:0,style:{float:"right","margin-top":"10px","margin-right":"20px","margin-bottom":"10px"},shape:"round",type:"primary",size:"small",onClick:pe=>f(e.order_id)},{default:t(()=>[_("\u9000\u5355")]),_:2},1032,["onClick"])):(a(),i(b,{key:1,style:{float:"right","margin-top":"10px","margin-right":"20px","margin-bottom":"10px"},shape:"round",type:"primary",size:"small",disabled:""},{default:t(()=>[_("\u9000\u5355")]),_:1}))])]),_:2},1032,["header"]))]))),128))]),_:1}),o(w,{style:{float:"right","margin-top":"10px","margin-right":"20px","margin-bottom":"10px"},total:u.current*10,simple:"",onChange:c},null,8,["total"]),o(z,{visible:r.value,"onUpdate:visible":v[0]||(v[0]=e=>r.value=e),width:"354px",onOk:m,onCancel:B},{title:t(()=>[_(" \u786E\u8BA4\u9000\u6B3E\u6B64\u8BA2\u5355\uFF1F ")]),default:t(()=>[s("div",null,"\u8BA2\u5355\u7F16\u53F7\uFF1A"+d(F.order_id),1)]),_:1},8,["visible"])],64)}}});const ne=x(ue,[["__scopeId","data-v-186d66b5"]]),le={class:"container"},_e={class:"left-side",style:{"margin-top":"12px"}},de={class:"panel"},ie={name:"Task"},ce=h({...ie,setup(A){return(r,F)=>{const u=J("Breadcrumb");return a(),n("div",le,[o(u,{items:["\u8BB0\u5F55","\u6211\u7684\u8BA2\u5355"]}),s("div",_e,[s("div",de,[o(Q),o(ne)])])])}}});const Me=x(ce,[["__scopeId","data-v-f2d0df0e"]]);export{Me as default};