import{c as S,d as v,r as k,A as y,aC as b,aF as t,aG as e,D as p,u as c,E as o,bB as A,bE as E,bF as w,bC as L,bi as $,bD as W,e as h,aJ as F,B as G,b8 as I,b9 as N}from"./arco.d4bcacf5.js";import{i as T,o as D,_ as z}from"./index.408bb41d.js";/* empty css               *//* empty css               *//* empty css               *//* empty css                */import{u as O,a as V}from"./chart-option.1589ed4e.js";/* empty css                */import{L as x}from"./chart.b90db84f.js";import"./vue.5c43e2ab.js";function J(){const _=T();return{isDark:S(()=>_.theme==="dark")}}function P(){return D.get("/v3/api/admin/panel/data")}function j(){return D.get("/v3/api/admin/panel/data/chart")}const q={class:"content-wrap"},H={class:"content"},K={class:"content-wrap"},M={class:"content"},Q={class:"content-wrap"},R={class:"content"},U={class:"content-wrap"},X={class:"content"},Y={class:"content-wrap"},Z={class:"content"},ee={class:"content-wrap"},te={class:"content"},ae={class:"content-wrap"},se={class:"content"},oe={class:"content-wrap"},re={class:"content"},ne=v({__name:"site-data",setup(_){const{isDark:n}=J(),s=k({to_day_user_consume:0,to_day_order:0,all_order:0,all_user:0,to_day_recharge:0,to_day_register_user:0,active_user:0,effective_user:0});return(async()=>{const{data:r}=await P();s.to_day_user_consume=r.to_day_user_consume,s.to_day_order=r.to_day_order,s.all_order=r.all_order,s.all_user=r.all_user,s.to_day_recharge=r.to_day_recharge,s.to_day_register_user=r.to_day_register_user,s.active_user=r.active_user,s.effective_user=r.effective_user})(),(r,g)=>{const i=A,d=E,u=w,a=L,l=$,m=W;return y(),b(d,{class:"general-card",title:"\u7AD9\u70B9\u6570\u636E","header-style":{paddingBottom:"12px"}},{default:t(()=>[e(m,{cols:24,"col-gap":12,"row-gap":12},{default:t(()=>[e(a,{span:{xs:12,sm:12,md:12,lg:12,xl:6,xxl:6}},{default:t(()=>[e(u,{style:{width:"100%"}},{default:t(()=>[e(d,{bordered:!1,style:p({background:c(n)?"linear-gradient(180deg, #284991 0%, #122B62 100%)":"linear-gradient(180deg, #f2f9fe 0%, #e6f4fe 100%)"})},{default:t(()=>[o("div",q,[o("div",H,[e(i,{title:"\u65E5\u6D88\u8D39\u6570",value:s.to_day_user_consume,"value-from":0,animation:"","show-group-separator":""},null,8,["value"])])])]),_:1},8,["style"])]),_:1})]),_:1}),e(a,{span:{xs:12,sm:12,md:12,lg:12,xl:6,xxl:6}},{default:t(()=>[e(u,{style:{width:"100%"}},{default:t(()=>[e(d,{bordered:!1,style:p({background:c(n)?" linear-gradient(180deg, #3D492E 0%, #263827 100%)":"linear-gradient(180deg, #F5FEF2 0%, #E6FEEE 100%)"})},{default:t(()=>[o("div",K,[o("div",M,[e(i,{title:"\u65E5\u8BA2\u5355\u6570",value:s.to_day_order,"value-from":0,animation:"","show-group-separator":""},null,8,["value"])])])]),_:1},8,["style"])]),_:1})]),_:1}),e(a,{span:{xs:12,sm:12,md:12,lg:12,xl:6,xxl:6}},{default:t(()=>[e(u,{style:{width:"100%"}},{default:t(()=>[e(d,{bordered:!1,style:p({background:c(n)?"linear-gradient(180deg, #294B94 0%, #0F275C 100%)":"linear-gradient(180deg, #f2f9fe 0%, #e6f4fe 100%)"})},{default:t(()=>[o("div",Q,[o("div",R,[e(i,{title:"\u603B\u8BA2\u5355\u6570",value:s.all_order,"value-from":0,animation:"","show-group-separator":""},null,8,["value"])])])]),_:1},8,["style"])]),_:1})]),_:1}),e(a,{span:{xs:12,sm:12,md:12,lg:12,xl:6,xxl:6}},{default:t(()=>[e(u,{style:{width:"100%"}},{default:t(()=>[e(d,{bordered:!1,style:p({background:c(n)?"linear-gradient(180deg, #312565 0%, #201936 100%)":"linear-gradient(180deg, #F7F7FF 0%, #ECECFF 100%)"})},{default:t(()=>[o("div",U,[o("div",X,[e(i,{title:"\u6CE8\u518C\u7528\u6237\u6570",value:s.all_user,"value-from":0,animation:"","show-group-separator":""},null,8,["value"])])])]),_:1},8,["style"])]),_:1})]),_:1}),e(a,{span:{xs:12,sm:12,md:12,lg:12,xl:6,xxl:6}},{default:t(()=>[e(u,{style:{width:"100%"}},{default:t(()=>[e(d,{bordered:!1,style:p({background:c(n)?"linear-gradient(180deg, #312565 0%, #201936 100%)":"linear-gradient(180deg, #F7F7FF 0%, #ECECFF 100%)"})},{default:t(()=>[o("div",Y,[o("div",Z,[e(i,{title:"\u65E5\u5145\u503C\u6570",value:s.to_day_recharge,"value-from":0,animation:"","show-group-separator":""},null,8,["value"])])])]),_:1},8,["style"])]),_:1})]),_:1}),e(a,{span:{xs:12,sm:12,md:12,lg:12,xl:6,xxl:6}},{default:t(()=>[e(u,{style:{width:"100%"}},{default:t(()=>[e(d,{bordered:!1,style:p({background:c(n)?"linear-gradient(180deg, #284991 0%, #122B62 100%)":"linear-gradient(180deg, #f2f9fe 0%, #e6f4fe 100%)"})},{default:t(()=>[o("div",ee,[o("div",te,[e(i,{title:"\u65E5\u6CE8\u518C\u6570",value:s.to_day_register_user,"value-from":0,animation:"","show-group-separator":""},null,8,["value"])])])]),_:1},8,["style"])]),_:1})]),_:1}),e(a,{span:{xs:12,sm:12,md:12,lg:12,xl:6,xxl:6}},{default:t(()=>[e(u,{style:{width:"100%"}},{default:t(()=>[e(d,{bordered:!1,style:p({background:c(n)?" linear-gradient(180deg, #3D492E 0%, #263827 100%)":"linear-gradient(180deg, #F5FEF2 0%, #E6FEEE 100%)"})},{default:t(()=>[o("div",ae,[o("div",se,[e(l,{content:"\u4E03\u65E5\u5185\u6D88\u8D39\u6216\u767B\u5F55\u7684\u7528\u6237","background-color":"#722ED1"},{default:t(()=>[e(i,{title:"\u6D3B\u8DC3\u7528\u6237\u6570",value:s.active_user,"value-from":0,animation:"","show-group-separator":""},null,8,["value"])]),_:1})])])]),_:1},8,["style"])]),_:1})]),_:1}),e(a,{span:{xs:12,sm:12,md:12,lg:12,xl:6,xxl:6}},{default:t(()=>[e(u,{style:{width:"100%"}},{default:t(()=>[e(d,{bordered:!1,style:p({background:c(n)?"linear-gradient(180deg, #284991 0%, #122B62 100%)":"linear-gradient(180deg, #f2f9fe 0%, #e6f4fe 100%)"})},{default:t(()=>[o("div",oe,[o("div",re,[e(l,{content:"\u4F59\u989D < 0\u7684\u7528\u6237","background-color":"#722ED1"},{default:t(()=>[e(i,{title:"\u6709\u6548\u7528\u6237\u6570",value:s.effective_user,"value-from":0,animation:"","show-group-separator":""},null,8,["value"])]),_:1})])])]),_:1},8,["style"])]),_:1})]),_:1})]),_:1})]),_:1})}}}),le=v({__name:"content-chart",setup(_){function n(a){return{type:"text",bottom:"8",...a,style:{text:"",textAlign:"center",fill:"#4E5969",fontSize:12}}}const{loading:s,setLoading:f}=O(!0),r=h([]),g=h([]),i=h([n({left:"2.6%"}),n({right:0})]);(async()=>{const{data:a}=await j();r.value=a.seven_days_date,g.value=a.seven_days_order})();const{chartOption:u}=V(()=>({grid:{left:"2.6%",right:"0",top:"10",bottom:"30"},xAxis:{type:"category",offset:2,data:r.value,boundaryGap:!1,axisLabel:{color:"#4E5969",formatter(a,l){return l===0||l===r.value.length-1?"":`${a}`}},axisLine:{show:!1},axisTick:{show:!1},splitLine:{show:!0,interval:a=>a===0?!1:a!==r.value.length-1,lineStyle:{color:"#E5E8EF"}},axisPointer:{show:!0,lineStyle:{color:"#23ADFF",width:2}}},yAxis:{type:"value",axisLine:{show:!1},axisLabel:{formatter(a,l){return l===0?a:`${a}`}},splitLine:{show:!0,lineStyle:{type:"dashed",color:"#E5E8EF"}}},tooltip:{trigger:"axis",formatter(a){const[l]=a;return`<div>
            <p class="tooltip-title">${l.axisValueLabel}</p>
            <div class="content-panel"><span>\u8BA2\u5355\u6570\uFF1A</span><span class="tooltip-value">${Number(l.value).toLocaleString()} </span></div>
          </div>`},className:"echarts-tooltip-diy"},graphic:{elements:i.value},series:[{data:g.value,type:"line",smooth:!0,symbolSize:12,emphasis:{focus:"series",itemStyle:{borderWidth:2}},lineStyle:{width:3,color:new x(0,0,1,0,[{offset:0,color:"rgba(30, 231, 255, 1)"},{offset:.5,color:"rgba(36, 154, 255, 1)"},{offset:1,color:"rgba(111, 66, 251, 1)"}])},showSymbol:!1,areaStyle:{opacity:.8,color:new x(0,0,0,1,[{offset:0,color:"rgba(17, 126, 255, 0.16)"},{offset:1,color:"rgba(17, 128, 255, 0)"}])}}]}));return f(!1),(a,l)=>{const m=F("Chart"),B=E,C=w;return y(),b(C,{loading:c(s),style:{width:"100%"}},{default:t(()=>[e(B,{class:"general-card","header-style":{paddingBottom:0},"body-style":{paddingTop:"20px"},title:"\u4E03\u65E5\u7F51\u7AD9\u8BA2\u5355\u6570\u636E"},{default:t(()=>[e(m,{height:"289px",option:c(u)},null,8,["option"])]),_:1})]),_:1},8,["loading"])}}}),ie=_=>(I("data-v-7baefa80"),_=_(),N(),_),de={class:"container"},ue=ie(()=>o("div",{class:"header"},null,-1)),ce={class:"left-side",style:{"margin-top":"12px"}},_e={class:"panel"},pe=v({__name:"index",setup(_){return(n,s)=>{const f=F("Breadcrumb");return y(),G("div",de,[e(f,{items:["\u7BA1\u7406\u4E2D\u5FC3","\u5DE5\u4F5C\u53F0"]}),ue,o("div",ce,[o("div",_e,[e(ne),e(le)])])])}}});const Fe=z(pe,[["__scopeId","data-v-7baefa80"]]);export{Fe as default};