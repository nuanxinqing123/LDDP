import{d as a,r as e,e as t,aJ as l,A as d,B as r,aG as o,aF as s,aL as n,E as i,bt as u,bu as c,aI as m,aT as p,aU as f,bI as _,bJ as b,b1 as x,b0 as g,b6 as j,aW as y,b7 as h,bP as v,bQ as I,bN as k}from"./arco.6fa229d2.js";import{_ as C}from"./index.3d5f452d.js";/* empty css              *//* empty css               *//* empty css               *//* empty css               *//* empty css                *//* empty css              *//* empty css                *//* empty css               *//* empty css               *//* empty css               *//* empty css               */import{m as Y}from"./moment.20765af1.js";import{e as A,o as D}from"./record.54ae7714.js";import"./chart.f78d3971.js";import"./vue.17197c5c.js";import"./_commonjs-dynamic-modules.30ae7933.js";const H=i("div",null,[i("p",null," 检索内容为「订单类名」时，欲检索内容应该填写需要检索的「订单类型」 "),i("p",null," 检索内容为「订单号」时，欲检索内容应该填写需要检索的「订单号」 "),i("p",null," 检索内容为「订单任务变量」时，欲检索内容应该填写需要检索的「订单内的任务变量值」 ")],-1),M={key:0},E={key:1},V={key:2},w={key:3},B={key:4},J={key:5},N=a({__name:"table-data",setup(a){const C={current:1,pageSize:20},N=[{title:"ID",dataIndex:"ID"},{title:"订单类名",dataIndex:"order_task_type"},{title:"订单号",dataIndex:"order_id"},{title:"订单归属用户",dataIndex:"order_uid"},{title:"订单任务数量",dataIndex:"order_number"},{title:"订单状态",dataIndex:"order_state",slotName:"order_state"},{title:"下单IP",dataIndex:"order_ip"},{title:"下单时间",dataIndex:"CreatedAt"}],P=e({table:[]}),T=e({fc:"",state:"",data:""}),U=t(!1),z=t(!1),F=async a=>{await A(a).then((a=>{C.current=a.data.page,P.table=a.data.page_data,P.table.forEach((a=>{a.CreatedAt=Y(a.CreatedAt).format("YYYY-MM-DD HH:mm:ss")}))}))};F(C.current);const G=()=>{U.value=!0},L=a=>{T.fc=a,z.value="订单类名"===a},O=a=>{T.state=a},Q=()=>{T.fc?T.data?(D(T).then((a=>{2e3===a.code&&(p.success("检索成功"),T.data="",P.table=a.data,P.table.forEach((a=>{a.CreatedAt=Y(a.CreatedAt).format("YYYY-MM-DD HH:mm:ss")})))})),C.current=1,U.value=!1):p.error("检索值不能为空"):p.error("检索方法不能为空")},S=()=>{U.value=!1},W=()=>{F(1)};return(a,e)=>{const t=l("icon-search"),p=f,Y=l("icon-refresh"),A=_,D=b,q=x,K=g,R=j,X=y,Z=h,$=v,aa=I,ea=k;return d(),r(m,null,[o(p,{type:"outline",onClick:G},{icon:s((()=>[o(t)])),default:s((()=>[n("搜索")])),_:1}),n("   "),o(p,{type:"outline",status:"success",onClick:W},{icon:s((()=>[o(Y)])),default:s((()=>[n("重置")])),_:1}),o(X,{width:500,visible:U.value,"unmount-on-close":"","ok-text":"检索",onOk:Q,onCancel:S},{title:s((()=>[n(" 订单搜索 ")])),default:s((()=>[i("div",null,[o(R,{model:T,layout:"vertical"},{default:s((()=>[o(q,{field:"email",label:"检索内容"},{default:s((()=>[o(D,{placeholder:"选择检索项",onChange:L},{default:s((()=>[o(A,null,{default:s((()=>[n("订单类名")])),_:1}),o(A,null,{default:s((()=>[n("订单号")])),_:1}),o(A,null,{default:s((()=>[n("订单任务变量")])),_:1})])),_:1})])),_:1}),u(o(q,{field:"state",label:"检索状态"},{default:s((()=>[o(D,{placeholder:"选择检索项",onChange:O},{default:s((()=>[o(A,null,{default:s((()=>[n("等待中")])),_:1}),o(A,null,{default:s((()=>[n("进行中")])),_:1}),o(A,null,{default:s((()=>[n("已完成")])),_:1}),o(A,null,{default:s((()=>[n("已终止")])),_:1}),o(A,null,{default:s((()=>[n("退款中")])),_:1}),o(A,null,{default:s((()=>[n("已退款")])),_:1})])),_:1})])),_:1},512),[[c,z.value]]),o(q,{field:"points",label:"欲检索内容"},{default:s((()=>[o(K,{modelValue:T.data,"onUpdate:modelValue":e[0]||(e[0]=a=>T.data=a),placeholder:"值（支持模糊搜索）"},null,8,["modelValue"])])),_:1})])),_:1},8,["model"]),H])])),_:1},8,["visible"]),o(Z),o(aa,{pagination:!1,columns:N,data:P.table},{order_state:s((({record:a})=>[-1===a.order_state?(d(),r("span",M,[o($,{bordered:"",color:"green"},{default:s((()=>[n("等待中")])),_:1})])):0===a.order_state?(d(),r("span",E,[o($,{bordered:"",color:"green"},{default:s((()=>[n("进行中")])),_:1})])):3===a.order_state?(d(),r("span",V,[o($,{bordered:"",color:"green"},{default:s((()=>[n("退款中")])),_:1})])):1===a.order_state?(d(),r("span",w,[o($,{bordered:"",color:"gray"},{default:s((()=>[n("已完成")])),_:1})])):2===a.order_state?(d(),r("span",B,[o($,{bordered:"",color:"gray"},{default:s((()=>[n("已终止")])),_:1})])):(d(),r("span",J,[o($,{bordered:"",color:"gray"},{default:s((()=>[n("已退款")])),_:1})]))])),_:1},8,["data"]),o(ea,{style:{float:"right","margin-top":"10px","margin-right":"20px","margin-bottom":"10px"},total:10*C.current,simple:"",onChange:F},null,8,["total"])],64)}}}),P={class:"container"},T={class:"left-side",style:{"margin-top":"12px"}},U={class:"panel"},z=C(a({name:"Task",setup:a=>(a,e)=>{const t=l("Breadcrumb");return d(),r("div",P,[o(t,{items:["订单管理"]}),i("div",T,[i("div",U,[o(N)])])])}}),[["__scopeId","data-v-418ca0c6"]]);export{z as default};