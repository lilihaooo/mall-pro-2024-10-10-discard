/*! 
 Build based on admin-vue 
 Time : 1721439181000 */
import{D as e,K as r,L as t}from"./087AC4D233B64EB0index.3a42c3d0.js";const a=(a,i)=>{const l={},n=e([]);return{children:n,addChild:e=>{l[e.uid]=e,n.value=((e,a,i)=>r(e.subTree).filter((e=>{var r;return t(e)&&(null==(r=e.type)?void 0:r.name)===a&&!!e.component})).map((e=>e.component.uid)).map((e=>i[e])).filter((e=>!!e)))(a,i,l)},removeChild:e=>{delete l[e],n.value=n.value.filter((r=>r.uid!==e))}}};export{a as u};
