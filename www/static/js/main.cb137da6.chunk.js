(window["webpackJsonpbraille-print"]=window["webpackJsonpbraille-print"]||[]).push([[0],{170:function(e,t,n){e.exports=n(337)},175:function(e,t,n){},176:function(e,t,n){},337:function(e,t,n){"use strict";n.r(t);var a=n(0),o=n.n(a),i=n(4),r=n.n(i),c=(n(175),n(176),n(152)),l=n(153),s=n(154),d=n(155),u=n(69),m=n(168),p=n(341),h=n(340),g=n(342),f=n(345),E=n(339),v=n(344),b=n(59),y=n(346),C=n(347),w=n(343),k=n(8),O=n(102),S=n.n(O),x=(n(272),n(273),p.a.Header),j=p.a.Content,T=p.a.Footer,M=p.a.Sider,A=h.a.confirm,H=function(e){function t(){var e,n;Object(c.a)(this,t);for(var a=arguments.length,i=new Array(a),r=0;r<a;r++)i[r]=arguments[r];return(n=Object(s.a)(this,(e=Object(d.a)(t)).call.apply(e,[this].concat(i)))).state={editorState:S.a.createEditorState(""),outputHTML:"",version:"Init",visible:!1,connected:!1,printConnected:!1,initConnected:!0},n.handleChange=function(e){n.setState({editorState:e,outputHTML:e.toHTML()})},n.handleMission=function(){var e=new Headers({"Access-Control-Allow-Origin":"*","Content-Type":"application/json"});fetch("http://localhost:4396/createMission",{method:"post",headers:e,body:JSON.stringify({html:n.state.outputHTML}),mode:"cors"}).then((function(e){return e.json()})).then((function(e){"OK"===e.code?g.a.success({message:"\u6210\u529f\uff01",description:"\u60a8\u5df2\u63d0\u4ea4\u4e86\u4e00\u4e2a\u6253\u5370\u4efb\u52a1\u5230\u6253\u5370\u673a\uff01"}):"NotFound"===e.code&&g.a.error({message:"\u9519\u8bef\uff01",description:"\u60a8\u7684\u7cfb\u7edf\u627e\u4e0d\u5230\u6253\u5370\u673a\u5728\u54ea\u91cc\uff01\u8bf7\u786e\u8ba4\u6253\u5370\u673a\u8fde\u63a5\u72b6\u6001\u540e\u91cd\u65b0\u542f\u52a8\u7cfb\u7edf\uff01"})})).catch((function(e){console.log(e)}))},n.connect=function(){var e=new Headers({"Access-Control-Allow-Origin":"*","Content-Type":"application/json"});fetch("http://localhost:4396/ping",{method:"post",headers:e,mode:"cors"}).then((function(e){return e.json()})).then((function(e){"OK"===e.code&&!1===n.state.connected&&(g.a.success({message:"\u6210\u529f\uff01",description:"\u7cfb\u7edf\u8fde\u63a5\u6210\u529f\uff01"}),n.setState({visible:!1,connected:!0,initConnected:!1,version:e.version}),"OK"===e.printConnected&&n.setState({printConnected:!0}))})).catch(function(e){this.setState({visible:!0,connected:!1,initConnected:!0}),console.log(e)}.bind(Object(u.a)(n)))},n.handleExit=function(){A({centered:!0,title:"\u60a8\u662f\u5426\u786e\u5b9a\u8981\u9000\u51fa\u6253\u5370\u673a?",content:"\u8bf7\u60a8\u518d\u4e09\u786e\u8ba4\u662f\u5426\u8981\u9000\u51fa\u6253\u5370\u673a\u7a0b\u5e8f\uff0c\n\u5728\u60a8\u70b9\u51fb\u786e\u8ba4\u6309\u94ae\u540e\uff0c\u7cfb\u7edf\u5c06\u4f1a\u5f3a\u5236\u5173\u95ed\u6253\u5370\u673a\u7a0b\u5e8f\uff0c\u8fd9\u5c06\u4f1a\u5bfc\u81f4\u60a8\u4e22\u5931\u672a\u4f20\u8f93\u7ed9\u6253\u5370\u673a\u7684\u6253\u5370\u6307\u4ee4\u961f\u5217\uff0c\u540c\u65f6\u4e5f\u610f\u5473\u7740\u60a8\u7684\u6253\u5370\u673a\u5c06\u4ee5\u4e0d\u53ef\u63a7\u72b6\u6001\u7ee7\u7eed\u6253\u5370\u672a\u5b8c\u6210\u7684\u6307\u4ee4\u96c6\uff01",onOk:n.exitSystem})},n.exitSystem=function(){var e=new Headers({"Access-Control-Allow-Origin":"*","Content-Type":"application/json"});fetch("http://localhost:4396/exit",{method:"post",headers:e,mode:"cors"}).then((function(e){return e.json()})).then((function(e){"OK"===e.code&&(g.a.success({message:"\u6210\u529f\uff01",description:"\u7cfb\u7edf\u5373\u5c06\u5728\u4e09\u79d2\u540e\u81ea\u52a8\u9000\u51fa"}),setTimeout(window.close,3e3))})).catch((function(e){g.a.error({message:"\u5931\u8d25\uff01",description:"\u7cfb\u7edf\u4e0e\u670d\u52a1\u5668\u8fde\u63a5\u5931\u8d25\uff01"}),console.log(e)}))},n.createMission=function(){!1!==n.state.printConnected?""!==n.state.outputHTML?(A({title:"\u786e\u8ba4\u6253\u5370",content:"\u8bf7\u95ee\u60a8\u786e\u8ba4\u6253\u5370\u5417\uff1f",onOk:n.handleMission}),console.log(n.state.outputHTML)):g.a.error({message:"\u9519\u8bef\uff01",description:"\u60a8\u4e0d\u80fd\u63d0\u4ea4\u4e00\u4e2a\u7a7a\u6587\u672c\u5230\u6253\u5370\u673a\uff01"}):h.a.error({title:"\u9519\u8bef\uff01",content:o.a.createElement("div",null,"\u60a8\u7684\u7cfb\u7edf\u627e\u4e0d\u5230\u6253\u5370\u673a\u5728\u54ea\u91cc\uff01",o.a.createElement("br",null),"\u8bf7\u786e\u8ba4\u6253\u5370\u673a\u8fde\u63a5\u72b6\u6001\u540e\u91cd\u65b0\u542f\u52a8\u7cfb\u7edf\uff01")})},n}return Object(m.a)(t,e),Object(l.a)(t,[{key:"componentDidMount",value:function(){setInterval(this.connect,3e3)}},{key:"render",value:function(){var e,t,n=this.state,a=n.editorState,i=n.connected,r=n.printConnected,c=n.initConnected,l=n.version;return e=i?o.a.createElement(f.a,{status:"processing",text:"Running"}):o.a.createElement(f.a,{status:"error",text:"Stop!"}),t=r?o.a.createElement(f.a,{status:"processing",text:"Running"}):o.a.createElement(f.a,{status:"error",text:"Stop!"}),o.a.createElement(E.a,{tip:"Loading...",spinning:c},o.a.createElement(p.a,null,o.a.createElement(M,{breakpoint:"lg",collapsedWidth:"0",width:"300",onBreakpoint:function(e){console.log(e)},onCollapse:function(e,t){console.log(e,t)},style:{height:"100vh",backgroundColor:"#555555"}},o.a.createElement("div",{className:"LOGO"},"\u76f2\u6587\u6253\u5370\u673a"),o.a.createElement(v.a,{style:{backgroundColor:"#555555"},title:"\u7cfb\u7edf\u72b6\u6001",column:1,bordered:!0},o.a.createElement(v.a.Item,{label:"\u670d\u52a1\u5668\u7aef\u8fde\u63a5\uff1a",style:{backgroundColor:"#555555"}},e),o.a.createElement(v.a.Item,{label:"\u6253\u5370\u673a\u7aef\u8fde\u63a5\uff1a",style:{backgroundColor:"#555555"}},t)),o.a.createElement("div",null,o.a.createElement(b.a,{style:{float:"bottom"},type:"danger",onClick:this.handleExit,block:!0},"\u9000\u51fa\u7cfb\u7edf"))),o.a.createElement(p.a,null,o.a.createElement(x,{style:{background:"#fff",padding:0,fontSize:"2em",textAlign:"left",paddingLeft:30}},"> \u6253\u5370\u4efb\u52a1\u521b\u5efa\u9762\u677f"),o.a.createElement(y.a,{message:"\u8b66\u544a",description:"\u672c\u7248\u672c\u7a0b\u5e8f\u4e3a\u5f00\u53d1\u8005\u9884\u89c8\u7248\uff0c\u975e\u6700\u7ec8Release\u53d1\u5e03\u7248\u672c\uff0c\u53ef\u80fd\u5b58\u5728\u8bef\u64cd\u4f5c\u7684\u98ce\u9669\uff0c\u8bf7\u52a1\u5fc5\u4eba\u4e3a\u786e\u8ba4\u6253\u5370\u673a\u7684\u72b6\u6001\uff01",type:"error"}),o.a.createElement(j,{style:{margin:"24px 16px 0"}},o.a.createElement("div",{style:{padding:24,background:"#fff",minHeight:360}},o.a.createElement("div",null,o.a.createElement("div",null,o.a.createElement(S.a,{controls:["undo","redo"],value:a,onChange:this.handleChange})),o.a.createElement("hr",null),o.a.createElement("div",{style:{textAlign:"right"}},o.a.createElement(b.a,{type:"primary",size:"large",onClick:this.createMission},"\u6253\u5370"))))),o.a.createElement(T,{style:{textAlign:"center"}},"Braille Print \xa92019 Created by Shanghai Jian Qiao U. Braille Print Team Project",o.a.createElement(C.a,{color:"purple"},l))),o.a.createElement(h.a,{centered:!0,title:"\u9519\u8bef\uff01",visible:this.state.visible,footer:null},o.a.createElement(w.a,{status:"404",title:"404",subTitle:o.a.createElement("div",null,o.a.createElement("p",null,"\u7cfb\u7edf\u65e0\u6cd5\u8fde\u63a5\u5230\u670d\u52a1\u5668"),o.a.createElement("p",null,"\u8bf7\u91cd\u65b0\u542f\u52a8\u7cfb\u7edf\u670d\u52a1\u5668"))}),","),o.a.createElement(h.a,{centered:!0,title:"\u7cfb\u7edf\u6b63\u5728\u8fde\u63a5\u670d\u52a1\u5668\u4e2d\uff01",visible:this.state.initConnected,footer:null},o.a.createElement(w.a,{icon:o.a.createElement(k.a,{type:"smile",theme:"twoTone"}),title:"\u7cfb\u7edf\u6b63\u5728\u8fde\u63a5\u4e2d!\u8bf7\u7a0d\u540e"}))))}}]),t}(o.a.Component);var L=function(){return o.a.createElement("div",{className:"App"},o.a.createElement(H,null))};Boolean("localhost"===window.location.hostname||"[::1]"===window.location.hostname||window.location.hostname.match(/^127(?:\.(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}$/));r.a.render(o.a.createElement(L,null),document.getElementById("root")),"serviceWorker"in navigator&&navigator.serviceWorker.ready.then((function(e){e.unregister()}))}},[[170,1,2]]]);
//# sourceMappingURL=main.cb137da6.chunk.js.map