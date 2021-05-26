/*! jQuery v3.5.1 | (c) JS Foundation and other contributors | jquery.org/license */
!function(e,t){"use strict";"object"==typeof module&&"object"==typeof module.exports?module.exports=e.document?t(e,!0):function(e){if(!e.document)throw new Error("jQuery requires a window with a document");return t(e)}:t(e)}("undefined"!=typeof window?window:this,function(C,e){"use strict";var t=[],r=Object.getPrototypeOf,s=t.slice,g=t.flat?function(e){return t.flat.call(e)}:function(e){return t.concat.apply([],e)},u=t.push,i=t.indexOf,n={},o=n.toString,v=n.hasOwnProperty,a=v.toString,l=a.call(Object),y={},m=function(e){return"function"==typeof e&&"number"!=typeof e.nodeType},x=function(e){return null!=e&&e===e.window},E=C.document,c={type:!0,src:!0,nonce:!0,noModule:!0};function b(e,t,n){var r,i,o=(n=n||E).createElement("script");if(o.text=e,t)for(r in c)(i=t[r]||t.getAttribute&&t.getAttribute(r))&&o.setAttribute(r,i);n.head.appendChild(o).parentNode.removeChild(o)}function w(e){return null==e?e+"":"object"==typeof e||"function"==typeof e?n[o.call(e)]||"object":typeof e}var f="3.5.1",S=function(e,t){return new S.fn.init(e,t)};function p(e){var t=!!e&&"length"in e&&e.length,n=w(e);return!m(e)&&!x(e)&&("array"===n||0===t||"number"==typeof t&&0<t&&t-1 in e)}S.fn=S.prototype={jquery:f,constructor:S,length:0,toArray:function(){return s.call(this)},get:function(e){return null==e?s.call(this):e<0?this[e+this.length]:this[e]},pushStack:function(e){var t=S.merge(this.constructor(),e);return t.prevObject=this,t},each:function(e){return S.each(this,e)},map:function(n){return this.pushStack(S.map(this,function(e,t){return n.call(e,t,e)}))},slice:function(){return this.pushStack(s.apply(this,arguments))},first:function(){return this.eq(0)},last:function(){return this.eq(-1)},even:function(){return this.pushStack(S.grep(this,function(e,t){return(t+1)%2}))},odd:function(){return this.pushStack(S.grep(this,function(e,t){return t%2}))},eq:function(e){var t=this.length,n=+e+(e<0?t:0);return this.pushStack(0<=n&&n<t?[this[n]]:[])},end:function(){return this.prevObject||this.constructor()},push:u,sort:t.sort,splice:t.splice},S.extend=S.fn.extend=function(){var e,t,n,r,i,o,a=arguments[0]||{},s=1,u=arguments.length,l=!1;for("boolean"==typeof a&&(l=a,a=arguments[s]||{},s++),"object"==typeof a||m(a)||(a={}),s===u&&(a=this,s--);s<u;s++)if(null!=(e=arguments[s]))for(t in e)r=e[t],"__proto__"!==t&&a!==r&&(l&&r&&(S.isPlainObject(r)||(i=Array.isArray(r)))?(n=a[t],o=i&&!Array.isArray(n)?[]:i||S.isPlainObject(n)?n:{},i=!1,a[t]=S.extend(l,o,r)):void 0!==r&&(a[t]=r));return a},S.extend({expando:"jQuery"+(f+Math.random()).replace(/\D/g,""),isReady:!0,error:function(e){throw new Error(e)},noop:function(){},isPlainObject:function(e){var t,n;return!(!e||"[object Object]"!==o.call(e))&&(!(t=r(e))||"function"==typeof(n=v.call(t,"constructor")&&t.constructor)&&a.call(n)===l)},isEmptyObject:function(e){var t;for(t in e)return!1;return!0},globalEval:function(e,t,n){b(e,{nonce:t&&t.nonce},n)},each:function(e,t){var n,r=0;if(p(e)){for(n=e.length;r<n;r++)if(!1===t.call(e[r],r,e[r]))break}else for(r in e)if(!1===t.call(e[r],r,e[r]))break;return e},makeArray:function(e,t){var n=t||[];return null!=e&&(p(Object(e))?S.merge(n,"string"==typeof e?[e]:e):u.call(n,e)),n},inArray:function(e,t,n){return null==t?-1:i.call(t,e,n)},merge:function(e,t){for(var n=+t.length,r=0,i=e.length;r<n;r++)e[i++]=t[r];return e.length=i,e},grep:function(e,t,n){for(var r=[],i=0,o=e.length,a=!n;i<o;i++)!t(e[i],i)!==a&&r.push(e[i]);return r},map:function(e,t,n){var r,i,o=0,a=[];if(p(e))for(r=e.length;o<r;o++)null!=(i=t(e[o],o,n))&&a.push(i);else for(o in e)null!=(i=t(e[o],o,n))&&a.push(i);return g(a)},guid:1,support:y}),"function"==typeof Symbol&&(S.fn[Symbol.iterator]=t[Symbol.iterator]),S.each("Boolean Number String Function Array Date RegExp Object Error Symbol".split(" "),function(e,t){n["[object "+t+"]"]=t.toLowerCase()});var d=function(n){var e,d,b,o,i,h,f,g,w,u,l,T,C,a,E,v,s,c,y,S="sizzle"+1*new Date,p=n.document,k=0,r=0,m=ue(),x=ue(),A=ue(),N=ue(),D=function(e,t){return e===t&&(l=!0),0},j={}.hasOwnProperty,t=[],q=t.pop,L=t.push,H=t.push,O=t.slice,P=function(e,t){for(var n=0,r=e.length;n<r;n++)if(e[n]===t)return n;return-1},R="checked|selected|async|autofocus|autoplay|controls|defer|disabled|hidden|ismap|loop|multiple|open|readonly|required|scoped",M="[\\x20\\t\\r\\n\\f]",I="(?:\\\\[\\da-fA-F]{1,6}"+M+"?|\\\\[^\\r\\n\\f]|[\\w-]|[^\0-\\x7f])+",W="\\["+M+"*("+I+")(?:"+M+"*([*^$|!~]?=)"+M+"*(?:'((?:\\\\.|[^\\\\'])*)'|\"((?:\\\\.|[^\\\\\"])*)\"|("+I+"))|)"+M+"*\\]",F=":("+I+")(?:\\((('((?:\\\\.|[^\\\\'])*)'|\"((?:\\\\.|[^\\\\\"])*)\")|((?:\\\\.|[^\\\\()[\\]]|"+W+")*)|.*)\\)|)",B=new RegExp(M+"+","g"),$=new RegExp("^"+M+"+|((?:^|[^\\\\])(?:\\\\.)*)"+M+"+$","g"),_=new RegExp("^"+M+"*,"+M+"*"),z=new RegExp("^"+M+"*([>+~]|"+M+")"+M+"*"),U=new RegExp(M+"|>"),X=new RegExp(F),V=new RegExp("^"+I+"$"),G={ID:new RegExp("^#("+I+")"),CLASS:new RegExp("^\\.("+I+")"),TAG:new RegExp("^("+I+"|[*])"),ATTR:new RegExp("^"+W),PSEUDO:new RegExp("^"+F),CHILD:new RegExp("^:(only|first|last|nth|nth-last)-(child|of-type)(?:\\("+M+"*(even|odd|(([+-]|)(\\d*)n|)"+M+"*(?:([+-]|)"+M+"*(\\d+)|))"+M+"*\\)|)","i"),bool:new RegExp("^(?:"+R+")$","i"),needsContext:new RegExp("^"+M+"*[>+~]|:(even|odd|eq|gt|lt|nth|first|last)(?:\\("+M+"*((?:-\\d)?\\d*)"+M+"*\\)|)(?=[^-]|$)","i")},Y=/HTML$/i,Q=/^(?:input|select|textarea|button)$/i,J=/^h\d$/i,K=/^[^{]+\{\s*\[native \w/,Z=/^(?:#([\w-]+)|(\w+)|\.([\w-]+))$/,ee=/[+~]/,te=new RegExp("\\\\[\\da-fA-F]{1,6}"+M+"?|\\\\([^\\r\\n\\f])","g"),ne=function(e,t){var n="0x"+e.slice(1)-65536;return t||(n<0?String.fromCharCode(n+65536):String.fromCharCode(n>>10|55296,1023&n|56320))},re=/([\0-\x1f\x7f]|^-?\d)|^-$|[^\0-\x1f\x7f-\uFFFF\w-]/g,ie=function(e,t){return t?"\0"===e?"\ufffd":e.slice(0,-1)+"\\"+e.charCodeAt(e.length-1).toString(16)+" ":"\\"+e},oe=function(){T()},ae=be(function(e){return!0===e.disabled&&"fieldset"===e.nodeName.toLowerCase()},{dir:"parentNode",next:"legend"});try{H.apply(t=O.call(p.childNodes),p.childNodes),t[p.childNodes.length].nodeType}catch(e){H={apply:t.length?function(e,t){L.apply(e,O.call(t))}:function(e,t){var n=e.length,r=0;while(e[n++]=t[r++]);e.length=n-1}}}function se(t,e,n,r){var i,o,a,s,u,l,c,f=e&&e.ownerDocument,p=e?e.nodeType:9;if(n=n||[],"string"!=typeof t||!t||1!==p&&9!==p&&11!==p)return n;if(!r&&(T(e),e=e||C,E)){if(11!==p&&(u=Z.exec(t)))if(i=u[1]){if(9===p){if(!(a=e.getElementById(i)))return n;if(a.id===i)return n.push(a),n}else if(f&&(a=f.getElementById(i))&&y(e,a)&&a.id===i)return n.push(a),n}else{if(u[2])return H.apply(n,e.getElementsByTagName(t)),n;if((i=u[3])&&d.getElementsByClassName&&e.getElementsByClassName)return H.apply(n,e.getElementsByClassName(i)),n}if(d.qsa&&!N[t+" "]&&(!v||!v.test(t))&&(1!==p||"object"!==e.nodeName.toLowerCase())){if(c=t,f=e,1===p&&(U.test(t)||z.test(t))){(f=ee.test(t)&&ye(e.parentNode)||e)===e&&d.scope||((s=e.getAttribute("id"))?s=s.replace(re,ie):e.setAttribute("id",s=S)),o=(l=h(t)).length;while(o--)l[o]=(s?"#"+s:":scope")+" "+xe(l[o]);c=l.join(",")}try{return H.apply(n,f.querySelectorAll(c)),n}catch(e){N(t,!0)}finally{s===S&&e.removeAttribute("id")}}}return g(t.replace($,"$1"),e,n,r)}function ue(){var r=[];return function e(t,n){return r.push(t+" ")>b.cacheLength&&delete e[r.shift()],e[t+" "]=n}}function le(e){return e[S]=!0,e}function ce(e){var t=C.createElement("fieldset");try{return!!e(t)}catch(e){return!1}finally{t.parentNode&&t.parentNode.removeChild(t),t=null}}function fe(e,t){var n=e.split("|"),r=n.length;while(r--)b.attrHandle[n[r]]=t}function pe(e,t){var n=t&&e,r=n&&1===e.nodeType&&1===t.nodeType&&e.sourceIndex-t.sourceIndex;if(r)return r;if(n)while(n=n.nextSibling)if(n===t)return-1;return e?1:-1}function de(t){return function(e){return"input"===e.nodeName.toLowerCase()&&e.type===t}}function he(n){return function(e){var t=e.nodeName.toLowerCase();return("input"===t||"button"===t)&&e.type===n}}function ge(t){return function(e){return"form"in e?e.parentNode&&!1===e.disabled?"label"in e?"label"in e.parentNode?e.parentNode.disabled===t:e.disabled===t:e.isDisabled===t||e.isDisabled!==!t&&ae(e)===t:e.disabled===t:"label"in e&&e.disabled===t}}function ve(a){return le(function(o){return o=+o,le(function(e,t){var n,r=a([],e.length,o),i=r.length;while(i--)e[n=r[i]]&&(e[n]=!(t[n]=e[n]))})})}function ye(e){return e&&"undefined"!=typeof e.getElementsByTagName&&e}for(e in d=se.support={},i=se.isXML=function(e){var t=e.namespaceURI,n=(e.ownerDocument||e).documentElement;return!Y.test(t||n&&n.nodeName||"HTML")},T=se.setDocument=function(e){var t,n,r=e?e.ownerDocument||e:p;return r!=C&&9===r.nodeType&&r.documentElement&&(a=(C=r).documentElement,E=!i(C),p!=C&&(n=C.defaultView)&&n.top!==n&&(n.addEventListener?n.addEventListener("unload",oe,!1):n.attachEvent&&n.attachEvent("onunload",oe)),d.scope=ce(function(e){return a.appendChild(e).appendChild(C.createElement("div")),"undefined"!=typeof e.querySelectorAll&&!e.querySelectorAll(":scope fieldset div").length}),d.attributes=ce(function(e){return e.className="i",!e.getAttribute("className")}),d.getElementsByTagName=ce(function(e){return e.appendChild(C.createComment("")),!e.getElementsByTagName("*").length}),d.getElementsByClassName=K.test(C.getElementsByClassName),d.getById=ce(function(e){return a.appendChild(e).id=S,!C.getElementsByName||!C.getElementsByName(S).length}),d.getById?(b.filter.ID=function(e){var t=e.replace(te,ne);return function(e){return e.getAttribute("id")===t}},b.find.ID=function(e,t){if("undefined"!=typeof t.getElementById&&E){var n=t.getElementById(e);return n?[n]:[]}}):(b.filter.ID=function(e){var n=e.replace(te,ne);return function(e){var t="undefined"!=typeof e.getAttributeNode&&e.getAttributeNode("id");return t&&t.value===n}},b.find.ID=function(e,t){if("undefined"!=typeof t.getElementById&&E){var n,r,i,o=t.getElementById(e);if(o){if((n=o.getAttributeNode("id"))&&n.value===e)return[o];i=t.getElementsByName(e),r=0;while(o=i[r++])if((n=o.getAttributeNode("id"))&&n.value===e)return[o]}return[]}}),b.find.TAG=d.getElementsByTagName?function(e,t){return"undefined"!=typeof t.getElementsByTagName?t.getElementsByTagName(e):d.qsa?t.querySelectorAll(e):void 0}:function(e,t){var n,r=[],i=0,o=t.getElementsByTagName(e);if("*"===e){while(n=o[i++])1===n.nodeType&&r.push(n);return r}return o},b.find.CLASS=d.getElementsByClassName&&function(e,t){if("undefined"!=typeof t.getElementsByClassName&&E)return t.getElementsByClassName(e)},s=[],v=[],(d.qsa=K.test(C.querySelectorAll))&&(ce(function(e){var t;a.appendChild(e).innerHTML="<a id='"+S+"'></a><select id='"+S+"-\r\\' msallowcapture=''><option selected=''></option></select>",e.querySelectorAll("[msallowcapture^='']").length&&v.push("[*^$]="+M+"*(?:''|\"\")"),e.querySelectorAll("[selected]").length||v.push("\\["+M+"*(?:value|"+R+")"),e.querySelectorAll("[id~="+S+"-]").length||v.push("~="),(t=C.createElement("input")).setAttribute("name",""),e.appendChild(t),e.querySelectorAll("[name='']").length||v.push("\\["+M+"*name"+M+"*="+M+"*(?:''|\"\")"),e.querySelectorAll(":checked").length||v.push(":checked"),e.querySelectorAll("a#"+S+"+*").length||v.push(".#.+[+~]"),e.querySelectorAll("\\\f"),v.push("[\\r\\n\\f]")}),ce(function(e){e.innerHTML="<a href='' disabled='disabled'></a><select disabled='disabled'><option/></select>";var t=C.createElement("input");t.setAttribute("type","hidden"),e.appendChild(t).setAttribute("name","D"),e.querySelectorAll("[name=d]").length&&v.push("name"+M+"*[*^$|!~]?="),2!==e.querySelectorAll(":enabled").length&&v.push(":enabled",":disabled"),a.appendChild(e).disabled=!0,2!==e.querySelectorAll(":disabled").length&&v.push(":enabled",":disabled"),e.querySelectorAll("*,:x"),v.push(",.*:")})),(d.matchesSelector=K.test(c=a.matches||a.webkitMatchesSelector||a.mozMatchesSelector||a.oMatchesSelector||a.msMatchesSelector))&&ce(function(e){d.disconnectedMatch=c.call(e,"*"),c.call(e,"[s!='']:x"),s.push("!=",F)}),v=v.length&&new RegExp(v.join("|")),s=s.length&&new RegExp(s.join("|")),t=K.test(a.compareDocumentPosition),y=t||K.test(a.contains)?function(e,t){var n=9===e.nodeType?e.documentElement:e,r=t&&t.parentNode;return e===r||!(!r||1!==r.nodeType||!(n.contains?n.contains(r):e.compareDocumentPosition&&16&e.compareDocumentPosition(r)))}:function(e,t){if(t)while(t=t.parentNode)if(t===e)return!0;return!1},D=t?function(e,t){if(e===t)return l=!0,0;var n=!e.compareDocumentPosition-!t.compareDocumentPosition;return n||(1&(n=(e.ownerDocument||e)==(t.ownerDocument||t)?e.compareDocumentPosition(t):1)||!d.sortDetached&&t.compareDocumentPosition(e)===n?e==C||e.ownerDocument==p&&y(p,e)?-1:t==C||t.ownerDocument==p&&y(p,t)?1:u?P(u,e)-P(u,t):0:4&n?-1:1)}:function(e,t){if(e===t)return l=!0,0;var n,r=0,i=e.parentNode,o=t.parentNode,a=[e],s=[t];if(!i||!o)return e==C?-1:t==C?1:i?-1:o?1:u?P(u,e)-P(u,t):0;if(i===o)return pe(e,t);n=e;while(n=n.parentNode)a.unshift(n);n=t;while(n=n.parentNode)s.unshift(n);while(a[r]===s[r])r++;return r?pe(a[r],s[r]):a[r]==p?-1:s[r]==p?1:0}),C},se.matches=function(e,t){return se(e,null,null,t)},se.matchesSelector=function(e,t){if(T(e),d.matchesSelector&&E&&!N[t+" "]&&(!s||!s.test(t))&&(!v||!v.test(t)))try{var n=c.call(e,t);if(n||d.disconnectedMatch||e.document&&11!==e.document.nodeType)return n}catch(e){N(t,!0)}return 0<se(t,C,null,[e]).length},se.contains=function(e,t){return(e.ownerDocument||e)!=C&&T(e),y(e,t)},se.attr=function(e,t){(e.ownerDocument||e)!=C&&T(e);var n=b.attrHandle[t.toLowerCase()],r=n&&j.call(b.attrHandle,t.toLowerCase())?n(e,t,!E):void 0;return void 0!==r?r:d.attributes||!E?e.getAttribute(t):(r=e.getAttributeNode(t))&&r.specified?r.value:null},se.escape=function(e){return(e+"").replace(re,ie)},se.error=function(e){throw new Error("Syntax error, unrecognized expression: "+e)},se.uniqueSort=function(e){var t,n=[],r=0,i=0;if(l=!d.detectDuplicates,u=!d.sortStable&&e.slice(0),e.sort(D),l){while(t=e[i++])t===e[i]&&(r=n.push(i));while(r--)e.splice(n[r],1)}return u=null,e},o=se.getText=function(e){var t,n="",r=0,i=e.nodeType;if(i){if(1===i||9===i||11===i){if("string"==typeof e.textContent)return e.textContent;for(e=e.firstChild;e;e=e.nextSibling)n+=o(e)}else if(3===i||4===i)return e.nodeValue}else while(t=e[r++])n+=o(t);return n},(b=se.selectors={cacheLength:50,createPseudo:le,match:G,attrHandle:{},find:{},relative:{">":{dir:"parentNode",first:!0}," ":{dir:"parentNode"},"+":{dir:"previousSibling",first:!0},"~":{dir:"previousSibling"}},preFilter:{ATTR:function(e){return e[1]=e[1].replace(te,ne),e[3]=(e[3]||e[4]||e[5]||"").replace(te,ne),"~="===e[2]&&(e[3]=" "+e[3]+" "),e.slice(0,4)},CHILD:function(e){return e[1]=e[1].toLowerCase(),"nth"===e[1].slice(0,3)?(e[3]||se.error(e[0]),e[4]=+(e[4]?e[5]+(e[6]||1):2*("even"===e[3]||"odd"===e[3])),e[5]=+(e[7]+e[8]||"odd"===e[3])):e[3]&&se.error(e[0]),e},PSEUDO:function(e){var t,n=!e[6]&&e[2];return G.CHILD.test(e[0])?null:(e[3]?e[2]=e[4]||e[5]||"":n&&X.test(n)&&(t=h(n,!0))&&(t=n.indexOf(")",n.length-t)-n.length)&&(e[0]=e[0].slice(0,t),e[2]=n.slice(0,t)),e.slice(0,3))}},filter:{TAG:function(e){var t=e.replace(te,ne).toLowerCase();return"*"===e?function(){return!0}:function(e){return e.nodeName&&e.nodeName.toLowerCase()===t}},CLASS:function(e){var t=m[e+" "];return t||(t=new RegExp("(^|"+M+")"+e+"("+M+"|$)"))&&m(e,function(e){return t.test("string"==typeof e.className&&e.className||"undefined"!=typeof e.getAttribute&&e.getAttribute("class")||"")})},ATTR:function(n,r,i){return function(e){var t=se.attr(e,n);return null==t?"!="===r:!r||(t+="","="===r?t===i:"!="===r?t!==i:"^="===r?i&&0===t.indexOf(i):"*="===r?i&&-1<t.indexOf(i):"$="===r?i&&t.slice(-i.length)===i:"~="===r?-1<(" "+t.replace(B," ")+" ").indexOf(i):"|="===r&&(t===i||t.slice(0,i.length+1)===i+"-"))}},CHILD:function(h,e,t,g,v){var y="nth"!==h.slice(0,3),m="last"!==h.slice(-4),x="of-type"===e;return 1===g&&0===v?function(e){return!!e.parentNode}:function(e,t,n){var r,i,o,a,s,u,l=y!==m?"nextSibling":"previousSibling",c=e.parentNode,f=x&&e.nodeName.toLowerCase(),p=!n&&!x,d=!1;if(c){if(y){while(l){a=e;while(a=a[l])if(x?a.nodeName.toLowerCase()===f:1===a.nodeType)return!1;u=l="only"===h&&!u&&"nextSibling"}return!0}if(u=[m?c.firstChild:c.lastChild],m&&p){d=(s=(r=(i=(o=(a=c)[S]||(a[S]={}))[a.uniqueID]||(o[a.uniqueID]={}))[h]||[])[0]===k&&r[1])&&r[2],a=s&&c.childNodes[s];while(a=++s&&a&&a[l]||(d=s=0)||u.pop())if(1===a.nodeType&&++d&&a===e){i[h]=[k,s,d];break}}else if(p&&(d=s=(r=(i=(o=(a=e)[S]||(a[S]={}))[a.uniqueID]||(o[a.uniqueID]={}))[h]||[])[0]===k&&r[1]),!1===d)while(a=++s&&a&&a[l]||(d=s=0)||u.pop())if((x?a.nodeName.toLowerCase()===f:1===a.nodeType)&&++d&&(p&&((i=(o=a[S]||(a[S]={}))[a.uniqueID]||(o[a.uniqueID]={}))[h]=[k,d]),a===e))break;return(d-=v)===g||d%g==0&&0<=d/g}}},PSEUDO:function(e,o){var t,a=b.pseudos[e]||b.setFilters[e.toLowerCase()]||se.error("unsupported pseudo: "+e);return a[S]?a(o):1<a.length?(t=[e,e,"",o],b.setFilters.hasOwnProperty(e.toLowerCase())?le(function(e,t){var n,r=a(e,o),i=r.length;while(i--)e[n=P(e,r[i])]=!(t[n]=r[i])}):function(e){return a(e,0,t)}):a}},pseudos:{not:le(function(e){var r=[],i=[],s=f(e.replace($,"$1"));return s[S]?le(function(e,t,n,r){var i,o=s(e,null,r,[]),a=e.length;while(a--)(i=o[a])&&(e[a]=!(t[a]=i))}):function(e,t,n){return r[0]=e,s(r,null,n,i),r[0]=null,!i.pop()}}),has:le(function(t){return function(e){return 0<se(t,e).length}}),contains:le(function(t){return t=t.replace(te,ne),function(e){return-1<(e.textContent||o(e)).indexOf(t)}}),lang:le(function(n){return V.test(n||"")||se.error("unsupported lang: "+n),n=n.replace(te,ne).toLowerCase(),function(e){var t;do{if(t=E?e.lang:e.getAttribute("xml:lang")||e.getAttribute("lang"))return(t=t.toLowerCase())===n||0===t.indexOf(n+"-")}while((e=e.parentNode)&&1===e.nodeType);return!1}}),target:function(e){var t=n.location&&n.location.hash;return t&&t.slice(1)===e.id},root:function(e){return e===a},focus:function(e){return e===C.activeElement&&(!C.hasFocus||C.hasFocus())&&!!(e.type||e.href||~e.tabIndex)},enabled:ge(!1),disabled:ge(!0),checked:function(e){var t=e.nodeName.toLowerCase();return"input"===t&&!!e.checked||"option"===t&&!!e.selected},selected:function(e){return e.parentNode&&e.parentNode.selectedIndex,!0===e.selected},empty:function(e){for(e=e.firstChild;e;e=e.nextSibling)if(e.nodeType<6)return!1;return!0},parent:function(e){return!b.pseudos.empty(e)},header:function(e){return J.test(e.nodeName)},input:function(e){return Q.test(e.nodeName)},button:function(e){var t=e.nodeName.toLowerCase();return"input"===t&&"button"===e.type||"button"===t},text:function(e){var t;return"input"===e.nodeName.toLowerCase()&&"text"===e.type&&(null==(t=e.getAttribute("type"))||"text"===t.toLowerCase())},first:ve(function(){return[0]}),last:ve(function(e,t){return[t-1]}),eq:ve(function(e,t,n){return[n<0?n+t:n]}),even:ve(function(e,t){for(var n=0;n<t;n+=2)e.push(n);return e}),odd:ve(function(e,t){for(var n=1;n<t;n+=2)e.push(n);return e}),lt:ve(function(e,t,n){for(var r=n<0?n+t:t<n?t:n;0<=--r;)e.push(r);return e}),gt:ve(function(e,t,n){for(var r=n<0?n+t:n;++r<t;)e.push(r);return e})}}).pseudos.nth=b.pseudos.eq,{radio:!0,checkbox:!0,file:!0,password:!0,image:!0})b.pseudos[e]=de(e);for(e in{submit:!0,reset:!0})b.pseudos[e]=he(e);function me(){}function xe(e){for(var t=0,n=e.length,r="";t<n;t++)r+=e[t].value;return r}function be(s,e,t){var u=e.dir,l=e.next,c=l||u,f=t&&"parentNode"===c,p=r++;return e.first?function(e,t,n){while(e=e[u])if(1===e.nodeType||f)return s(e,t,n);return!1}:function(e,t,n){var r,i,o,a=[k,p];if(n){while(e=e[u])if((1===e.nodeType||f)&&s(e,t,n))return!0}else while(e=e[u])if(1===e.nodeType||f)if(i=(o=e[S]||(e[S]={}))[e.uniqueID]||(o[e.uniqueID]={}),l&&l===e.nodeName.toLowerCase())e=e[u]||e;else{if((r=i[c])&&r[0]===k&&r[1]===p)return a[2]=r[2];if((i[c]=a)[2]=s(e,t,n))return!0}return!1}}function we(i){return 1<i.length?function(e,t,n){var r=i.length;while(r--)if(!i[r](e,t,n))return!1;return!0}:i[0]}function Te(e,t,n,r,i){for(var o,a=[],s=0,u=e.length,l=null!=t;s<u;s++)(o=e[s])&&(n&&!n(o,r,i)||(a.push(o),l&&t.push(s)));return a}function Ce(d,h,g,v,y,e){return v&&!v[S]&&(v=Ce(v)),y&&!y[S]&&(y=Ce(y,e)),le(function(e,t,n,r){var i,o,a,s=[],u=[],l=t.length,c=e||function(e,t,n){for(var r=0,i=t.length;r<i;r++)se(e,t[r],n);return n}(h||"*",n.nodeType?[n]:n,[]),f=!d||!e&&h?c:Te(c,s,d,n,r),p=g?y||(e?d:l||v)?[]:t:f;if(g&&g(f,p,n,r),v){i=Te(p,u),v(i,[],n,r),o=i.length;while(o--)(a=i[o])&&(p[u[o]]=!(f[u[o]]=a))}if(e){if(y||d){if(y){i=[],o=p.length;while(o--)(a=p[o])&&i.push(f[o]=a);y(null,p=[],i,r)}o=p.length;while(o--)(a=p[o])&&-1<(i=y?P(e,a):s[o])&&(e[i]=!(t[i]=a))}}else p=Te(p===t?p.splice(l,p.length):p),y?y(null,t,p,r):H.apply(t,p)})}function Ee(e){for(var i,t,n,r=e.length,o=b.relative[e[0].type],a=o||b.relative[" "],s=o?1:0,u=be(function(e){return e===i},a,!0),l=be(function(e){return-1<P(i,e)},a,!0),c=[function(e,t,n){var r=!o&&(n||t!==w)||((i=t).nodeType?u(e,t,n):l(e,t,n));return i=null,r}];s<r;s++)if(t=b.relative[e[s].type])c=[be(we(c),t)];else{if((t=b.filter[e[s].type].apply(null,e[s].matches))[S]){for(n=++s;n<r;n++)if(b.relative[e[n].type])break;return Ce(1<s&&we(c),1<s&&xe(e.slice(0,s-1).concat({value:" "===e[s-2].type?"*":""})).replace($,"$1"),t,s<n&&Ee(e.slice(s,n)),n<r&&Ee(e=e.slice(n)),n<r&&xe(e))}c.push(t)}return we(c)}return me.prototype=b.filters=b.pseudos,b.setFilters=new me,h=se.tokenize=function(e,t){var n,r,i,o,a,s,u,l=x[e+" "];if(l)return t?0:l.slice(0);a=e,s=[],u=b.preFilter;while(a){for(o in n&&!(r=_.exec(a))||(r&&(a=a.slice(r[0].length)||a),s.push(i=[])),n=!1,(r=z.exec(a))&&(n=r.shift(),i.push({value:n,type:r[0].replace($," ")}),a=a.slice(n.length)),b.filter)!(r=G[o].exec(a))||u[o]&&!(r=u[o](r))||(n=r.shift(),i.push({value:n,type:o,matches:r}),a=a.slice(n.length));if(!n)break}return t?a.length:a?se.error(e):x(e,s).slice(0)},f=se.compile=function(e,t){var n,v,y,m,x,r,i=[],o=[],a=A[e+" "];if(!a){t||(t=h(e)),n=t.length;while(n--)(a=Ee(t[n]))[S]?i.push(a):o.push(a);(a=A(e,(v=o,m=0<(y=i).length,x=0<v.length,r=function(e,t,n,r,i){var o,a,s,u=0,l="0",c=e&&[],f=[],p=w,d=e||x&&b.find.TAG("*",i),h=k+=null==p?1:Math.random()||.1,g=d.length;for(i&&(w=t==C||t||i);l!==g&&null!=(o=d[l]);l++){if(x&&o){a=0,t||o.ownerDocument==C||(T(o),n=!E);while(s=v[a++])if(s(o,t||C,n)){r.push(o);break}i&&(k=h)}m&&((o=!s&&o)&&u--,e&&c.push(o))}if(u+=l,m&&l!==u){a=0;while(s=y[a++])s(c,f,t,n);if(e){if(0<u)while(l--)c[l]||f[l]||(f[l]=q.call(r));f=Te(f)}H.apply(r,f),i&&!e&&0<f.length&&1<u+y.length&&se.uniqueSort(r)}return i&&(k=h,w=p),c},m?le(r):r))).selector=e}return a},g=se.select=function(e,t,n,r){var i,o,a,s,u,l="function"==typeof e&&e,c=!r&&h(e=l.selector||e);if(n=n||[],1===c.length){if(2<(o=c[0]=c[0].slice(0)).length&&"ID"===(a=o[0]).type&&9===t.nodeType&&E&&b.relative[o[1].type]){if(!(t=(b.find.ID(a.matches[0].replace(te,ne),t)||[])[0]))return n;l&&(t=t.parentNode),e=e.slice(o.shift().value.length)}i=G.needsContext.test(e)?0:o.length;while(i--){if(a=o[i],b.relative[s=a.type])break;if((u=b.find[s])&&(r=u(a.matches[0].replace(te,ne),ee.test(o[0].type)&&ye(t.parentNode)||t))){if(o.splice(i,1),!(e=r.length&&xe(o)))return H.apply(n,r),n;break}}}return(l||f(e,c))(r,t,!E,n,!t||ee.test(e)&&ye(t.parentNode)||t),n},d.sortStable=S.split("").sort(D).join("")===S,d.detectDuplicates=!!l,T(),d.sortDetached=ce(function(e){return 1&e.compareDocumentPosition(C.createElement("fieldset"))}),ce(function(e){return e.innerHTML="<a href='#'></a>","#"===e.firstChild.getAttribute("href")})||fe("type|href|height|width",function(e,t,n){if(!n)return e.getAttribute(t,"type"===t.toLowerCase()?1:2)}),d.attributes&&ce(function(e){return e.innerHTML="<input/>",e.firstChild.setAttribute("value",""),""===e.firstChild.getAttribute("value")})||fe("value",function(e,t,n){if(!n&&"input"===e.nodeName.toLowerCase())return e.defaultValue}),ce(function(e){return null==e.getAttribute("disabled")})||fe(R,function(e,t,n){var r;if(!n)return!0===e[t]?t.toLowerCase():(r=e.getAttributeNode(t))&&r.specified?r.value:null}),se}(C);S.find=d,S.expr=d.selectors,S.expr[":"]=S.expr.pseudos,S.uniqueSort=S.unique=d.uniqueSort,S.text=d.getText,S.isXMLDoc=d.isXML,S.contains=d.contains,S.escapeSelector=d.escape;var h=function(e,t,n){var r=[],i=void 0!==n;while((e=e[t])&&9!==e.nodeType)if(1===e.nodeType){if(i&&S(e).is(n))break;r.push(e)}return r},T=function(e,t){for(var n=[];e;e=e.nextSibling)1===e.nodeType&&e!==t&&n.push(e);return n},k=S.expr.match.needsContext;function A(e,t){return e.nodeName&&e.nodeName.toLowerCase()===t.toLowerCase()}var N=/^<([a-z][^\/\0>:\x20\t\r\n\f]*)[\x20\t\r\n\f]*\/?>(?:<\/\1>|)$/i;function D(e,n,r){return m(n)?S.grep(e,function(e,t){return!!n.call(e,t,e)!==r}):n.nodeType?S.grep(e,function(e){return e===n!==r}):"string"!=typeof n?S.grep(e,function(e){return-1<i.call(n,e)!==r}):S.filter(n,e,r)}S.filter=function(e,t,n){var r=t[0];return n&&(e=":not("+e+")"),1===t.length&&1===r.nodeType?S.find.matchesSelector(r,e)?[r]:[]:S.find.matches(e,S.grep(t,function(e){return 1===e.nodeType}))},S.fn.extend({find:function(e){var t,n,r=this.length,i=this;if("string"!=typeof e)return this.pushStack(S(e).filter(function(){for(t=0;t<r;t++)if(S.contains(i[t],this))return!0}));for(n=this.pushStack([]),t=0;t<r;t++)S.find(e,i[t],n);return 1<r?S.uniqueSort(n):n},filter:function(e){return this.pushStack(D(this,e||[],!1))},not:function(e){return this.pushStack(D(this,e||[],!0))},is:function(e){return!!D(this,"string"==typeof e&&k.test(e)?S(e):e||[],!1).length}});var j,q=/^(?:\s*(<[\w\W]+>)[^>]*|#([\w-]+))$/;(S.fn.init=function(e,t,n){var r,i;if(!e)return this;if(n=n||j,"string"==typeof e){if(!(r="<"===e[0]&&">"===e[e.length-1]&&3<=e.length?[null,e,null]:q.exec(e))||!r[1]&&t)return!t||t.jquery?(t||n).find(e):this.constructor(t).find(e);if(r[1]){if(t=t instanceof S?t[0]:t,S.merge(this,S.parseHTML(r[1],t&&t.nodeType?t.ownerDocument||t:E,!0)),N.test(r[1])&&S.isPlainObject(t))for(r in t)m(this[r])?this[r](t[r]):this.attr(r,t[r]);return this}return(i=E.getElementById(r[2]))&&(this[0]=i,this.length=1),this}return e.nodeType?(this[0]=e,this.length=1,this):m(e)?void 0!==n.ready?n.ready(e):e(S):S.makeArray(e,this)}).prototype=S.fn,j=S(E);var L=/^(?:parents|prev(?:Until|All))/,H={children:!0,contents:!0,next:!0,prev:!0};function O(e,t){while((e=e[t])&&1!==e.nodeType);return e}S.fn.extend({has:function(e){var t=S(e,this),n=t.length;return this.filter(function(){for(var e=0;e<n;e++)if(S.contains(this,t[e]))return!0})},closest:function(e,t){var n,r=0,i=this.length,o=[],a="string"!=typeof e&&S(e);if(!k.test(e))for(;r<i;r++)for(n=this[r];n&&n!==t;n=n.parentNode)if(n.nodeType<11&&(a?-1<a.index(n):1===n.nodeType&&S.find.matchesSelector(n,e))){o.push(n);break}return this.pushStack(1<o.length?S.uniqueSort(o):o)},index:function(e){return e?"string"==typeof e?i.call(S(e),this[0]):i.call(this,e.jquery?e[0]:e):this[0]&&this[0].parentNode?this.first().prevAll().length:-1},add:function(e,t){return this.pushStack(S.uniqueSort(S.merge(this.get(),S(e,t))))},addBack:function(e){return this.add(null==e?this.prevObject:this.prevObject.filter(e))}}),S.each({parent:function(e){var t=e.parentNode;return t&&11!==t.nodeType?t:null},parents:function(e){return h(e,"parentNode")},parentsUntil:function(e,t,n){return h(e,"parentNode",n)},next:function(e){return O(e,"nextSibling")},prev:function(e){return O(e,"previousSibling")},nextAll:function(e){return h(e,"nextSibling")},prevAll:function(e){return h(e,"previousSibling")},nextUntil:function(e,t,n){return h(e,"nextSibling",n)},prevUntil:function(e,t,n){return h(e,"previousSibling",n)},siblings:function(e){return T((e.parentNode||{}).firstChild,e)},children:function(e){return T(e.firstChild)},contents:function(e){return null!=e.contentDocument&&r(e.contentDocument)?e.contentDocument:(A(e,"template")&&(e=e.content||e),S.merge([],e.childNodes))}},function(r,i){S.fn[r]=function(e,t){var n=S.map(this,i,e);return"Until"!==r.slice(-5)&&(t=e),t&&"string"==typeof t&&(n=S.filter(t,n)),1<this.length&&(H[r]||S.uniqueSort(n),L.test(r)&&n.reverse()),this.pushStack(n)}});var P=/[^\x20\t\r\n\f]+/g;function R(e){return e}function M(e){throw e}function I(e,t,n,r){var i;try{e&&m(i=e.promise)?i.call(e).done(t).fail(n):e&&m(i=e.then)?i.call(e,t,n):t.apply(void 0,[e].slice(r))}catch(e){n.apply(void 0,[e])}}S.Callbacks=function(r){var e,n;r="string"==typeof r?(e=r,n={},S.each(e.match(P)||[],function(e,t){n[t]=!0}),n):S.extend({},r);var i,t,o,a,s=[],u=[],l=-1,c=function(){for(a=a||r.once,o=i=!0;u.length;l=-1){t=u.shift();while(++l<s.length)!1===s[l].apply(t[0],t[1])&&r.stopOnFalse&&(l=s.length,t=!1)}r.memory||(t=!1),i=!1,a&&(s=t?[]:"")},f={add:function(){return s&&(t&&!i&&(l=s.length-1,u.push(t)),function n(e){S.each(e,function(e,t){m(t)?r.unique&&f.has(t)||s.push(t):t&&t.length&&"string"!==w(t)&&n(t)})}(arguments),t&&!i&&c()),this},remove:function(){return S.each(arguments,function(e,t){var n;while(-1<(n=S.inArray(t,s,n)))s.splice(n,1),n<=l&&l--}),this},has:function(e){return e?-1<S.inArray(e,s):0<s.length},empty:function(){return s&&(s=[]),this},disable:function(){return a=u=[],s=t="",this},disabled:function(){return!s},lock:function(){return a=u=[],t||i||(s=t=""),this},locked:function(){return!!a},fireWith:function(e,t){return a||(t=[e,(t=t||[]).slice?t.slice():t],u.push(t),i||c()),this},fire:function(){return f.fireWith(this,arguments),this},fired:function(){return!!o}};return f},S.extend({Deferred:function(e){var o=[["notify","progress",S.Callbacks("memory"),S.Callbacks("memory"),2],["resolve","done",S.Callbacks("once memory"),S.Callbacks("once memory"),0,"resolved"],["reject","fail",S.Callbacks("once memory"),S.Callbacks("once memory"),1,"rejected"]],i="pending",a={state:function(){return i},always:function(){return s.done(arguments).fail(arguments),this},"catch":function(e){return a.then(null,e)},pipe:function(){var i=arguments;return S.Deferred(function(r){S.each(o,function(e,t){var n=m(i[t[4]])&&i[t[4]];s[t[1]](function(){var e=n&&n.apply(this,arguments);e&&m(e.promise)?e.promise().progress(r.notify).done(r.resolve).fail(r.reject):r[t[0]+"With"](this,n?[e]:arguments)})}),i=null}).promise()},then:function(t,n,r){var u=0;function l(i,o,a,s){return function(){var n=this,r=arguments,e=function(){var e,t;if(!(i<u)){if((e=a.apply(n,r))===o.promise())throw new TypeError("Thenable self-resolution");t=e&&("object"==typeof e||"function"==typeof e)&&e.then,m(t)?s?t.call(e,l(u,o,R,s),l(u,o,M,s)):(u++,t.call(e,l(u,o,R,s),l(u,o,M,s),l(u,o,R,o.notifyWith))):(a!==R&&(n=void 0,r=[e]),(s||o.resolveWith)(n,r))}},t=s?e:function(){try{e()}catch(e){S.Deferred.exceptionHook&&S.Deferred.exceptionHook(e,t.stackTrace),u<=i+1&&(a!==M&&(n=void 0,r=[e]),o.rejectWith(n,r))}};i?t():(S.Deferred.getStackHook&&(t.stackTrace=S.Deferred.getStackHook()),C.setTimeout(t))}}return S.Deferred(function(e){o[0][3].add(l(0,e,m(r)?r:R,e.notifyWith)),o[1][3].add(l(0,e,m(t)?t:R)),o[2][3].add(l(0,e,m(n)?n:M))}).promise()},promise:function(e){return null!=e?S.extend(e,a):a}},s={};return S.each(o,function(e,t){var n=t[2],r=t[5];a[t[1]]=n.add,r&&n.add(function(){i=r},o[3-e][2].disable,o[3-e][3].disable,o[0][2].lock,o[0][3].lock),n.add(t[3].fire),s[t[0]]=function(){return s[t[0]+"With"](this===s?void 0:this,arguments),this},s[t[0]+"With"]=n.fireWith}),a.promise(s),e&&e.call(s,s),s},when:function(e){var n=arguments.length,t=n,r=Array(t),i=s.call(arguments),o=S.Deferred(),a=function(t){return function(e){r[t]=this,i[t]=1<arguments.length?s.call(arguments):e,--n||o.resolveWith(r,i)}};if(n<=1&&(I(e,o.done(a(t)).resolve,o.reject,!n),"pending"===o.state()||m(i[t]&&i[t].then)))return o.then();while(t--)I(i[t],a(t),o.reject);return o.promise()}});var W=/^(Eval|Internal|Range|Reference|Syntax|Type|URI)Error$/;S.Deferred.exceptionHook=function(e,t){C.console&&C.console.warn&&e&&W.test(e.name)&&C.console.warn("jQuery.Deferred exception: "+e.message,e.stack,t)},S.readyException=function(e){C.setTimeout(function(){throw e})};var F=S.Deferred();function B(){E.removeEventListener("DOMContentLoaded",B),C.removeEventListener("load",B),S.ready()}S.fn.ready=function(e){return F.then(e)["catch"](function(e){S.readyException(e)}),this},S.extend({isReady:!1,readyWait:1,ready:function(e){(!0===e?--S.readyWait:S.isReady)||(S.isReady=!0)!==e&&0<--S.readyWait||F.resolveWith(E,[S])}}),S.ready.then=F.then,"complete"===E.readyState||"loading"!==E.readyState&&!E.documentElement.doScroll?C.setTimeout(S.ready):(E.addEventListener("DOMContentLoaded",B),C.addEventListener("load",B));var $=function(e,t,n,r,i,o,a){var s=0,u=e.length,l=null==n;if("object"===w(n))for(s in i=!0,n)$(e,t,s,n[s],!0,o,a);else if(void 0!==r&&(i=!0,m(r)||(a=!0),l&&(a?(t.call(e,r),t=null):(l=t,t=function(e,t,n){return l.call(S(e),n)})),t))for(;s<u;s++)t(e[s],n,a?r:r.call(e[s],s,t(e[s],n)));return i?e:l?t.call(e):u?t(e[0],n):o},_=/^-ms-/,z=/-([a-z])/g;function U(e,t){return t.toUpperCase()}function X(e){return e.replace(_,"ms-").replace(z,U)}var V=function(e){return 1===e.nodeType||9===e.nodeType||!+e.nodeType};function G(){this.expando=S.expando+G.uid++}G.uid=1,G.prototype={cache:function(e){var t=e[this.expando];return t||(t={},V(e)&&(e.nodeType?e[this.expando]=t:Object.defineProperty(e,this.expando,{value:t,configurable:!0}))),t},set:function(e,t,n){var r,i=this.cache(e);if("string"==typeof t)i[X(t)]=n;else for(r in t)i[X(r)]=t[r];return i},get:function(e,t){return void 0===t?this.cache(e):e[this.expando]&&e[this.expando][X(t)]},access:function(e,t,n){return void 0===t||t&&"string"==typeof t&&void 0===n?this.get(e,t):(this.set(e,t,n),void 0!==n?n:t)},remove:function(e,t){var n,r=e[this.expando];if(void 0!==r){if(void 0!==t){n=(t=Array.isArray(t)?t.map(X):(t=X(t))in r?[t]:t.match(P)||[]).length;while(n--)delete r[t[n]]}(void 0===t||S.isEmptyObject(r))&&(e.nodeType?e[this.expando]=void 0:delete e[this.expando])}},hasData:function(e){var t=e[this.expando];return void 0!==t&&!S.isEmptyObject(t)}};var Y=new G,Q=new G,J=/^(?:\{[\w\W]*\}|\[[\w\W]*\])$/,K=/[A-Z]/g;function Z(e,t,n){var r,i;if(void 0===n&&1===e.nodeType)if(r="data-"+t.replace(K,"-$&").toLowerCase(),"string"==typeof(n=e.getAttribute(r))){try{n="true"===(i=n)||"false"!==i&&("null"===i?null:i===+i+""?+i:J.test(i)?JSON.parse(i):i)}catch(e){}Q.set(e,t,n)}else n=void 0;return n}S.extend({hasData:function(e){return Q.hasData(e)||Y.hasData(e)},data:function(e,t,n){return Q.access(e,t,n)},removeData:function(e,t){Q.remove(e,t)},_data:function(e,t,n){return Y.access(e,t,n)},_removeData:function(e,t){Y.remove(e,t)}}),S.fn.extend({data:function(n,e){var t,r,i,o=this[0],a=o&&o.attributes;if(void 0===n){if(this.length&&(i=Q.get(o),1===o.nodeType&&!Y.get(o,"hasDataAttrs"))){t=a.length;while(t--)a[t]&&0===(r=a[t].name).indexOf("data-")&&(r=X(r.slice(5)),Z(o,r,i[r]));Y.set(o,"hasDataAttrs",!0)}return i}return"object"==typeof n?this.each(function(){Q.set(this,n)}):$(this,function(e){var t;if(o&&void 0===e)return void 0!==(t=Q.get(o,n))?t:void 0!==(t=Z(o,n))?t:void 0;this.each(function(){Q.set(this,n,e)})},null,e,1<arguments.length,null,!0)},removeData:function(e){return this.each(function(){Q.remove(this,e)})}}),S.extend({queue:function(e,t,n){var r;if(e)return t=(t||"fx")+"queue",r=Y.get(e,t),n&&(!r||Array.isArray(n)?r=Y.access(e,t,S.makeArray(n)):r.push(n)),r||[]},dequeue:function(e,t){t=t||"fx";var n=S.queue(e,t),r=n.length,i=n.shift(),o=S._queueHooks(e,t);"inprogress"===i&&(i=n.shift(),r--),i&&("fx"===t&&n.unshift("inprogress"),delete o.stop,i.call(e,function(){S.dequeue(e,t)},o)),!r&&o&&o.empty.fire()},_queueHooks:function(e,t){var n=t+"queueHooks";return Y.get(e,n)||Y.access(e,n,{empty:S.Callbacks("once memory").add(function(){Y.remove(e,[t+"queue",n])})})}}),S.fn.extend({queue:function(t,n){var e=2;return"string"!=typeof t&&(n=t,t="fx",e--),arguments.length<e?S.queue(this[0],t):void 0===n?this:this.each(function(){var e=S.queue(this,t,n);S._queueHooks(this,t),"fx"===t&&"inprogress"!==e[0]&&S.dequeue(this,t)})},dequeue:function(e){return this.each(function(){S.dequeue(this,e)})},clearQueue:function(e){return this.queue(e||"fx",[])},promise:function(e,t){var n,r=1,i=S.Deferred(),o=this,a=this.length,s=function(){--r||i.resolveWith(o,[o])};"string"!=typeof e&&(t=e,e=void 0),e=e||"fx";while(a--)(n=Y.get(o[a],e+"queueHooks"))&&n.empty&&(r++,n.empty.add(s));return s(),i.promise(t)}});var ee=/[+-]?(?:\d*\.|)\d+(?:[eE][+-]?\d+|)/.source,te=new RegExp("^(?:([+-])=|)("+ee+")([a-z%]*)$","i"),ne=["Top","Right","Bottom","Left"],re=E.documentElement,ie=function(e){return S.contains(e.ownerDocument,e)},oe={composed:!0};re.getRootNode&&(ie=function(e){return S.contains(e.ownerDocument,e)||e.getRootNode(oe)===e.ownerDocument});var ae=function(e,t){return"none"===(e=t||e).style.display||""===e.style.display&&ie(e)&&"none"===S.css(e,"display")};function se(e,t,n,r){var i,o,a=20,s=r?function(){return r.cur()}:function(){return S.css(e,t,"")},u=s(),l=n&&n[3]||(S.cssNumber[t]?"":"px"),c=e.nodeType&&(S.cssNumber[t]||"px"!==l&&+u)&&te.exec(S.css(e,t));if(c&&c[3]!==l){u/=2,l=l||c[3],c=+u||1;while(a--)S.style(e,t,c+l),(1-o)*(1-(o=s()/u||.5))<=0&&(a=0),c/=o;c*=2,S.style(e,t,c+l),n=n||[]}return n&&(c=+c||+u||0,i=n[1]?c+(n[1]+1)*n[2]:+n[2],r&&(r.unit=l,r.start=c,r.end=i)),i}var ue={};function le(e,t){for(var n,r,i,o,a,s,u,l=[],c=0,f=e.length;c<f;c++)(r=e[c]).style&&(n=r.style.display,t?("none"===n&&(l[c]=Y.get(r,"display")||null,l[c]||(r.style.display="")),""===r.style.display&&ae(r)&&(l[c]=(u=a=o=void 0,a=(i=r).ownerDocument,s=i.nodeName,(u=ue[s])||(o=a.body.appendChild(a.createElement(s)),u=S.css(o,"display"),o.parentNode.removeChild(o),"none"===u&&(u="block"),ue[s]=u)))):"none"!==n&&(l[c]="none",Y.set(r,"display",n)));for(c=0;c<f;c++)null!=l[c]&&(e[c].style.display=l[c]);return e}S.fn.extend({show:function(){return le(this,!0)},hide:function(){return le(this)},toggle:function(e){return"boolean"==typeof e?e?this.show():this.hide():this.each(function(){ae(this)?S(this).show():S(this).hide()})}});var ce,fe,pe=/^(?:checkbox|radio)$/i,de=/<([a-z][^\/\0>\x20\t\r\n\f]*)/i,he=/^$|^module$|\/(?:java|ecma)script/i;ce=E.createDocumentFragment().appendChild(E.createElement("div")),(fe=E.createElement("input")).setAttribute("type","radio"),fe.setAttribute("checked","checked"),fe.setAttribute("name","t"),ce.appendChild(fe),y.checkClone=ce.cloneNode(!0).cloneNode(!0).lastChild.checked,ce.innerHTML="<textarea>x</textarea>",y.noCloneChecked=!!ce.cloneNode(!0).lastChild.defaultValue,ce.innerHTML="<option></option>",y.option=!!ce.lastChild;var ge={thead:[1,"<table>","</table>"],col:[2,"<table><colgroup>","</colgroup></table>"],tr:[2,"<table><tbody>","</tbody></table>"],td:[3,"<table><tbody><tr>","</tr></tbody></table>"],_default:[0,"",""]};function ve(e,t){var n;return n="undefined"!=typeof e.getElementsByTagName?e.getElementsByTagName(t||"*"):"undefined"!=typeof e.querySelectorAll?e.querySelectorAll(t||"*"):[],void 0===t||t&&A(e,t)?S.merge([e],n):n}function ye(e,t){for(var n=0,r=e.length;n<r;n++)Y.set(e[n],"globalEval",!t||Y.get(t[n],"globalEval"))}ge.tbody=ge.tfoot=ge.colgroup=ge.caption=ge.thead,ge.th=ge.td,y.option||(ge.optgroup=ge.option=[1,"<select multiple='multiple'>","</select>"]);var me=/<|&#?\w+;/;function xe(e,t,n,r,i){for(var o,a,s,u,l,c,f=t.createDocumentFragment(),p=[],d=0,h=e.length;d<h;d++)if((o=e[d])||0===o)if("object"===w(o))S.merge(p,o.nodeType?[o]:o);else if(me.test(o)){a=a||f.appendChild(t.createElement("div")),s=(de.exec(o)||["",""])[1].toLowerCase(),u=ge[s]||ge._default,a.innerHTML=u[1]+S.htmlPrefilter(o)+u[2],c=u[0];while(c--)a=a.lastChild;S.merge(p,a.childNodes),(a=f.firstChild).textContent=""}else p.push(t.createTextNode(o));f.textContent="",d=0;while(o=p[d++])if(r&&-1<S.inArray(o,r))i&&i.push(o);else if(l=ie(o),a=ve(f.appendChild(o),"script"),l&&ye(a),n){c=0;while(o=a[c++])he.test(o.type||"")&&n.push(o)}return f}var be=/^key/,we=/^(?:mouse|pointer|contextmenu|drag|drop)|click/,Te=/^([^.]*)(?:\.(.+)|)/;function Ce(){return!0}function Ee(){return!1}function Se(e,t){return e===function(){try{return E.activeElement}catch(e){}}()==("focus"===t)}function ke(e,t,n,r,i,o){var a,s;if("object"==typeof t){for(s in"string"!=typeof n&&(r=r||n,n=void 0),t)ke(e,s,n,r,t[s],o);return e}if(null==r&&null==i?(i=n,r=n=void 0):null==i&&("string"==typeof n?(i=r,r=void 0):(i=r,r=n,n=void 0)),!1===i)i=Ee;else if(!i)return e;return 1===o&&(a=i,(i=function(e){return S().off(e),a.apply(this,arguments)}).guid=a.guid||(a.guid=S.guid++)),e.each(function(){S.event.add(this,t,i,r,n)})}function Ae(e,i,o){o?(Y.set(e,i,!1),S.event.add(e,i,{namespace:!1,handler:function(e){var t,n,r=Y.get(this,i);if(1&e.isTrigger&&this[i]){if(r.length)(S.event.special[i]||{}).delegateType&&e.stopPropagation();else if(r=s.call(arguments),Y.set(this,i,r),t=o(this,i),this[i](),r!==(n=Y.get(this,i))||t?Y.set(this,i,!1):n={},r!==n)return e.stopImmediatePropagation(),e.preventDefault(),n.value}else r.length&&(Y.set(this,i,{value:S.event.trigger(S.extend(r[0],S.Event.prototype),r.slice(1),this)}),e.stopImmediatePropagation())}})):void 0===Y.get(e,i)&&S.event.add(e,i,Ce)}S.event={global:{},add:function(t,e,n,r,i){var o,a,s,u,l,c,f,p,d,h,g,v=Y.get(t);if(V(t)){n.handler&&(n=(o=n).handler,i=o.selector),i&&S.find.matchesSelector(re,i),n.guid||(n.guid=S.guid++),(u=v.events)||(u=v.events=Object.create(null)),(a=v.handle)||(a=v.handle=function(e){return"undefined"!=typeof S&&S.event.triggered!==e.type?S.event.dispatch.apply(t,arguments):void 0}),l=(e=(e||"").match(P)||[""]).length;while(l--)d=g=(s=Te.exec(e[l])||[])[1],h=(s[2]||"").split(".").sort(),d&&(f=S.event.special[d]||{},d=(i?f.delegateType:f.bindType)||d,f=S.event.special[d]||{},c=S.extend({type:d,origType:g,data:r,handler:n,guid:n.guid,selector:i,needsContext:i&&S.expr.match.needsContext.test(i),namespace:h.join(".")},o),(p=u[d])||((p=u[d]=[]).delegateCount=0,f.setup&&!1!==f.setup.call(t,r,h,a)||t.addEventListener&&t.addEventListener(d,a)),f.add&&(f.add.call(t,c),c.handler.guid||(c.handler.guid=n.guid)),i?p.splice(p.delegateCount++,0,c):p.push(c),S.event.global[d]=!0)}},remove:function(e,t,n,r,i){var o,a,s,u,l,c,f,p,d,h,g,v=Y.hasData(e)&&Y.get(e);if(v&&(u=v.events)){l=(t=(t||"").match(P)||[""]).length;while(l--)if(d=g=(s=Te.exec(t[l])||[])[1],h=(s[2]||"").split(".").sort(),d){f=S.event.special[d]||{},p=u[d=(r?f.delegateType:f.bindType)||d]||[],s=s[2]&&new RegExp("(^|\\.)"+h.join("\\.(?:.*\\.|)")+"(\\.|$)"),a=o=p.length;while(o--)c=p[o],!i&&g!==c.origType||n&&n.guid!==c.guid||s&&!s.test(c.namespace)||r&&r!==c.selector&&("**"!==r||!c.selector)||(p.splice(o,1),c.selector&&p.delegateCount--,f.remove&&f.remove.call(e,c));a&&!p.length&&(f.teardown&&!1!==f.teardown.call(e,h,v.handle)||S.removeEvent(e,d,v.handle),delete u[d])}else for(d in u)S.event.remove(e,d+t[l],n,r,!0);S.isEmptyObject(u)&&Y.remove(e,"handle events")}},dispatch:function(e){var t,n,r,i,o,a,s=new Array(arguments.length),u=S.event.fix(e),l=(Y.get(this,"events")||Object.create(null))[u.type]||[],c=S.event.special[u.type]||{};for(s[0]=u,t=1;t<arguments.length;t++)s[t]=arguments[t];if(u.delegateTarget=this,!c.preDispatch||!1!==c.preDispatch.call(this,u)){a=S.event.handlers.call(this,u,l),t=0;while((i=a[t++])&&!u.isPropagationStopped()){u.currentTarget=i.elem,n=0;while((o=i.handlers[n++])&&!u.isImmediatePropagationStopped())u.rnamespace&&!1!==o.namespace&&!u.rnamespace.test(o.namespace)||(u.handleObj=o,u.data=o.data,void 0!==(r=((S.event.special[o.origType]||{}).handle||o.handler).apply(i.elem,s))&&!1===(u.result=r)&&(u.preventDefault(),u.stopPropagation()))}return c.postDispatch&&c.postDispatch.call(this,u),u.result}},handlers:function(e,t){var n,r,i,o,a,s=[],u=t.delegateCount,l=e.target;if(u&&l.nodeType&&!("click"===e.type&&1<=e.button))for(;l!==this;l=l.parentNode||this)if(1===l.nodeType&&("click"!==e.type||!0!==l.disabled)){for(o=[],a={},n=0;n<u;n++)void 0===a[i=(r=t[n]).selector+" "]&&(a[i]=r.needsContext?-1<S(i,this).index(l):S.find(i,this,null,[l]).length),a[i]&&o.push(r);o.length&&s.push({elem:l,handlers:o})}return l=this,u<t.length&&s.push({elem:l,handlers:t.slice(u)}),s},addProp:function(t,e){Object.defineProperty(S.Event.prototype,t,{enumerable:!0,configurable:!0,get:m(e)?function(){if(this.originalEvent)return e(this.originalEvent)}:function(){if(this.originalEvent)return this.originalEvent[t]},set:function(e){Object.defineProperty(this,t,{enumerable:!0,configurable:!0,writable:!0,value:e})}})},fix:function(e){return e[S.expando]?e:new S.Event(e)},special:{load:{noBubble:!0},click:{setup:function(e){var t=this||e;return pe.test(t.type)&&t.click&&A(t,"input")&&Ae(t,"click",Ce),!1},trigger:function(e){var t=this||e;return pe.test(t.type)&&t.click&&A(t,"input")&&Ae(t,"click"),!0},_default:function(e){var t=e.target;return pe.test(t.type)&&t.click&&A(t,"input")&&Y.get(t,"click")||A(t,"a")}},beforeunload:{postDispatch:function(e){void 0!==e.result&&e.originalEvent&&(e.originalEvent.returnValue=e.result)}}}},S.removeEvent=function(e,t,n){e.removeEventListener&&e.removeEventListener(t,n)},S.Event=function(e,t){if(!(this instanceof S.Event))return new S.Event(e,t);e&&e.type?(this.originalEvent=e,this.type=e.type,this.isDefaultPrevented=e.defaultPrevented||void 0===e.defaultPrevented&&!1===e.returnValue?Ce:Ee,this.target=e.target&&3===e.target.nodeType?e.target.parentNode:e.target,this.currentTarget=e.currentTarget,this.relatedTarget=e.relatedTarget):this.type=e,t&&S.extend(this,t),this.timeStamp=e&&e.timeStamp||Date.now(),this[S.expando]=!0},S.Event.prototype={constructor:S.Event,isDefaultPrevented:Ee,isPropagationStopped:Ee,isImmediatePropagationStopped:Ee,isSimulated:!1,preventDefault:function(){var e=this.originalEvent;this.isDefaultPrevented=Ce,e&&!this.isSimulated&&e.preventDefault()},stopPropagation:function(){var e=this.originalEvent;this.isPropagationStopped=Ce,e&&!this.isSimulated&&e.stopPropagation()},stopImmediatePropagation:function(){var e=this.originalEvent;this.isImmediatePropagationStopped=Ce,e&&!this.isSimulated&&e.stopImmediatePropagation(),this.stopPropagation()}},S.each({altKey:!0,bubbles:!0,cancelable:!0,changedTouches:!0,ctrlKey:!0,detail:!0,eventPhase:!0,metaKey:!0,pageX:!0,pageY:!0,shiftKey:!0,view:!0,"char":!0,code:!0,charCode:!0,key:!0,keyCode:!0,button:!0,buttons:!0,clientX:!0,clientY:!0,offsetX:!0,offsetY:!0,pointerId:!0,pointerType:!0,screenX:!0,screenY:!0,targetTouches:!0,toElement:!0,touches:!0,which:function(e){var t=e.button;return null==e.which&&be.test(e.type)?null!=e.charCode?e.charCode:e.keyCode:!e.which&&void 0!==t&&we.test(e.type)?1&t?1:2&t?3:4&t?2:0:e.which}},S.event.addProp),S.each({focus:"focusin",blur:"focusout"},function(e,t){S.event.special[e]={setup:function(){return Ae(this,e,Se),!1},trigger:function(){return Ae(this,e),!0},delegateType:t}}),S.each({mouseenter:"mouseover",mouseleave:"mouseout",pointerenter:"pointerover",pointerleave:"pointerout"},function(e,i){S.event.special[e]={delegateType:i,bindType:i,handle:function(e){var t,n=e.relatedTarget,r=e.handleObj;return n&&(n===this||S.contains(this,n))||(e.type=r.origType,t=r.handler.apply(this,arguments),e.type=i),t}}}),S.fn.extend({on:function(e,t,n,r){return ke(this,e,t,n,r)},one:function(e,t,n,r){return ke(this,e,t,n,r,1)},off:function(e,t,n){var r,i;if(e&&e.preventDefault&&e.handleObj)return r=e.handleObj,S(e.delegateTarget).off(r.namespace?r.origType+"."+r.namespace:r.origType,r.selector,r.handler),this;if("object"==typeof e){for(i in e)this.off(i,t,e[i]);return this}return!1!==t&&"function"!=typeof t||(n=t,t=void 0),!1===n&&(n=Ee),this.each(function(){S.event.remove(this,e,n,t)})}});var Ne=/<script|<style|<link/i,De=/checked\s*(?:[^=]|=\s*.checked.)/i,je=/^\s*<!(?:\[CDATA\[|--)|(?:\]\]|--)>\s*$/g;function qe(e,t){return A(e,"table")&&A(11!==t.nodeType?t:t.firstChild,"tr")&&S(e).children("tbody")[0]||e}function Le(e){return e.type=(null!==e.getAttribute("type"))+"/"+e.type,e}function He(e){return"true/"===(e.type||"").slice(0,5)?e.type=e.type.slice(5):e.removeAttribute("type"),e}function Oe(e,t){var n,r,i,o,a,s;if(1===t.nodeType){if(Y.hasData(e)&&(s=Y.get(e).events))for(i in Y.remove(t,"handle events"),s)for(n=0,r=s[i].length;n<r;n++)S.event.add(t,i,s[i][n]);Q.hasData(e)&&(o=Q.access(e),a=S.extend({},o),Q.set(t,a))}}function Pe(n,r,i,o){r=g(r);var e,t,a,s,u,l,c=0,f=n.length,p=f-1,d=r[0],h=m(d);if(h||1<f&&"string"==typeof d&&!y.checkClone&&De.test(d))return n.each(function(e){var t=n.eq(e);h&&(r[0]=d.call(this,e,t.html())),Pe(t,r,i,o)});if(f&&(t=(e=xe(r,n[0].ownerDocument,!1,n,o)).firstChild,1===e.childNodes.length&&(e=t),t||o)){for(s=(a=S.map(ve(e,"script"),Le)).length;c<f;c++)u=e,c!==p&&(u=S.clone(u,!0,!0),s&&S.merge(a,ve(u,"script"))),i.call(n[c],u,c);if(s)for(l=a[a.length-1].ownerDocument,S.map(a,He),c=0;c<s;c++)u=a[c],he.test(u.type||"")&&!Y.access(u,"globalEval")&&S.contains(l,u)&&(u.src&&"module"!==(u.type||"").toLowerCase()?S._evalUrl&&!u.noModule&&S._evalUrl(u.src,{nonce:u.nonce||u.getAttribute("nonce")},l):b(u.textContent.replace(je,""),u,l))}return n}function Re(e,t,n){for(var r,i=t?S.filter(t,e):e,o=0;null!=(r=i[o]);o++)n||1!==r.nodeType||S.cleanData(ve(r)),r.parentNode&&(n&&ie(r)&&ye(ve(r,"script")),r.parentNode.removeChild(r));return e}S.extend({htmlPrefilter:function(e){return e},clone:function(e,t,n){var r,i,o,a,s,u,l,c=e.cloneNode(!0),f=ie(e);if(!(y.noCloneChecked||1!==e.nodeType&&11!==e.nodeType||S.isXMLDoc(e)))for(a=ve(c),r=0,i=(o=ve(e)).length;r<i;r++)s=o[r],u=a[r],void 0,"input"===(l=u.nodeName.toLowerCase())&&pe.test(s.type)?u.checked=s.checked:"input"!==l&&"textarea"!==l||(u.defaultValue=s.defaultValue);if(t)if(n)for(o=o||ve(e),a=a||ve(c),r=0,i=o.length;r<i;r++)Oe(o[r],a[r]);else Oe(e,c);return 0<(a=ve(c,"script")).length&&ye(a,!f&&ve(e,"script")),c},cleanData:function(e){for(var t,n,r,i=S.event.special,o=0;void 0!==(n=e[o]);o++)if(V(n)){if(t=n[Y.expando]){if(t.events)for(r in t.events)i[r]?S.event.remove(n,r):S.removeEvent(n,r,t.handle);n[Y.expando]=void 0}n[Q.expando]&&(n[Q.expando]=void 0)}}}),S.fn.extend({detach:function(e){return Re(this,e,!0)},remove:function(e){return Re(this,e)},text:function(e){return $(this,function(e){return void 0===e?S.text(this):this.empty().each(function(){1!==this.nodeType&&11!==this.nodeType&&9!==this.nodeType||(this.textContent=e)})},null,e,arguments.length)},append:function(){return Pe(this,arguments,function(e){1!==this.nodeType&&11!==this.nodeType&&9!==this.nodeType||qe(this,e).appendChild(e)})},prepend:function(){return Pe(this,arguments,function(e){if(1===this.nodeType||11===this.nodeType||9===this.nodeType){var t=qe(this,e);t.insertBefore(e,t.firstChild)}})},before:function(){return Pe(this,arguments,function(e){this.parentNode&&this.parentNode.insertBefore(e,this)})},after:function(){return Pe(this,arguments,function(e){this.parentNode&&this.parentNode.insertBefore(e,this.nextSibling)})},empty:function(){for(var e,t=0;null!=(e=this[t]);t++)1===e.nodeType&&(S.cleanData(ve(e,!1)),e.textContent="");return this},clone:function(e,t){return e=null!=e&&e,t=null==t?e:t,this.map(function(){return S.clone(this,e,t)})},html:function(e){return $(this,function(e){var t=this[0]||{},n=0,r=this.length;if(void 0===e&&1===t.nodeType)return t.innerHTML;if("string"==typeof e&&!Ne.test(e)&&!ge[(de.exec(e)||["",""])[1].toLowerCase()]){e=S.htmlPrefilter(e);try{for(;n<r;n++)1===(t=this[n]||{}).nodeType&&(S.cleanData(ve(t,!1)),t.innerHTML=e);t=0}catch(e){}}t&&this.empty().append(e)},null,e,arguments.length)},replaceWith:function(){var n=[];return Pe(this,arguments,function(e){var t=this.parentNode;S.inArray(this,n)<0&&(S.cleanData(ve(this)),t&&t.replaceChild(e,this))},n)}}),S.each({appendTo:"append",prependTo:"prepend",insertBefore:"before",insertAfter:"after",replaceAll:"replaceWith"},function(e,a){S.fn[e]=function(e){for(var t,n=[],r=S(e),i=r.length-1,o=0;o<=i;o++)t=o===i?this:this.clone(!0),S(r[o])[a](t),u.apply(n,t.get());return this.pushStack(n)}});var Me=new RegExp("^("+ee+")(?!px)[a-z%]+$","i"),Ie=function(e){var t=e.ownerDocument.defaultView;return t&&t.opener||(t=C),t.getComputedStyle(e)},We=function(e,t,n){var r,i,o={};for(i in t)o[i]=e.style[i],e.style[i]=t[i];for(i in r=n.call(e),t)e.style[i]=o[i];return r},Fe=new RegExp(ne.join("|"),"i");function Be(e,t,n){var r,i,o,a,s=e.style;return(n=n||Ie(e))&&(""!==(a=n.getPropertyValue(t)||n[t])||ie(e)||(a=S.style(e,t)),!y.pixelBoxStyles()&&Me.test(a)&&Fe.test(t)&&(r=s.width,i=s.minWidth,o=s.maxWidth,s.minWidth=s.maxWidth=s.width=a,a=n.width,s.width=r,s.minWidth=i,s.maxWidth=o)),void 0!==a?a+"":a}function $e(e,t){return{get:function(){if(!e())return(this.get=t).apply(this,arguments);delete this.get}}}!function(){function e(){if(l){u.style.cssText="position:absolute;left:-11111px;width:60px;margin-top:1px;padding:0;border:0",l.style.cssText="position:relative;display:block;box-sizing:border-box;overflow:scroll;margin:auto;border:1px;padding:1px;width:60%;top:1%",re.appendChild(u).appendChild(l);var e=C.getComputedStyle(l);n="1%"!==e.top,s=12===t(e.marginLeft),l.style.right="60%",o=36===t(e.right),r=36===t(e.width),l.style.position="absolute",i=12===t(l.offsetWidth/3),re.removeChild(u),l=null}}function t(e){return Math.round(parseFloat(e))}var n,r,i,o,a,s,u=E.createElement("div"),l=E.createElement("div");l.style&&(l.style.backgroundClip="content-box",l.cloneNode(!0).style.backgroundClip="",y.clearCloneStyle="content-box"===l.style.backgroundClip,S.extend(y,{boxSizingReliable:function(){return e(),r},pixelBoxStyles:function(){return e(),o},pixelPosition:function(){return e(),n},reliableMarginLeft:function(){return e(),s},scrollboxSize:function(){return e(),i},reliableTrDimensions:function(){var e,t,n,r;return null==a&&(e=E.createElement("table"),t=E.createElement("tr"),n=E.createElement("div"),e.style.cssText="position:absolute;left:-11111px",t.style.height="1px",n.style.height="9px",re.appendChild(e).appendChild(t).appendChild(n),r=C.getComputedStyle(t),a=3<parseInt(r.height),re.removeChild(e)),a}}))}();var _e=["Webkit","Moz","ms"],ze=E.createElement("div").style,Ue={};function Xe(e){var t=S.cssProps[e]||Ue[e];return t||(e in ze?e:Ue[e]=function(e){var t=e[0].toUpperCase()+e.slice(1),n=_e.length;while(n--)if((e=_e[n]+t)in ze)return e}(e)||e)}var Ve=/^(none|table(?!-c[ea]).+)/,Ge=/^--/,Ye={position:"absolute",visibility:"hidden",display:"block"},Qe={letterSpacing:"0",fontWeight:"400"};function Je(e,t,n){var r=te.exec(t);return r?Math.max(0,r[2]-(n||0))+(r[3]||"px"):t}function Ke(e,t,n,r,i,o){var a="width"===t?1:0,s=0,u=0;if(n===(r?"border":"content"))return 0;for(;a<4;a+=2)"margin"===n&&(u+=S.css(e,n+ne[a],!0,i)),r?("content"===n&&(u-=S.css(e,"padding"+ne[a],!0,i)),"margin"!==n&&(u-=S.css(e,"border"+ne[a]+"Width",!0,i))):(u+=S.css(e,"padding"+ne[a],!0,i),"padding"!==n?u+=S.css(e,"border"+ne[a]+"Width",!0,i):s+=S.css(e,"border"+ne[a]+"Width",!0,i));return!r&&0<=o&&(u+=Math.max(0,Math.ceil(e["offset"+t[0].toUpperCase()+t.slice(1)]-o-u-s-.5))||0),u}function Ze(e,t,n){var r=Ie(e),i=(!y.boxSizingReliable()||n)&&"border-box"===S.css(e,"boxSizing",!1,r),o=i,a=Be(e,t,r),s="offset"+t[0].toUpperCase()+t.slice(1);if(Me.test(a)){if(!n)return a;a="auto"}return(!y.boxSizingReliable()&&i||!y.reliableTrDimensions()&&A(e,"tr")||"auto"===a||!parseFloat(a)&&"inline"===S.css(e,"display",!1,r))&&e.getClientRects().length&&(i="border-box"===S.css(e,"boxSizing",!1,r),(o=s in e)&&(a=e[s])),(a=parseFloat(a)||0)+Ke(e,t,n||(i?"border":"content"),o,r,a)+"px"}function et(e,t,n,r,i){return new et.prototype.init(e,t,n,r,i)}S.extend({cssHooks:{opacity:{get:function(e,t){if(t){var n=Be(e,"opacity");return""===n?"1":n}}}},cssNumber:{animationIterationCount:!0,columnCount:!0,fillOpacity:!0,flexGrow:!0,flexShrink:!0,fontWeight:!0,gridArea:!0,gridColumn:!0,gridColumnEnd:!0,gridColumnStart:!0,gridRow:!0,gridRowEnd:!0,gridRowStart:!0,lineHeight:!0,opacity:!0,order:!0,orphans:!0,widows:!0,zIndex:!0,zoom:!0},cssProps:{},style:function(e,t,n,r){if(e&&3!==e.nodeType&&8!==e.nodeType&&e.style){var i,o,a,s=X(t),u=Ge.test(t),l=e.style;if(u||(t=Xe(s)),a=S.cssHooks[t]||S.cssHooks[s],void 0===n)return a&&"get"in a&&void 0!==(i=a.get(e,!1,r))?i:l[t];"string"===(o=typeof n)&&(i=te.exec(n))&&i[1]&&(n=se(e,t,i),o="number"),null!=n&&n==n&&("number"!==o||u||(n+=i&&i[3]||(S.cssNumber[s]?"":"px")),y.clearCloneStyle||""!==n||0!==t.indexOf("background")||(l[t]="inherit"),a&&"set"in a&&void 0===(n=a.set(e,n,r))||(u?l.setProperty(t,n):l[t]=n))}},css:function(e,t,n,r){var i,o,a,s=X(t);return Ge.test(t)||(t=Xe(s)),(a=S.cssHooks[t]||S.cssHooks[s])&&"get"in a&&(i=a.get(e,!0,n)),void 0===i&&(i=Be(e,t,r)),"normal"===i&&t in Qe&&(i=Qe[t]),""===n||n?(o=parseFloat(i),!0===n||isFinite(o)?o||0:i):i}}),S.each(["height","width"],function(e,u){S.cssHooks[u]={get:function(e,t,n){if(t)return!Ve.test(S.css(e,"display"))||e.getClientRects().length&&e.getBoundingClientRect().width?Ze(e,u,n):We(e,Ye,function(){return Ze(e,u,n)})},set:function(e,t,n){var r,i=Ie(e),o=!y.scrollboxSize()&&"absolute"===i.position,a=(o||n)&&"border-box"===S.css(e,"boxSizing",!1,i),s=n?Ke(e,u,n,a,i):0;return a&&o&&(s-=Math.ceil(e["offset"+u[0].toUpperCase()+u.slice(1)]-parseFloat(i[u])-Ke(e,u,"border",!1,i)-.5)),s&&(r=te.exec(t))&&"px"!==(r[3]||"px")&&(e.style[u]=t,t=S.css(e,u)),Je(0,t,s)}}}),S.cssHooks.marginLeft=$e(y.reliableMarginLeft,function(e,t){if(t)return(parseFloat(Be(e,"marginLeft"))||e.getBoundingClientRect().left-We(e,{marginLeft:0},function(){return e.getBoundingClientRect().left}))+"px"}),S.each({margin:"",padding:"",border:"Width"},function(i,o){S.cssHooks[i+o]={expand:function(e){for(var t=0,n={},r="string"==typeof e?e.split(" "):[e];t<4;t++)n[i+ne[t]+o]=r[t]||r[t-2]||r[0];return n}},"margin"!==i&&(S.cssHooks[i+o].set=Je)}),S.fn.extend({css:function(e,t){return $(this,function(e,t,n){var r,i,o={},a=0;if(Array.isArray(t)){for(r=Ie(e),i=t.length;a<i;a++)o[t[a]]=S.css(e,t[a],!1,r);return o}return void 0!==n?S.style(e,t,n):S.css(e,t)},e,t,1<arguments.length)}}),((S.Tween=et).prototype={constructor:et,init:function(e,t,n,r,i,o){this.elem=e,this.prop=n,this.easing=i||S.easing._default,this.options=t,this.start=this.now=this.cur(),this.end=r,this.unit=o||(S.cssNumber[n]?"":"px")},cur:function(){var e=et.propHooks[this.prop];return e&&e.get?e.get(this):et.propHooks._default.get(this)},run:function(e){var t,n=et.propHooks[this.prop];return this.options.duration?this.pos=t=S.easing[this.easing](e,this.options.duration*e,0,1,this.options.duration):this.pos=t=e,this.now=(this.end-this.start)*t+this.start,this.options.step&&this.options.step.call(this.elem,this.now,this),n&&n.set?n.set(this):et.propHooks._default.set(this),this}}).init.prototype=et.prototype,(et.propHooks={_default:{get:function(e){var t;return 1!==e.elem.nodeType||null!=e.elem[e.prop]&&null==e.elem.style[e.prop]?e.elem[e.prop]:(t=S.css(e.elem,e.prop,""))&&"auto"!==t?t:0},set:function(e){S.fx.step[e.prop]?S.fx.step[e.prop](e):1!==e.elem.nodeType||!S.cssHooks[e.prop]&&null==e.elem.style[Xe(e.prop)]?e.elem[e.prop]=e.now:S.style(e.elem,e.prop,e.now+e.unit)}}}).scrollTop=et.propHooks.scrollLeft={set:function(e){e.elem.nodeType&&e.elem.parentNode&&(e.elem[e.prop]=e.now)}},S.easing={linear:function(e){return e},swing:function(e){return.5-Math.cos(e*Math.PI)/2},_default:"swing"},S.fx=et.prototype.init,S.fx.step={};var tt,nt,rt,it,ot=/^(?:toggle|show|hide)$/,at=/queueHooks$/;function st(){nt&&(!1===E.hidden&&C.requestAnimationFrame?C.requestAnimationFrame(st):C.setTimeout(st,S.fx.interval),S.fx.tick())}function ut(){return C.setTimeout(function(){tt=void 0}),tt=Date.now()}function lt(e,t){var n,r=0,i={height:e};for(t=t?1:0;r<4;r+=2-t)i["margin"+(n=ne[r])]=i["padding"+n]=e;return t&&(i.opacity=i.width=e),i}function ct(e,t,n){for(var r,i=(ft.tweeners[t]||[]).concat(ft.tweeners["*"]),o=0,a=i.length;o<a;o++)if(r=i[o].call(n,t,e))return r}function ft(o,e,t){var n,a,r=0,i=ft.prefilters.length,s=S.Deferred().always(function(){delete u.elem}),u=function(){if(a)return!1;for(var e=tt||ut(),t=Math.max(0,l.startTime+l.duration-e),n=1-(t/l.duration||0),r=0,i=l.tweens.length;r<i;r++)l.tweens[r].run(n);return s.notifyWith(o,[l,n,t]),n<1&&i?t:(i||s.notifyWith(o,[l,1,0]),s.resolveWith(o,[l]),!1)},l=s.promise({elem:o,props:S.extend({},e),opts:S.extend(!0,{specialEasing:{},easing:S.easing._default},t),originalProperties:e,originalOptions:t,startTime:tt||ut(),duration:t.duration,tweens:[],createTween:function(e,t){var n=S.Tween(o,l.opts,e,t,l.opts.specialEasing[e]||l.opts.easing);return l.tweens.push(n),n},stop:function(e){var t=0,n=e?l.tweens.length:0;if(a)return this;for(a=!0;t<n;t++)l.tweens[t].run(1);return e?(s.notifyWith(o,[l,1,0]),s.resolveWith(o,[l,e])):s.rejectWith(o,[l,e]),this}}),c=l.props;for(!function(e,t){var n,r,i,o,a;for(n in e)if(i=t[r=X(n)],o=e[n],Array.isArray(o)&&(i=o[1],o=e[n]=o[0]),n!==r&&(e[r]=o,delete e[n]),(a=S.cssHooks[r])&&"expand"in a)for(n in o=a.expand(o),delete e[r],o)n in e||(e[n]=o[n],t[n]=i);else t[r]=i}(c,l.opts.specialEasing);r<i;r++)if(n=ft.prefilters[r].call(l,o,c,l.opts))return m(n.stop)&&(S._queueHooks(l.elem,l.opts.queue).stop=n.stop.bind(n)),n;return S.map(c,ct,l),m(l.opts.start)&&l.opts.start.call(o,l),l.progress(l.opts.progress).done(l.opts.done,l.opts.complete).fail(l.opts.fail).always(l.opts.always),S.fx.timer(S.extend(u,{elem:o,anim:l,queue:l.opts.queue})),l}S.Animation=S.extend(ft,{tweeners:{"*":[function(e,t){var n=this.createTween(e,t);return se(n.elem,e,te.exec(t),n),n}]},tweener:function(e,t){m(e)?(t=e,e=["*"]):e=e.match(P);for(var n,r=0,i=e.length;r<i;r++)n=e[r],ft.tweeners[n]=ft.tweeners[n]||[],ft.tweeners[n].unshift(t)},prefilters:[function(e,t,n){var r,i,o,a,s,u,l,c,f="width"in t||"height"in t,p=this,d={},h=e.style,g=e.nodeType&&ae(e),v=Y.get(e,"fxshow");for(r in n.queue||(null==(a=S._queueHooks(e,"fx")).unqueued&&(a.unqueued=0,s=a.empty.fire,a.empty.fire=function(){a.unqueued||s()}),a.unqueued++,p.always(function(){p.always(function(){a.unqueued--,S.queue(e,"fx").length||a.empty.fire()})})),t)if(i=t[r],ot.test(i)){if(delete t[r],o=o||"toggle"===i,i===(g?"hide":"show")){if("show"!==i||!v||void 0===v[r])continue;g=!0}d[r]=v&&v[r]||S.style(e,r)}if((u=!S.isEmptyObject(t))||!S.isEmptyObject(d))for(r in f&&1===e.nodeType&&(n.overflow=[h.overflow,h.overflowX,h.overflowY],null==(l=v&&v.display)&&(l=Y.get(e,"display")),"none"===(c=S.css(e,"display"))&&(l?c=l:(le([e],!0),l=e.style.display||l,c=S.css(e,"display"),le([e]))),("inline"===c||"inline-block"===c&&null!=l)&&"none"===S.css(e,"float")&&(u||(p.done(function(){h.display=l}),null==l&&(c=h.display,l="none"===c?"":c)),h.display="inline-block")),n.overflow&&(h.overflow="hidden",p.always(function(){h.overflow=n.overflow[0],h.overflowX=n.overflow[1],h.overflowY=n.overflow[2]})),u=!1,d)u||(v?"hidden"in v&&(g=v.hidden):v=Y.access(e,"fxshow",{display:l}),o&&(v.hidden=!g),g&&le([e],!0),p.done(function(){for(r in g||le([e]),Y.remove(e,"fxshow"),d)S.style(e,r,d[r])})),u=ct(g?v[r]:0,r,p),r in v||(v[r]=u.start,g&&(u.end=u.start,u.start=0))}],prefilter:function(e,t){t?ft.prefilters.unshift(e):ft.prefilters.push(e)}}),S.speed=function(e,t,n){var r=e&&"object"==typeof e?S.extend({},e):{complete:n||!n&&t||m(e)&&e,duration:e,easing:n&&t||t&&!m(t)&&t};return S.fx.off?r.duration=0:"number"!=typeof r.duration&&(r.duration in S.fx.speeds?r.duration=S.fx.speeds[r.duration]:r.duration=S.fx.speeds._default),null!=r.queue&&!0!==r.queue||(r.queue="fx"),r.old=r.complete,r.complete=function(){m(r.old)&&r.old.call(this),r.queue&&S.dequeue(this,r.queue)},r},S.fn.extend({fadeTo:function(e,t,n,r){return this.filter(ae).css("opacity",0).show().end().animate({opacity:t},e,n,r)},animate:function(t,e,n,r){var i=S.isEmptyObject(t),o=S.speed(e,n,r),a=function(){var e=ft(this,S.extend({},t),o);(i||Y.get(this,"finish"))&&e.stop(!0)};return a.finish=a,i||!1===o.queue?this.each(a):this.queue(o.queue,a)},stop:function(i,e,o){var a=function(e){var t=e.stop;delete e.stop,t(o)};return"string"!=typeof i&&(o=e,e=i,i=void 0),e&&this.queue(i||"fx",[]),this.each(function(){var e=!0,t=null!=i&&i+"queueHooks",n=S.timers,r=Y.get(this);if(t)r[t]&&r[t].stop&&a(r[t]);else for(t in r)r[t]&&r[t].stop&&at.test(t)&&a(r[t]);for(t=n.length;t--;)n[t].elem!==this||null!=i&&n[t].queue!==i||(n[t].anim.stop(o),e=!1,n.splice(t,1));!e&&o||S.dequeue(this,i)})},finish:function(a){return!1!==a&&(a=a||"fx"),this.each(function(){var e,t=Y.get(this),n=t[a+"queue"],r=t[a+"queueHooks"],i=S.timers,o=n?n.length:0;for(t.finish=!0,S.queue(this,a,[]),r&&r.stop&&r.stop.call(this,!0),e=i.length;e--;)i[e].elem===this&&i[e].queue===a&&(i[e].anim.stop(!0),i.splice(e,1));for(e=0;e<o;e++)n[e]&&n[e].finish&&n[e].finish.call(this);delete t.finish})}}),S.each(["toggle","show","hide"],function(e,r){var i=S.fn[r];S.fn[r]=function(e,t,n){return null==e||"boolean"==typeof e?i.apply(this,arguments):this.animate(lt(r,!0),e,t,n)}}),S.each({slideDown:lt("show"),slideUp:lt("hide"),slideToggle:lt("toggle"),fadeIn:{opacity:"show"},fadeOut:{opacity:"hide"},fadeToggle:{opacity:"toggle"}},function(e,r){S.fn[e]=function(e,t,n){return this.animate(r,e,t,n)}}),S.timers=[],S.fx.tick=function(){var e,t=0,n=S.timers;for(tt=Date.now();t<n.length;t++)(e=n[t])()||n[t]!==e||n.splice(t--,1);n.length||S.fx.stop(),tt=void 0},S.fx.timer=function(e){S.timers.push(e),S.fx.start()},S.fx.interval=13,S.fx.start=function(){nt||(nt=!0,st())},S.fx.stop=function(){nt=null},S.fx.speeds={slow:600,fast:200,_default:400},S.fn.delay=function(r,e){return r=S.fx&&S.fx.speeds[r]||r,e=e||"fx",this.queue(e,function(e,t){var n=C.setTimeout(e,r);t.stop=function(){C.clearTimeout(n)}})},rt=E.createElement("input"),it=E.createElement("select").appendChild(E.createElement("option")),rt.type="checkbox",y.checkOn=""!==rt.value,y.optSelected=it.selected,(rt=E.createElement("input")).value="t",rt.type="radio",y.radioValue="t"===rt.value;var pt,dt=S.expr.attrHandle;S.fn.extend({attr:function(e,t){return $(this,S.attr,e,t,1<arguments.length)},removeAttr:function(e){return this.each(function(){S.removeAttr(this,e)})}}),S.extend({attr:function(e,t,n){var r,i,o=e.nodeType;if(3!==o&&8!==o&&2!==o)return"undefined"==typeof e.getAttribute?S.prop(e,t,n):(1===o&&S.isXMLDoc(e)||(i=S.attrHooks[t.toLowerCase()]||(S.expr.match.bool.test(t)?pt:void 0)),void 0!==n?null===n?void S.removeAttr(e,t):i&&"set"in i&&void 0!==(r=i.set(e,n,t))?r:(e.setAttribute(t,n+""),n):i&&"get"in i&&null!==(r=i.get(e,t))?r:null==(r=S.find.attr(e,t))?void 0:r)},attrHooks:{type:{set:function(e,t){if(!y.radioValue&&"radio"===t&&A(e,"input")){var n=e.value;return e.setAttribute("type",t),n&&(e.value=n),t}}}},removeAttr:function(e,t){var n,r=0,i=t&&t.match(P);if(i&&1===e.nodeType)while(n=i[r++])e.removeAttribute(n)}}),pt={set:function(e,t,n){return!1===t?S.removeAttr(e,n):e.setAttribute(n,n),n}},S.each(S.expr.match.bool.source.match(/\w+/g),function(e,t){var a=dt[t]||S.find.attr;dt[t]=function(e,t,n){var r,i,o=t.toLowerCase();return n||(i=dt[o],dt[o]=r,r=null!=a(e,t,n)?o:null,dt[o]=i),r}});var ht=/^(?:input|select|textarea|button)$/i,gt=/^(?:a|area)$/i;function vt(e){return(e.match(P)||[]).join(" ")}function yt(e){return e.getAttribute&&e.getAttribute("class")||""}function mt(e){return Array.isArray(e)?e:"string"==typeof e&&e.match(P)||[]}S.fn.extend({prop:function(e,t){return $(this,S.prop,e,t,1<arguments.length)},removeProp:function(e){return this.each(function(){delete this[S.propFix[e]||e]})}}),S.extend({prop:function(e,t,n){var r,i,o=e.nodeType;if(3!==o&&8!==o&&2!==o)return 1===o&&S.isXMLDoc(e)||(t=S.propFix[t]||t,i=S.propHooks[t]),void 0!==n?i&&"set"in i&&void 0!==(r=i.set(e,n,t))?r:e[t]=n:i&&"get"in i&&null!==(r=i.get(e,t))?r:e[t]},propHooks:{tabIndex:{get:function(e){var t=S.find.attr(e,"tabindex");return t?parseInt(t,10):ht.test(e.nodeName)||gt.test(e.nodeName)&&e.href?0:-1}}},propFix:{"for":"htmlFor","class":"className"}}),y.optSelected||(S.propHooks.selected={get:function(e){var t=e.parentNode;return t&&t.parentNode&&t.parentNode.selectedIndex,null},set:function(e){var t=e.parentNode;t&&(t.selectedIndex,t.parentNode&&t.parentNode.selectedIndex)}}),S.each(["tabIndex","readOnly","maxLength","cellSpacing","cellPadding","rowSpan","colSpan","useMap","frameBorder","contentEditable"],function(){S.propFix[this.toLowerCase()]=this}),S.fn.extend({addClass:function(t){var e,n,r,i,o,a,s,u=0;if(m(t))return this.each(function(e){S(this).addClass(t.call(this,e,yt(this)))});if((e=mt(t)).length)while(n=this[u++])if(i=yt(n),r=1===n.nodeType&&" "+vt(i)+" "){a=0;while(o=e[a++])r.indexOf(" "+o+" ")<0&&(r+=o+" ");i!==(s=vt(r))&&n.setAttribute("class",s)}return this},removeClass:function(t){var e,n,r,i,o,a,s,u=0;if(m(t))return this.each(function(e){S(this).removeClass(t.call(this,e,yt(this)))});if(!arguments.length)return this.attr("class","");if((e=mt(t)).length)while(n=this[u++])if(i=yt(n),r=1===n.nodeType&&" "+vt(i)+" "){a=0;while(o=e[a++])while(-1<r.indexOf(" "+o+" "))r=r.replace(" "+o+" "," ");i!==(s=vt(r))&&n.setAttribute("class",s)}return this},toggleClass:function(i,t){var o=typeof i,a="string"===o||Array.isArray(i);return"boolean"==typeof t&&a?t?this.addClass(i):this.removeClass(i):m(i)?this.each(function(e){S(this).toggleClass(i.call(this,e,yt(this),t),t)}):this.each(function(){var e,t,n,r;if(a){t=0,n=S(this),r=mt(i);while(e=r[t++])n.hasClass(e)?n.removeClass(e):n.addClass(e)}else void 0!==i&&"boolean"!==o||((e=yt(this))&&Y.set(this,"__className__",e),this.setAttribute&&this.setAttribute("class",e||!1===i?"":Y.get(this,"__className__")||""))})},hasClass:function(e){var t,n,r=0;t=" "+e+" ";while(n=this[r++])if(1===n.nodeType&&-1<(" "+vt(yt(n))+" ").indexOf(t))return!0;return!1}});var xt=/\r/g;S.fn.extend({val:function(n){var r,e,i,t=this[0];return arguments.length?(i=m(n),this.each(function(e){var t;1===this.nodeType&&(null==(t=i?n.call(this,e,S(this).val()):n)?t="":"number"==typeof t?t+="":Array.isArray(t)&&(t=S.map(t,function(e){return null==e?"":e+""})),(r=S.valHooks[this.type]||S.valHooks[this.nodeName.toLowerCase()])&&"set"in r&&void 0!==r.set(this,t,"value")||(this.value=t))})):t?(r=S.valHooks[t.type]||S.valHooks[t.nodeName.toLowerCase()])&&"get"in r&&void 0!==(e=r.get(t,"value"))?e:"string"==typeof(e=t.value)?e.replace(xt,""):null==e?"":e:void 0}}),S.extend({valHooks:{option:{get:function(e){var t=S.find.attr(e,"value");return null!=t?t:vt(S.text(e))}},select:{get:function(e){var t,n,r,i=e.options,o=e.selectedIndex,a="select-one"===e.type,s=a?null:[],u=a?o+1:i.length;for(r=o<0?u:a?o:0;r<u;r++)if(((n=i[r]).selected||r===o)&&!n.disabled&&(!n.parentNode.disabled||!A(n.parentNode,"optgroup"))){if(t=S(n).val(),a)return t;s.push(t)}return s},set:function(e,t){var n,r,i=e.options,o=S.makeArray(t),a=i.length;while(a--)((r=i[a]).selected=-1<S.inArray(S.valHooks.option.get(r),o))&&(n=!0);return n||(e.selectedIndex=-1),o}}}}),S.each(["radio","checkbox"],function(){S.valHooks[this]={set:function(e,t){if(Array.isArray(t))return e.checked=-1<S.inArray(S(e).val(),t)}},y.checkOn||(S.valHooks[this].get=function(e){return null===e.getAttribute("value")?"on":e.value})}),y.focusin="onfocusin"in C;var bt=/^(?:focusinfocus|focusoutblur)$/,wt=function(e){e.stopPropagation()};S.extend(S.event,{trigger:function(e,t,n,r){var i,o,a,s,u,l,c,f,p=[n||E],d=v.call(e,"type")?e.type:e,h=v.call(e,"namespace")?e.namespace.split("."):[];if(o=f=a=n=n||E,3!==n.nodeType&&8!==n.nodeType&&!bt.test(d+S.event.triggered)&&(-1<d.indexOf(".")&&(d=(h=d.split(".")).shift(),h.sort()),u=d.indexOf(":")<0&&"on"+d,(e=e[S.expando]?e:new S.Event(d,"object"==typeof e&&e)).isTrigger=r?2:3,e.namespace=h.join("."),e.rnamespace=e.namespace?new RegExp("(^|\\.)"+h.join("\\.(?:.*\\.|)")+"(\\.|$)"):null,e.result=void 0,e.target||(e.target=n),t=null==t?[e]:S.makeArray(t,[e]),c=S.event.special[d]||{},r||!c.trigger||!1!==c.trigger.apply(n,t))){if(!r&&!c.noBubble&&!x(n)){for(s=c.delegateType||d,bt.test(s+d)||(o=o.parentNode);o;o=o.parentNode)p.push(o),a=o;a===(n.ownerDocument||E)&&p.push(a.defaultView||a.parentWindow||C)}i=0;while((o=p[i++])&&!e.isPropagationStopped())f=o,e.type=1<i?s:c.bindType||d,(l=(Y.get(o,"events")||Object.create(null))[e.type]&&Y.get(o,"handle"))&&l.apply(o,t),(l=u&&o[u])&&l.apply&&V(o)&&(e.result=l.apply(o,t),!1===e.result&&e.preventDefault());return e.type=d,r||e.isDefaultPrevented()||c._default&&!1!==c._default.apply(p.pop(),t)||!V(n)||u&&m(n[d])&&!x(n)&&((a=n[u])&&(n[u]=null),S.event.triggered=d,e.isPropagationStopped()&&f.addEventListener(d,wt),n[d](),e.isPropagationStopped()&&f.removeEventListener(d,wt),S.event.triggered=void 0,a&&(n[u]=a)),e.result}},simulate:function(e,t,n){var r=S.extend(new S.Event,n,{type:e,isSimulated:!0});S.event.trigger(r,null,t)}}),S.fn.extend({trigger:function(e,t){return this.each(function(){S.event.trigger(e,t,this)})},triggerHandler:function(e,t){var n=this[0];if(n)return S.event.trigger(e,t,n,!0)}}),y.focusin||S.each({focus:"focusin",blur:"focusout"},function(n,r){var i=function(e){S.event.simulate(r,e.target,S.event.fix(e))};S.event.special[r]={setup:function(){var e=this.ownerDocument||this.document||this,t=Y.access(e,r);t||e.addEventListener(n,i,!0),Y.access(e,r,(t||0)+1)},teardown:function(){var e=this.ownerDocument||this.document||this,t=Y.access(e,r)-1;t?Y.access(e,r,t):(e.removeEventListener(n,i,!0),Y.remove(e,r))}}});var Tt=C.location,Ct={guid:Date.now()},Et=/\?/;S.parseXML=function(e){var t;if(!e||"string"!=typeof e)return null;try{t=(new C.DOMParser).parseFromString(e,"text/xml")}catch(e){t=void 0}return t&&!t.getElementsByTagName("parsererror").length||S.error("Invalid XML: "+e),t};var St=/\[\]$/,kt=/\r?\n/g,At=/^(?:submit|button|image|reset|file)$/i,Nt=/^(?:input|select|textarea|keygen)/i;function Dt(n,e,r,i){var t;if(Array.isArray(e))S.each(e,function(e,t){r||St.test(n)?i(n,t):Dt(n+"["+("object"==typeof t&&null!=t?e:"")+"]",t,r,i)});else if(r||"object"!==w(e))i(n,e);else for(t in e)Dt(n+"["+t+"]",e[t],r,i)}S.param=function(e,t){var n,r=[],i=function(e,t){var n=m(t)?t():t;r[r.length]=encodeURIComponent(e)+"="+encodeURIComponent(null==n?"":n)};if(null==e)return"";if(Array.isArray(e)||e.jquery&&!S.isPlainObject(e))S.each(e,function(){i(this.name,this.value)});else for(n in e)Dt(n,e[n],t,i);return r.join("&")},S.fn.extend({serialize:function(){return S.param(this.serializeArray())},serializeArray:function(){return this.map(function(){var e=S.prop(this,"elements");return e?S.makeArray(e):this}).filter(function(){var e=this.type;return this.name&&!S(this).is(":disabled")&&Nt.test(this.nodeName)&&!At.test(e)&&(this.checked||!pe.test(e))}).map(function(e,t){var n=S(this).val();return null==n?null:Array.isArray(n)?S.map(n,function(e){return{name:t.name,value:e.replace(kt,"\r\n")}}):{name:t.name,value:n.replace(kt,"\r\n")}}).get()}});var jt=/%20/g,qt=/#.*$/,Lt=/([?&])_=[^&]*/,Ht=/^(.*?):[ \t]*([^\r\n]*)$/gm,Ot=/^(?:GET|HEAD)$/,Pt=/^\/\//,Rt={},Mt={},It="*/".concat("*"),Wt=E.createElement("a");function Ft(o){return function(e,t){"string"!=typeof e&&(t=e,e="*");var n,r=0,i=e.toLowerCase().match(P)||[];if(m(t))while(n=i[r++])"+"===n[0]?(n=n.slice(1)||"*",(o[n]=o[n]||[]).unshift(t)):(o[n]=o[n]||[]).push(t)}}function Bt(t,i,o,a){var s={},u=t===Mt;function l(e){var r;return s[e]=!0,S.each(t[e]||[],function(e,t){var n=t(i,o,a);return"string"!=typeof n||u||s[n]?u?!(r=n):void 0:(i.dataTypes.unshift(n),l(n),!1)}),r}return l(i.dataTypes[0])||!s["*"]&&l("*")}function $t(e,t){var n,r,i=S.ajaxSettings.flatOptions||{};for(n in t)void 0!==t[n]&&((i[n]?e:r||(r={}))[n]=t[n]);return r&&S.extend(!0,e,r),e}Wt.href=Tt.href,S.extend({active:0,lastModified:{},etag:{},ajaxSettings:{url:Tt.href,type:"GET",isLocal:/^(?:about|app|app-storage|.+-extension|file|res|widget):$/.test(Tt.protocol),global:!0,processData:!0,async:!0,contentType:"application/x-www-form-urlencoded; charset=UTF-8",accepts:{"*":It,text:"text/plain",html:"text/html",xml:"application/xml, text/xml",json:"application/json, text/javascript"},contents:{xml:/\bxml\b/,html:/\bhtml/,json:/\bjson\b/},responseFields:{xml:"responseXML",text:"responseText",json:"responseJSON"},converters:{"* text":String,"text html":!0,"text json":JSON.parse,"text xml":S.parseXML},flatOptions:{url:!0,context:!0}},ajaxSetup:function(e,t){return t?$t($t(e,S.ajaxSettings),t):$t(S.ajaxSettings,e)},ajaxPrefilter:Ft(Rt),ajaxTransport:Ft(Mt),ajax:function(e,t){"object"==typeof e&&(t=e,e=void 0),t=t||{};var c,f,p,n,d,r,h,g,i,o,v=S.ajaxSetup({},t),y=v.context||v,m=v.context&&(y.nodeType||y.jquery)?S(y):S.event,x=S.Deferred(),b=S.Callbacks("once memory"),w=v.statusCode||{},a={},s={},u="canceled",T={readyState:0,getResponseHeader:function(e){var t;if(h){if(!n){n={};while(t=Ht.exec(p))n[t[1].toLowerCase()+" "]=(n[t[1].toLowerCase()+" "]||[]).concat(t[2])}t=n[e.toLowerCase()+" "]}return null==t?null:t.join(", ")},getAllResponseHeaders:function(){return h?p:null},setRequestHeader:function(e,t){return null==h&&(e=s[e.toLowerCase()]=s[e.toLowerCase()]||e,a[e]=t),this},overrideMimeType:function(e){return null==h&&(v.mimeType=e),this},statusCode:function(e){var t;if(e)if(h)T.always(e[T.status]);else for(t in e)w[t]=[w[t],e[t]];return this},abort:function(e){var t=e||u;return c&&c.abort(t),l(0,t),this}};if(x.promise(T),v.url=((e||v.url||Tt.href)+"").replace(Pt,Tt.protocol+"//"),v.type=t.method||t.type||v.method||v.type,v.dataTypes=(v.dataType||"*").toLowerCase().match(P)||[""],null==v.crossDomain){r=E.createElement("a");try{r.href=v.url,r.href=r.href,v.crossDomain=Wt.protocol+"//"+Wt.host!=r.protocol+"//"+r.host}catch(e){v.crossDomain=!0}}if(v.data&&v.processData&&"string"!=typeof v.data&&(v.data=S.param(v.data,v.traditional)),Bt(Rt,v,t,T),h)return T;for(i in(g=S.event&&v.global)&&0==S.active++&&S.event.trigger("ajaxStart"),v.type=v.type.toUpperCase(),v.hasContent=!Ot.test(v.type),f=v.url.replace(qt,""),v.hasContent?v.data&&v.processData&&0===(v.contentType||"").indexOf("application/x-www-form-urlencoded")&&(v.data=v.data.replace(jt,"+")):(o=v.url.slice(f.length),v.data&&(v.processData||"string"==typeof v.data)&&(f+=(Et.test(f)?"&":"?")+v.data,delete v.data),!1===v.cache&&(f=f.replace(Lt,"$1"),o=(Et.test(f)?"&":"?")+"_="+Ct.guid+++o),v.url=f+o),v.ifModified&&(S.lastModified[f]&&T.setRequestHeader("If-Modified-Since",S.lastModified[f]),S.etag[f]&&T.setRequestHeader("If-None-Match",S.etag[f])),(v.data&&v.hasContent&&!1!==v.contentType||t.contentType)&&T.setRequestHeader("Content-Type",v.contentType),T.setRequestHeader("Accept",v.dataTypes[0]&&v.accepts[v.dataTypes[0]]?v.accepts[v.dataTypes[0]]+("*"!==v.dataTypes[0]?", "+It+"; q=0.01":""):v.accepts["*"]),v.headers)T.setRequestHeader(i,v.headers[i]);if(v.beforeSend&&(!1===v.beforeSend.call(y,T,v)||h))return T.abort();if(u="abort",b.add(v.complete),T.done(v.success),T.fail(v.error),c=Bt(Mt,v,t,T)){if(T.readyState=1,g&&m.trigger("ajaxSend",[T,v]),h)return T;v.async&&0<v.timeout&&(d=C.setTimeout(function(){T.abort("timeout")},v.timeout));try{h=!1,c.send(a,l)}catch(e){if(h)throw e;l(-1,e)}}else l(-1,"No Transport");function l(e,t,n,r){var i,o,a,s,u,l=t;h||(h=!0,d&&C.clearTimeout(d),c=void 0,p=r||"",T.readyState=0<e?4:0,i=200<=e&&e<300||304===e,n&&(s=function(e,t,n){var r,i,o,a,s=e.contents,u=e.dataTypes;while("*"===u[0])u.shift(),void 0===r&&(r=e.mimeType||t.getResponseHeader("Content-Type"));if(r)for(i in s)if(s[i]&&s[i].test(r)){u.unshift(i);break}if(u[0]in n)o=u[0];else{for(i in n){if(!u[0]||e.converters[i+" "+u[0]]){o=i;break}a||(a=i)}o=o||a}if(o)return o!==u[0]&&u.unshift(o),n[o]}(v,T,n)),!i&&-1<S.inArray("script",v.dataTypes)&&(v.converters["text script"]=function(){}),s=function(e,t,n,r){var i,o,a,s,u,l={},c=e.dataTypes.slice();if(c[1])for(a in e.converters)l[a.toLowerCase()]=e.converters[a];o=c.shift();while(o)if(e.responseFields[o]&&(n[e.responseFields[o]]=t),!u&&r&&e.dataFilter&&(t=e.dataFilter(t,e.dataType)),u=o,o=c.shift())if("*"===o)o=u;else if("*"!==u&&u!==o){if(!(a=l[u+" "+o]||l["* "+o]))for(i in l)if((s=i.split(" "))[1]===o&&(a=l[u+" "+s[0]]||l["* "+s[0]])){!0===a?a=l[i]:!0!==l[i]&&(o=s[0],c.unshift(s[1]));break}if(!0!==a)if(a&&e["throws"])t=a(t);else try{t=a(t)}catch(e){return{state:"parsererror",error:a?e:"No conversion from "+u+" to "+o}}}return{state:"success",data:t}}(v,s,T,i),i?(v.ifModified&&((u=T.getResponseHeader("Last-Modified"))&&(S.lastModified[f]=u),(u=T.getResponseHeader("etag"))&&(S.etag[f]=u)),204===e||"HEAD"===v.type?l="nocontent":304===e?l="notmodified":(l=s.state,o=s.data,i=!(a=s.error))):(a=l,!e&&l||(l="error",e<0&&(e=0))),T.status=e,T.statusText=(t||l)+"",i?x.resolveWith(y,[o,l,T]):x.rejectWith(y,[T,l,a]),T.statusCode(w),w=void 0,g&&m.trigger(i?"ajaxSuccess":"ajaxError",[T,v,i?o:a]),b.fireWith(y,[T,l]),g&&(m.trigger("ajaxComplete",[T,v]),--S.active||S.event.trigger("ajaxStop")))}return T},getJSON:function(e,t,n){return S.get(e,t,n,"json")},getScript:function(e,t){return S.get(e,void 0,t,"script")}}),S.each(["get","post"],function(e,i){S[i]=function(e,t,n,r){return m(t)&&(r=r||n,n=t,t=void 0),S.ajax(S.extend({url:e,type:i,dataType:r,data:t,success:n},S.isPlainObject(e)&&e))}}),S.ajaxPrefilter(function(e){var t;for(t in e.headers)"content-type"===t.toLowerCase()&&(e.contentType=e.headers[t]||"")}),S._evalUrl=function(e,t,n){return S.ajax({url:e,type:"GET",dataType:"script",cache:!0,async:!1,global:!1,converters:{"text script":function(){}},dataFilter:function(e){S.globalEval(e,t,n)}})},S.fn.extend({wrapAll:function(e){var t;return this[0]&&(m(e)&&(e=e.call(this[0])),t=S(e,this[0].ownerDocument).eq(0).clone(!0),this[0].parentNode&&t.insertBefore(this[0]),t.map(function(){var e=this;while(e.firstElementChild)e=e.firstElementChild;return e}).append(this)),this},wrapInner:function(n){return m(n)?this.each(function(e){S(this).wrapInner(n.call(this,e))}):this.each(function(){var e=S(this),t=e.contents();t.length?t.wrapAll(n):e.append(n)})},wrap:function(t){var n=m(t);return this.each(function(e){S(this).wrapAll(n?t.call(this,e):t)})},unwrap:function(e){return this.parent(e).not("body").each(function(){S(this).replaceWith(this.childNodes)}),this}}),S.expr.pseudos.hidden=function(e){return!S.expr.pseudos.visible(e)},S.expr.pseudos.visible=function(e){return!!(e.offsetWidth||e.offsetHeight||e.getClientRects().length)},S.ajaxSettings.xhr=function(){try{return new C.XMLHttpRequest}catch(e){}};var _t={0:200,1223:204},zt=S.ajaxSettings.xhr();y.cors=!!zt&&"withCredentials"in zt,y.ajax=zt=!!zt,S.ajaxTransport(function(i){var o,a;if(y.cors||zt&&!i.crossDomain)return{send:function(e,t){var n,r=i.xhr();if(r.open(i.type,i.url,i.async,i.username,i.password),i.xhrFields)for(n in i.xhrFields)r[n]=i.xhrFields[n];for(n in i.mimeType&&r.overrideMimeType&&r.overrideMimeType(i.mimeType),i.crossDomain||e["X-Requested-With"]||(e["X-Requested-With"]="XMLHttpRequest"),e)r.setRequestHeader(n,e[n]);o=function(e){return function(){o&&(o=a=r.onload=r.onerror=r.onabort=r.ontimeout=r.onreadystatechange=null,"abort"===e?r.abort():"error"===e?"number"!=typeof r.status?t(0,"error"):t(r.status,r.statusText):t(_t[r.status]||r.status,r.statusText,"text"!==(r.responseType||"text")||"string"!=typeof r.responseText?{binary:r.response}:{text:r.responseText},r.getAllResponseHeaders()))}},r.onload=o(),a=r.onerror=r.ontimeout=o("error"),void 0!==r.onabort?r.onabort=a:r.onreadystatechange=function(){4===r.readyState&&C.setTimeout(function(){o&&a()})},o=o("abort");try{r.send(i.hasContent&&i.data||null)}catch(e){if(o)throw e}},abort:function(){o&&o()}}}),S.ajaxPrefilter(function(e){e.crossDomain&&(e.contents.script=!1)}),S.ajaxSetup({accepts:{script:"text/javascript, application/javascript, application/ecmascript, application/x-ecmascript"},contents:{script:/\b(?:java|ecma)script\b/},converters:{"text script":function(e){return S.globalEval(e),e}}}),S.ajaxPrefilter("script",function(e){void 0===e.cache&&(e.cache=!1),e.crossDomain&&(e.type="GET")}),S.ajaxTransport("script",function(n){var r,i;if(n.crossDomain||n.scriptAttrs)return{send:function(e,t){r=S("<script>").attr(n.scriptAttrs||{}).prop({charset:n.scriptCharset,src:n.url}).on("load error",i=function(e){r.remove(),i=null,e&&t("error"===e.type?404:200,e.type)}),E.head.appendChild(r[0])},abort:function(){i&&i()}}});var Ut,Xt=[],Vt=/(=)\?(?=&|$)|\?\?/;S.ajaxSetup({jsonp:"callback",jsonpCallback:function(){var e=Xt.pop()||S.expando+"_"+Ct.guid++;return this[e]=!0,e}}),S.ajaxPrefilter("json jsonp",function(e,t,n){var r,i,o,a=!1!==e.jsonp&&(Vt.test(e.url)?"url":"string"==typeof e.data&&0===(e.contentType||"").indexOf("application/x-www-form-urlencoded")&&Vt.test(e.data)&&"data");if(a||"jsonp"===e.dataTypes[0])return r=e.jsonpCallback=m(e.jsonpCallback)?e.jsonpCallback():e.jsonpCallback,a?e[a]=e[a].replace(Vt,"$1"+r):!1!==e.jsonp&&(e.url+=(Et.test(e.url)?"&":"?")+e.jsonp+"="+r),e.converters["script json"]=function(){return o||S.error(r+" was not called"),o[0]},e.dataTypes[0]="json",i=C[r],C[r]=function(){o=arguments},n.always(function(){void 0===i?S(C).removeProp(r):C[r]=i,e[r]&&(e.jsonpCallback=t.jsonpCallback,Xt.push(r)),o&&m(i)&&i(o[0]),o=i=void 0}),"script"}),y.createHTMLDocument=((Ut=E.implementation.createHTMLDocument("").body).innerHTML="<form></form><form></form>",2===Ut.childNodes.length),S.parseHTML=function(e,t,n){return"string"!=typeof e?[]:("boolean"==typeof t&&(n=t,t=!1),t||(y.createHTMLDocument?((r=(t=E.implementation.createHTMLDocument("")).createElement("base")).href=E.location.href,t.head.appendChild(r)):t=E),o=!n&&[],(i=N.exec(e))?[t.createElement(i[1])]:(i=xe([e],t,o),o&&o.length&&S(o).remove(),S.merge([],i.childNodes)));var r,i,o},S.fn.load=function(e,t,n){var r,i,o,a=this,s=e.indexOf(" ");return-1<s&&(r=vt(e.slice(s)),e=e.slice(0,s)),m(t)?(n=t,t=void 0):t&&"object"==typeof t&&(i="POST"),0<a.length&&S.ajax({url:e,type:i||"GET",dataType:"html",data:t}).done(function(e){o=arguments,a.html(r?S("<div>").append(S.parseHTML(e)).find(r):e)}).always(n&&function(e,t){a.each(function(){n.apply(this,o||[e.responseText,t,e])})}),this},S.expr.pseudos.animated=function(t){return S.grep(S.timers,function(e){return t===e.elem}).length},S.offset={setOffset:function(e,t,n){var r,i,o,a,s,u,l=S.css(e,"position"),c=S(e),f={};"static"===l&&(e.style.position="relative"),s=c.offset(),o=S.css(e,"top"),u=S.css(e,"left"),("absolute"===l||"fixed"===l)&&-1<(o+u).indexOf("auto")?(a=(r=c.position()).top,i=r.left):(a=parseFloat(o)||0,i=parseFloat(u)||0),m(t)&&(t=t.call(e,n,S.extend({},s))),null!=t.top&&(f.top=t.top-s.top+a),null!=t.left&&(f.left=t.left-s.left+i),"using"in t?t.using.call(e,f):("number"==typeof f.top&&(f.top+="px"),"number"==typeof f.left&&(f.left+="px"),c.css(f))}},S.fn.extend({offset:function(t){if(arguments.length)return void 0===t?this:this.each(function(e){S.offset.setOffset(this,t,e)});var e,n,r=this[0];return r?r.getClientRects().length?(e=r.getBoundingClientRect(),n=r.ownerDocument.defaultView,{top:e.top+n.pageYOffset,left:e.left+n.pageXOffset}):{top:0,left:0}:void 0},position:function(){if(this[0]){var e,t,n,r=this[0],i={top:0,left:0};if("fixed"===S.css(r,"position"))t=r.getBoundingClientRect();else{t=this.offset(),n=r.ownerDocument,e=r.offsetParent||n.documentElement;while(e&&(e===n.body||e===n.documentElement)&&"static"===S.css(e,"position"))e=e.parentNode;e&&e!==r&&1===e.nodeType&&((i=S(e).offset()).top+=S.css(e,"borderTopWidth",!0),i.left+=S.css(e,"borderLeftWidth",!0))}return{top:t.top-i.top-S.css(r,"marginTop",!0),left:t.left-i.left-S.css(r,"marginLeft",!0)}}},offsetParent:function(){return this.map(function(){var e=this.offsetParent;while(e&&"static"===S.css(e,"position"))e=e.offsetParent;return e||re})}}),S.each({scrollLeft:"pageXOffset",scrollTop:"pageYOffset"},function(t,i){var o="pageYOffset"===i;S.fn[t]=function(e){return $(this,function(e,t,n){var r;if(x(e)?r=e:9===e.nodeType&&(r=e.defaultView),void 0===n)return r?r[i]:e[t];r?r.scrollTo(o?r.pageXOffset:n,o?n:r.pageYOffset):e[t]=n},t,e,arguments.length)}}),S.each(["top","left"],function(e,n){S.cssHooks[n]=$e(y.pixelPosition,function(e,t){if(t)return t=Be(e,n),Me.test(t)?S(e).position()[n]+"px":t})}),S.each({Height:"height",Width:"width"},function(a,s){S.each({padding:"inner"+a,content:s,"":"outer"+a},function(r,o){S.fn[o]=function(e,t){var n=arguments.length&&(r||"boolean"!=typeof e),i=r||(!0===e||!0===t?"margin":"border");return $(this,function(e,t,n){var r;return x(e)?0===o.indexOf("outer")?e["inner"+a]:e.document.documentElement["client"+a]:9===e.nodeType?(r=e.documentElement,Math.max(e.body["scroll"+a],r["scroll"+a],e.body["offset"+a],r["offset"+a],r["client"+a])):void 0===n?S.css(e,t,i):S.style(e,t,n,i)},s,n?e:void 0,n)}})}),S.each(["ajaxStart","ajaxStop","ajaxComplete","ajaxError","ajaxSuccess","ajaxSend"],function(e,t){S.fn[t]=function(e){return this.on(t,e)}}),S.fn.extend({bind:function(e,t,n){return this.on(e,null,t,n)},unbind:function(e,t){return this.off(e,null,t)},delegate:function(e,t,n,r){return this.on(t,e,n,r)},undelegate:function(e,t,n){return 1===arguments.length?this.off(e,"**"):this.off(t,e||"**",n)},hover:function(e,t){return this.mouseenter(e).mouseleave(t||e)}}),S.each("blur focus focusin focusout resize scroll click dblclick mousedown mouseup mousemove mouseover mouseout mouseenter mouseleave change select submit keydown keypress keyup contextmenu".split(" "),function(e,n){S.fn[n]=function(e,t){return 0<arguments.length?this.on(n,null,e,t):this.trigger(n)}});var Gt=/^[\s\uFEFF\xA0]+|[\s\uFEFF\xA0]+$/g;S.proxy=function(e,t){var n,r,i;if("string"==typeof t&&(n=e[t],t=e,e=n),m(e))return r=s.call(arguments,2),(i=function(){return e.apply(t||this,r.concat(s.call(arguments)))}).guid=e.guid=e.guid||S.guid++,i},S.holdReady=function(e){e?S.readyWait++:S.ready(!0)},S.isArray=Array.isArray,S.parseJSON=JSON.parse,S.nodeName=A,S.isFunction=m,S.isWindow=x,S.camelCase=X,S.type=w,S.now=Date.now,S.isNumeric=function(e){var t=S.type(e);return("number"===t||"string"===t)&&!isNaN(e-parseFloat(e))},S.trim=function(e){return null==e?"":(e+"").replace(Gt,"")},"function"==typeof define&&define.amd&&define("jquery",[],function(){return S});var Yt=C.jQuery,Qt=C.$;return S.noConflict=function(e){return C.$===S&&(C.$=Qt),e&&C.jQuery===S&&(C.jQuery=Yt),S},"undefined"==typeof e&&(C.jQuery=C.$=S),S});
/*!
* TableSorter 2.17.8 min - Client-side table sorting with ease!
* Copyright (c) 2007 Christian Bach
*/
!function(h){h.extend({tablesorter:new function(){function d(){var b=arguments[0],a=1<arguments.length?Array.prototype.slice.call(arguments):b;if("undefined"!==typeof console&&"undefined"!==typeof console.log)console[/error/i.test(b)?"error":/warn/i.test(b)?"warn":"log"](a);else alert(a)}function q(b,a){d(b+" ("+((new Date).getTime()-a.getTime())+"ms)")}function p(b){for(var a in b)return!1;return!0}function r(b,a,c){if(!a)return"";var f,e=b.config,l=e.textExtraction||"",d="",d="basic"===l?h(a).attr(e.textAttribute)|| a.textContent||a.innerText||h(a).text()||"":"function"===typeof l?l(a,b,c):"function"===typeof(f=g.getColumnData(b,l,c))?f(a,b,c):a.textContent||a.innerText||h(a).text()||"";return h.trim(d)}function v(b){var a,c,f=b.config,e=f.$tbodies=f.$table.children("tbody:not(."+f.cssInfoBlock+")"),l,x,k,h,m,B,u,s,t,p=0,v="",w=e.length;if(0===w)return f.debug?d("Warning: *Empty table!* Not building a parser cache"):"";f.debug&&(t=new Date,d("Detecting parsers for each column"));a=[];for(c=[];p<w;){l=e[p].rows; if(l[p])for(x=f.columns,k=0;k<x;k++){h=f.$headers.filter('[data-column="'+k+'"]:last');m=g.getColumnData(b,f.headers,k);s=g.getParserById(g.getData(h,m,"extractor"));u=g.getParserById(g.getData(h,m,"sorter"));B="false"===g.getData(h,m,"parser");f.empties[k]=(g.getData(h,m,"empty")||f.emptyTo||(f.emptyToBottom?"bottom":"top")).toLowerCase();f.strings[k]=(g.getData(h,m,"string")||f.stringTo||"max").toLowerCase();B&&(u=g.getParserById("no-parser"));s||(s=!1);if(!u)a:{h=b;m=l;B=-1;u=k;for(var A=void 0, K=g.parsers.length,G=!1,z="",A=!0;""===z&&A;)B++,m[B]?(G=m[B].cells[u],z=r(h,G,u),h.config.debug&&d("Checking if value was empty on row "+B+", column: "+u+': "'+z+'"')):A=!1;for(;0<=--K;)if((A=g.parsers[K])&&"text"!==A.id&&A.is&&A.is(z,h,G)){u=A;break a}u=g.getParserById("text")}f.debug&&(v+="column:"+k+"; extractor:"+s.id+"; parser:"+u.id+"; string:"+f.strings[k]+"; empty: "+f.empties[k]+"\n");c[k]=u;a[k]=s}p+=c.length?w:1}f.debug&&(d(v?v:"No parsers detected"),q("Completed detecting parsers",t)); f.parsers=c;f.extractors=a}function w(b){var a,c,f,e,l,x,k,n,m,p,u,s=b.config,t=s.$table.children("tbody"),v=s.extractors,w=s.parsers;s.cache={};s.totalRows=0;if(!w)return s.debug?d("Warning: *Empty table!* Not building a cache"):"";s.debug&&(n=new Date);s.showProcessing&&g.isProcessing(b,!0);for(l=0;l<t.length;l++)if(u=[],a=s.cache[l]={normalized:[]},!t.eq(l).hasClass(s.cssInfoBlock)){m=t[l]&&t[l].rows.length||0;for(f=0;f<m;++f)if(p={child:[]},x=h(t[l].rows[f]),k=[],x.hasClass(s.cssChildRow)&&0!== f)c=a.normalized.length-1,a.normalized[c][s.columns].$row=a.normalized[c][s.columns].$row.add(x),x.prev().hasClass(s.cssChildRow)||x.prev().addClass(g.css.cssHasChild),p.child[c]=h.trim(x[0].textContent||x[0].innerText||x.text()||"");else{p.$row=x;p.order=f;for(e=0;e<s.columns;++e)"undefined"===typeof w[e]?s.debug&&d("No parser found for cell:",x[0].cells[e],"does it have a header?"):(c=r(b,x[0].cells[e],e),c="undefined"===typeof v[e].id?c:v[e].format(c,b,x[0].cells[e],e),c="no-parser"===w[e].id? "":w[e].format(c,b,x[0].cells[e],e),k.push(s.ignoreCase&&"string"===typeof c?c.toLowerCase():c),"numeric"===(w[e].type||"").toLowerCase()&&(u[e]=Math.max(Math.abs(c)||0,u[e]||0)));k[s.columns]=p;a.normalized.push(k)}a.colMax=u;s.totalRows+=a.normalized.length}s.showProcessing&&g.isProcessing(b);s.debug&&q("Building cache for "+m+" rows",n)}function z(b,a){var c=b.config,f=c.widgetOptions,e=b.tBodies,l=[],d=c.cache,k,n,m,r,u,s;if(p(d))return c.appender?c.appender(b,l):b.isUpdating?c.$table.trigger("updateComplete", b):"";c.debug&&(s=new Date);for(u=0;u<e.length;u++)if(k=h(e[u]),k.length&&!k.hasClass(c.cssInfoBlock)){m=g.processTbody(b,k,!0);k=d[u].normalized;n=k.length;for(r=0;r<n;r++)l.push(k[r][c.columns].$row),c.appender&&(!c.pager||c.pager.removeRows&&f.pager_removeRows||c.pager.ajax)||m.append(k[r][c.columns].$row);g.processTbody(b,m,!1)}c.appender&&c.appender(b,l);c.debug&&q("Rebuilt table",s);a||c.appender||g.applyWidget(b);b.isUpdating&&c.$table.trigger("updateComplete",b)}function D(b){return/^d/i.test(b)|| 1===b}function E(b){var a,c,f,e,l,x,k,n=b.config;n.headerList=[];n.headerContent=[];n.debug&&(k=new Date);n.columns=g.computeColumnIndex(n.$table.children("thead, tfoot").children("tr"));e=n.cssIcon?'<i class="'+(n.cssIcon===g.css.icon?g.css.icon:n.cssIcon+" "+g.css.icon)+'"></i>':"";n.$headers=h(b).find(n.selectorHeaders).each(function(k){c=h(this);a=g.getColumnData(b,n.headers,k,!0);n.headerContent[k]=h(this).html();""!==n.headerTemplate&&(l=n.headerTemplate.replace(/\{content\}/g,h(this).html()).replace(/\{icon\}/g, e),n.onRenderTemplate&&(f=n.onRenderTemplate.apply(c,[k,l]))&&"string"===typeof f&&(l=f),h(this).html('<div class="'+g.css.headerIn+'">'+l+"</div>"));n.onRenderHeader&&n.onRenderHeader.apply(c,[k]);this.column=parseInt(h(this).attr("data-column"),10);this.order=D(g.getData(c,a,"sortInitialOrder")||n.sortInitialOrder)?[1,0,2]:[0,1,2];this.count=-1;this.lockedOrder=!1;x=g.getData(c,a,"lockedOrder")||!1;"undefined"!==typeof x&&!1!==x&&(this.order=this.lockedOrder=D(x)?[1,1,1]:[0,0,0]);c.addClass(g.css.header+ " "+n.cssHeader);n.headerList[k]=this;c.parent().addClass(g.css.headerRow+" "+n.cssHeaderRow).attr("role","row");n.tabIndex&&c.attr("tabindex",0)}).attr({scope:"col",role:"columnheader"});H(b);n.debug&&(q("Built headers:",k),d(n.$headers))}function C(b,a,c){var f=b.config;f.$table.find(f.selectorRemove).remove();v(b);w(b);I(f.$table,a,c)}function H(b){var a,c,f,e=b.config;e.$headers.each(function(l,d){c=h(d);f=g.getColumnData(b,e.headers,l,!0);a="false"===g.getData(d,f,"sorter")||"false"===g.getData(d, f,"parser");d.sortDisabled=a;c[a?"addClass":"removeClass"]("sorter-false").attr("aria-disabled",""+a);b.id&&(a?c.removeAttr("aria-controls"):c.attr("aria-controls",b.id))})}function F(b){var a,c,f=b.config,e=f.sortList,l=e.length,d=g.css.sortNone+" "+f.cssNone,k=[g.css.sortAsc+" "+f.cssAsc,g.css.sortDesc+" "+f.cssDesc],n=["ascending","descending"],m=h(b).find("tfoot tr").children().add(f.$extraHeaders).removeClass(k.join(" "));f.$headers.removeClass(k.join(" ")).addClass(d).attr("aria-sort","none"); for(a=0;a<l;a++)if(2!==e[a][1]&&(b=f.$headers.not(".sorter-false").filter('[data-column="'+e[a][0]+'"]'+(1===l?":last":"")),b.length)){for(c=0;c<b.length;c++)b[c].sortDisabled||b.eq(c).removeClass(d).addClass(k[e[a][1]]).attr("aria-sort",n[e[a][1]]);m.length&&m.filter('[data-column="'+e[a][0]+'"]').removeClass(d).addClass(k[e[a][1]])}f.$headers.not(".sorter-false").each(function(){var b=h(this),a=this.order[(this.count+1)%(f.sortReset?3:2)],a=b.text()+": "+g.language[b.hasClass(g.css.sortAsc)?"sortAsc": b.hasClass(g.css.sortDesc)?"sortDesc":"sortNone"]+g.language[0===a?"nextAsc":1===a?"nextDesc":"nextNone"];b.attr("aria-label",a)})}function O(b){var a,c,f=b.config;f.widthFixed&&0===f.$table.find("colgroup").length&&(a=h("<colgroup>"),c=h(b).width(),h(b.tBodies).not("."+f.cssInfoBlock).find("tr:first").children(":visible").each(function(){a.append(h("<col>").css("width",parseInt(h(this).width()/c*1E3,10)/10+"%"))}),f.$table.prepend(a))}function P(b,a){var c,f,e,l,g,k=b.config,d=a||k.sortList;k.sortList= [];h.each(d,function(b,a){l=parseInt(a[0],10);if(e=k.$headers.filter('[data-column="'+l+'"]:last')[0]){f=(f=(""+a[1]).match(/^(1|d|s|o|n)/))?f[0]:"";switch(f){case "1":case "d":f=1;break;case "s":f=g||0;break;case "o":c=e.order[(g||0)%(k.sortReset?3:2)];f=0===c?1:1===c?0:2;break;case "n":e.count+=1;f=e.order[e.count%(k.sortReset?3:2)];break;default:f=0}g=0===b?f:g;c=[l,parseInt(f,10)||0];k.sortList.push(c);f=h.inArray(c[1],e.order);e.count=0<=f?f:c[1]%(k.sortReset?3:2)}})}function Q(b,a){return b&& b[a]?b[a].type||"":""}function L(b,a,c){if(b.isUpdating)return setTimeout(function(){L(b,a,c)},50);var f,e,l,d,k=b.config,n=!c[k.sortMultiSortKey],m=k.$table;m.trigger("sortStart",b);a.count=c[k.sortResetKey]?2:(a.count+1)%(k.sortReset?3:2);k.sortRestart&&(e=a,k.$headers.each(function(){this===e||!n&&h(this).is("."+g.css.sortDesc+",."+g.css.sortAsc)||(this.count=-1)}));e=a.column;if(n){k.sortList=[];if(null!==k.sortForce)for(f=k.sortForce,l=0;l<f.length;l++)f[l][0]!==e&&k.sortList.push(f[l]);f=a.order[a.count]; if(2>f&&(k.sortList.push([e,f]),1<a.colSpan))for(l=1;l<a.colSpan;l++)k.sortList.push([e+l,f])}else{if(k.sortAppend&&1<k.sortList.length)for(l=0;l<k.sortAppend.length;l++)d=g.isValueInArray(k.sortAppend[l][0],k.sortList),0<=d&&k.sortList.splice(d,1);if(0<=g.isValueInArray(e,k.sortList))for(l=0;l<k.sortList.length;l++)d=k.sortList[l],f=k.$headers.filter('[data-column="'+d[0]+'"]:last')[0],d[0]===e&&(d[1]=f.order[a.count],2===d[1]&&(k.sortList.splice(l,1),f.count=-1));else if(f=a.order[a.count],2>f&& (k.sortList.push([e,f]),1<a.colSpan))for(l=1;l<a.colSpan;l++)k.sortList.push([e+l,f])}if(null!==k.sortAppend)for(f=k.sortAppend,l=0;l<f.length;l++)f[l][0]!==e&&k.sortList.push(f[l]);m.trigger("sortBegin",b);setTimeout(function(){F(b);J(b);z(b);m.trigger("sortEnd",b)},1)}function J(b){var a,c,f,e,l,d,k,h,m,r,u,s=0,t=b.config,v=t.textSorter||"",w=t.sortList,y=w.length,z=b.tBodies.length;if(!t.serverSideSorting&&!p(t.cache)){t.debug&&(l=new Date);for(c=0;c<z;c++)d=t.cache[c].colMax,k=t.cache[c].normalized, k.sort(function(c,l){for(a=0;a<y;a++){e=w[a][0];h=w[a][1];s=0===h;if(t.sortStable&&c[e]===l[e]&&1===y)break;(f=/n/i.test(Q(t.parsers,e)))&&t.strings[e]?(f="boolean"===typeof t.string[t.strings[e]]?(s?1:-1)*(t.string[t.strings[e]]?-1:1):t.strings[e]?t.string[t.strings[e]]||0:0,m=t.numberSorter?t.numberSorter(c[e],l[e],s,d[e],b):g["sortNumeric"+(s?"Asc":"Desc")](c[e],l[e],f,d[e],e,b)):(r=s?c:l,u=s?l:c,m="function"===typeof v?v(r[e],u[e],s,e,b):"object"===typeof v&&v.hasOwnProperty(e)?v[e](r[e],u[e], s,e,b):g["sortNatural"+(s?"Asc":"Desc")](c[e],l[e],e,b,t));if(m)return m}return c[t.columns].order-l[t.columns].order});t.debug&&q("Sorting on "+w.toString()+" and dir "+h+" time",l)}}function M(b,a){var c=b[0];c.isUpdating&&b.trigger("updateComplete",c);h.isFunction(a)&&a(b[0])}function I(b,a,c){var f=b[0].config.sortList;!1!==a&&!b[0].isProcessing&&f.length?b.trigger("sorton",[f,function(){M(b,c)},!0]):(M(b,c),g.applyWidget(b[0],!1))}function N(b){var a=b.config,c=a.$table;c.unbind("sortReset update updateRows updateCell updateAll addRows updateComplete sorton appendCache updateCache applyWidgetId applyWidgets refreshWidgets destroy mouseup mouseleave ".split(" ").join(a.namespace+ " ")).bind("sortReset"+a.namespace,function(f,e){f.stopPropagation();a.sortList=[];F(b);J(b);z(b);h.isFunction(e)&&e(b)}).bind("updateAll"+a.namespace,function(f,e,c){f.stopPropagation();b.isUpdating=!0;g.refreshWidgets(b,!0,!0);g.restoreHeaders(b);E(b);g.bindEvents(b,a.$headers,!0);N(b);C(b,e,c)}).bind("update"+a.namespace+" updateRows"+a.namespace,function(a,e,c){a.stopPropagation();b.isUpdating=!0;H(b);C(b,e,c)}).bind("updateCell"+a.namespace,function(f,e,l,g){f.stopPropagation();b.isUpdating= !0;c.find(a.selectorRemove).remove();var d,n,m;n=c.find("tbody");m=h(e);f=n.index(h.fn.closest?m.closest("tbody"):m.parents("tbody").filter(":first"));d=h.fn.closest?m.closest("tr"):m.parents("tr").filter(":first");e=m[0];n.length&&0<=f&&(n=n.eq(f).find("tr").index(d),m=m.index(),a.cache[f].normalized[n][a.columns].$row=d,d="undefined"===typeof a.extractors[m].id?r(b,e,m):a.extractors[m].format(r(b,e,m),b,e,m),e="no-parser"===a.parsers[m].id?"":a.parsers[m].format(d,b,e,m),a.cache[f].normalized[n][m]= a.ignoreCase&&"string"===typeof e?e.toLowerCase():e,"numeric"===(a.parsers[m].type||"").toLowerCase()&&(a.cache[f].colMax[m]=Math.max(Math.abs(e)||0,a.cache[f].colMax[m]||0)),I(c,l,g))}).bind("addRows"+a.namespace,function(f,e,l,g){f.stopPropagation();b.isUpdating=!0;if(p(a.cache))H(b),C(b,l,g);else{e=h(e).attr("role","row");var d,n,m,q,u,s=e.filter("tr").length,t=c.find("tbody").index(e.parents("tbody").filter(":first"));a.parsers&&a.parsers.length||v(b);for(f=0;f<s;f++){n=e[f].cells.length;u=[]; q={child:[],$row:e.eq(f),order:a.cache[t].normalized.length};for(d=0;d<n;d++)m="undefined"===typeof a.extractors[d].id?r(b,e[f].cells[d],d):a.extractors[d].format(r(b,e[f].cells[d],d),b,e[f].cells[d],d),m="no-parser"===a.parsers[d].id?"":a.parsers[d].format(m,b,e[f].cells[d],d),u[d]=a.ignoreCase&&"string"===typeof m?m.toLowerCase():m,"numeric"===(a.parsers[d].type||"").toLowerCase()&&(a.cache[t].colMax[d]=Math.max(Math.abs(u[d])||0,a.cache[t].colMax[d]||0));u.push(q);a.cache[t].normalized.push(u)}I(c, l,g)}}).bind("updateComplete"+a.namespace,function(){b.isUpdating=!1}).bind("sorton"+a.namespace,function(a,e,d,x){var k=b.config;a.stopPropagation();c.trigger("sortStart",this);P(b,e);F(b);k.delayInit&&p(k.cache)&&w(b);c.trigger("sortBegin",this);J(b);z(b,x);c.trigger("sortEnd",this);g.applyWidget(b);h.isFunction(d)&&d(b)}).bind("appendCache"+a.namespace,function(a,e,c){a.stopPropagation();z(b,c);h.isFunction(e)&&e(b)}).bind("updateCache"+a.namespace,function(c,e){a.parsers&&a.parsers.length||v(b); w(b);h.isFunction(e)&&e(b)}).bind("applyWidgetId"+a.namespace,function(c,e){c.stopPropagation();g.getWidgetById(e).format(b,a,a.widgetOptions)}).bind("applyWidgets"+a.namespace,function(a,c){a.stopPropagation();g.applyWidget(b,c)}).bind("refreshWidgets"+a.namespace,function(a,c,d){a.stopPropagation();g.refreshWidgets(b,c,d)}).bind("destroy"+a.namespace,function(a,c,d){a.stopPropagation();g.destroy(b,c,d)}).bind("resetToLoadState"+a.namespace,function(){g.refreshWidgets(b,!0,!0);a=h.extend(!0,g.defaults, a.originalSettings);b.hasInitialized=!1;g.setup(b,a)})}var g=this;g.version="2.17.8";g.parsers=[];g.widgets=[];g.defaults={theme:"default",widthFixed:!1,showProcessing:!1,headerTemplate:"{content}",onRenderTemplate:null,onRenderHeader:null,cancelSelection:!0,tabIndex:!0,dateFormat:"mmddyyyy",sortMultiSortKey:"shiftKey",sortResetKey:"ctrlKey",usNumberFormat:!0,delayInit:!1,serverSideSorting:!1,headers:{},ignoreCase:!0,sortForce:null,sortList:[],sortAppend:null,sortStable:!1,sortInitialOrder:"asc", sortLocaleCompare:!1,sortReset:!1,sortRestart:!1,emptyTo:"bottom",stringTo:"max",textExtraction:"basic",textAttribute:"data-text",textSorter:null,numberSorter:null,widgets:[],widgetOptions:{zebra:["even","odd"]},initWidgets:!0,initialized:null,tableClass:"",cssAsc:"",cssDesc:"",cssNone:"",cssHeader:"",cssHeaderRow:"",cssProcessing:"",cssChildRow:"tablesorter-childRow",cssIcon:"tablesorter-icon",cssInfoBlock:"tablesorter-infoOnly",selectorHeaders:"> thead th, > thead td",selectorSort:"th, td",selectorRemove:".remove-me", debug:!1,headerList:[],empties:{},strings:{},parsers:[]};g.css={table:"tablesorter",cssHasChild:"tablesorter-hasChildRow",childRow:"tablesorter-childRow",header:"tablesorter-header",headerRow:"tablesorter-headerRow",headerIn:"tablesorter-header-inner",icon:"tablesorter-icon",info:"tablesorter-infoOnly",processing:"tablesorter-processing",sortAsc:"tablesorter-headerAsc",sortDesc:"tablesorter-headerDesc",sortNone:"tablesorter-headerUnSorted"};g.language={sortAsc:"Ascending sort applied, ",sortDesc:"Descending sort applied, ", sortNone:"No sort applied, ",nextAsc:"activate to apply an ascending sort",nextDesc:"activate to apply a descending sort",nextNone:"activate to remove the sort"};g.log=d;g.benchmark=q;g.construct=function(b){return this.each(function(){var a=h.extend(!0,{},g.defaults,b);a.originalSettings=b;!this.hasInitialized&&g.buildTable&&"TABLE"!==this.tagName?g.buildTable(this,a):g.setup(this,a)})};g.setup=function(b,a){if(!b||!b.tHead||0===b.tBodies.length||!0===b.hasInitialized)return a.debug?d("ERROR: stopping initialization! No table, thead, tbody or tablesorter has already been initialized"): "";var c="",f=h(b),e=h.metadata;b.hasInitialized=!1;b.isProcessing=!0;b.config=a;h.data(b,"tablesorter",a);a.debug&&h.data(b,"startoveralltimer",new Date);a.supportsDataObject=function(a){a[0]=parseInt(a[0],10);return 1<a[0]||1===a[0]&&4<=parseInt(a[1],10)}(h.fn.jquery.split("."));a.string={max:1,min:-1,emptymin:1,emptymax:-1,zero:0,none:0,"null":0,top:!0,bottom:!1};a.emptyTo=a.emptyTo.toLowerCase();a.stringTo=a.stringTo.toLowerCase();/tablesorter\-/.test(f.attr("class"))||(c=""!==a.theme?" tablesorter-"+ a.theme:"");a.table=b;a.$table=f.addClass(g.css.table+" "+a.tableClass+c).attr("role","grid");a.$headers=f.find(a.selectorHeaders);a.namespace=a.namespace?"."+a.namespace.replace(/\W/g,""):".tablesorter"+Math.random().toString(16).slice(2);a.$table.children().children("tr").attr("role","row");a.$tbodies=f.children("tbody:not(."+a.cssInfoBlock+")").attr({"aria-live":"polite","aria-relevant":"all"});a.$table.find("caption").length&&a.$table.attr("aria-labelledby","theCaption");a.widgetInit={};a.textExtraction= a.$table.attr("data-text-extraction")||a.textExtraction||"basic";E(b);O(b);v(b);a.totalRows=0;a.delayInit||w(b);g.bindEvents(b,a.$headers,!0);N(b);a.supportsDataObject&&"undefined"!==typeof f.data().sortlist?a.sortList=f.data().sortlist:e&&f.metadata()&&f.metadata().sortlist&&(a.sortList=f.metadata().sortlist);g.applyWidget(b,!0);0<a.sortList.length?f.trigger("sorton",[a.sortList,{},!a.initWidgets,!0]):(F(b),a.initWidgets&&g.applyWidget(b,!1));a.showProcessing&&f.unbind("sortBegin"+a.namespace+" sortEnd"+ a.namespace).bind("sortBegin"+a.namespace+" sortEnd"+a.namespace,function(c){clearTimeout(a.processTimer);g.isProcessing(b);"sortBegin"===c.type&&(a.processTimer=setTimeout(function(){g.isProcessing(b,!0)},500))});b.hasInitialized=!0;b.isProcessing=!1;a.debug&&g.benchmark("Overall initialization time",h.data(b,"startoveralltimer"));f.trigger("tablesorter-initialized",b);"function"===typeof a.initialized&&a.initialized(b)};g.getColumnData=function(b,a,c,f){if("undefined"!==typeof a&&null!==a){b=h(b)[0]; var e,d=b.config;if(a[c])return f?a[c]:a[d.$headers.index(d.$headers.filter('[data-column="'+c+'"]:last'))];for(e in a)if("string"===typeof e&&(b=f?d.$headers.eq(c).filter(e):d.$headers.filter('[data-column="'+c+'"]:last').filter(e),b.length))return a[e]}};g.computeColumnIndex=function(b){var a=[],c=0,f,e,d,g,k,n,m,p,q,s;for(f=0;f<b.length;f++)for(k=b[f].cells,e=0;e<k.length;e++){d=k[e];g=h(d);n=d.parentNode.rowIndex;g.index();m=d.rowSpan||1;p=d.colSpan||1;"undefined"===typeof a[n]&&(a[n]=[]);for(d= 0;d<a[n].length+1;d++)if("undefined"===typeof a[n][d]){q=d;break}c=Math.max(q,c);g.attr({"data-column":q});for(d=n;d<n+m;d++)for("undefined"===typeof a[d]&&(a[d]=[]),s=a[d],g=q;g<q+p;g++)s[g]="x"}return c+1};g.isProcessing=function(b,a,c){b=h(b);var f=b[0].config,e=c||b.find("."+g.css.header);a?("undefined"!==typeof c&&0<f.sortList.length&&(e=e.filter(function(){return this.sortDisabled?!1:0<=g.isValueInArray(parseFloat(h(this).attr("data-column")),f.sortList)})),b.add(e).addClass(g.css.processing+ " "+f.cssProcessing)):b.add(e).removeClass(g.css.processing+" "+f.cssProcessing)};g.processTbody=function(b,a,c){b=h(b)[0];if(c)return b.isProcessing=!0,a.before('<span class="tablesorter-savemyplace"/>'),c=h.fn.detach?a.detach():a.remove();c=h(b).find("span.tablesorter-savemyplace");a.insertAfter(c);c.remove();b.isProcessing=!1};g.clearTableBody=function(b){h(b)[0].config.$tbodies.children().detach()};g.bindEvents=function(b,a,c){b=h(b)[0];var f,e=b.config;!0!==c&&(e.$extraHeaders=e.$extraHeaders? e.$extraHeaders.add(a):a);a.find(e.selectorSort).add(a.filter(e.selectorSort)).unbind(["mousedown","mouseup","sort","keyup",""].join(e.namespace+" ")).bind(["mousedown","mouseup","sort","keyup",""].join(e.namespace+" "),function(c,d){var g;g=c.type;if(!(1!==(c.which||c.button)&&!/sort|keyup/.test(g)||"keyup"===g&&13!==c.which||"mouseup"===g&&!0!==d&&250<(new Date).getTime()-f)){if("mousedown"===g)return f=(new Date).getTime(),/(input|select|button|textarea)/i.test(c.target.tagName)?"":!e.cancelSelection; e.delayInit&&p(e.cache)&&w(b);g=h.fn.closest?h(this).closest("th, td")[0]:/TH|TD/.test(this.tagName)?this:h(this).parents("th, td")[0];g=e.$headers[a.index(g)];g.sortDisabled||L(b,g,c)}});e.cancelSelection&&a.attr("unselectable","on").bind("selectstart",!1).css({"user-select":"none",MozUserSelect:"none"})};g.restoreHeaders=function(b){var a=h(b)[0].config;a.$table.find(a.selectorHeaders).each(function(b){h(this).find("."+g.css.headerIn).length&&h(this).html(a.headerContent[b])})};g.destroy=function(b, a,c){b=h(b)[0];if(b.hasInitialized){g.refreshWidgets(b,!0,!0);var f=h(b),e=b.config,d=f.find("thead:first"),q=d.find("tr."+g.css.headerRow).removeClass(g.css.headerRow+" "+e.cssHeaderRow),k=f.find("tfoot:first > tr").children("th, td");!1===a&&0<=h.inArray("uitheme",e.widgets)&&(f.trigger("applyWidgetId",["uitheme"]),f.trigger("applyWidgetId",["zebra"]));d.find("tr").not(q).remove();f.removeData("tablesorter").unbind("sortReset update updateAll updateRows updateCell addRows updateComplete sorton appendCache updateCache applyWidgetId applyWidgets refreshWidgets destroy mouseup mouseleave keypress sortBegin sortEnd resetToLoadState ".split(" ").join(e.namespace+ " "));e.$headers.add(k).removeClass([g.css.header,e.cssHeader,e.cssAsc,e.cssDesc,g.css.sortAsc,g.css.sortDesc,g.css.sortNone].join(" ")).removeAttr("data-column").removeAttr("aria-label").attr("aria-disabled","true");q.find(e.selectorSort).unbind(["mousedown","mouseup","keypress",""].join(e.namespace+" "));g.restoreHeaders(b);f.toggleClass(g.css.table+" "+e.tableClass+" tablesorter-"+e.theme,!1===a);b.hasInitialized=!1;delete b.config.cache;"function"===typeof c&&c(b)}};g.regex={chunk:/(^([+\-]?(?:0|[1-9]\d*)(?:\.\d*)?(?:[eE][+\-]?\d+)?)?$|^0x[0-9a-f]+$|\d+)/gi, chunks:/(^\\0|\\0$)/,hex:/^0x[0-9a-f]+$/i};g.sortNatural=function(b,a){if(b===a)return 0;var c,f,e,d,h,k;f=g.regex;if(f.hex.test(a)){c=parseInt(b.match(f.hex),16);e=parseInt(a.match(f.hex),16);if(c<e)return-1;if(c>e)return 1}c=b.replace(f.chunk,"\\0$1\\0").replace(f.chunks,"").split("\\0");f=a.replace(f.chunk,"\\0$1\\0").replace(f.chunks,"").split("\\0");k=Math.max(c.length,f.length);for(h=0;h<k;h++){e=isNaN(c[h])?c[h]||0:parseFloat(c[h])||0;d=isNaN(f[h])?f[h]||0:parseFloat(f[h])||0;if(isNaN(e)!== isNaN(d))return isNaN(e)?1:-1;typeof e!==typeof d&&(e+="",d+="");if(e<d)return-1;if(e>d)return 1}return 0};g.sortNaturalAsc=function(b,a,c,f,e){if(b===a)return 0;c=e.string[e.empties[c]||e.emptyTo];return""===b&&0!==c?"boolean"===typeof c?c?-1:1:-c||-1:""===a&&0!==c?"boolean"===typeof c?c?1:-1:c||1:g.sortNatural(b,a)};g.sortNaturalDesc=function(b,a,c,f,e){if(b===a)return 0;c=e.string[e.empties[c]||e.emptyTo];return""===b&&0!==c?"boolean"===typeof c?c?-1:1:c||1:""===a&&0!==c?"boolean"===typeof c?c? 1:-1:-c||-1:g.sortNatural(a,b)};g.sortText=function(b,a){return b>a?1:b<a?-1:0};g.getTextValue=function(b,a,c){if(c){var f=b?b.length:0,e=c+a;for(c=0;c<f;c++)e+=b.charCodeAt(c);return a*e}return 0};g.sortNumericAsc=function(b,a,c,f,e,d){if(b===a)return 0;d=d.config;e=d.string[d.empties[e]||d.emptyTo];if(""===b&&0!==e)return"boolean"===typeof e?e?-1:1:-e||-1;if(""===a&&0!==e)return"boolean"===typeof e?e?1:-1:e||1;isNaN(b)&&(b=g.getTextValue(b,c,f));isNaN(a)&&(a=g.getTextValue(a,c,f));return b-a};g.sortNumericDesc= function(b,a,c,f,e,d){if(b===a)return 0;d=d.config;e=d.string[d.empties[e]||d.emptyTo];if(""===b&&0!==e)return"boolean"===typeof e?e?-1:1:e||1;if(""===a&&0!==e)return"boolean"===typeof e?e?1:-1:-e||-1;isNaN(b)&&(b=g.getTextValue(b,c,f));isNaN(a)&&(a=g.getTextValue(a,c,f));return a-b};g.sortNumeric=function(b,a){return b-a};g.characterEquivalents={a:"\u00e1\u00e0\u00e2\u00e3\u00e4\u0105\u00e5",A:"\u00c1\u00c0\u00c2\u00c3\u00c4\u0104\u00c5",c:"\u00e7\u0107\u010d",C:"\u00c7\u0106\u010c",e:"\u00e9\u00e8\u00ea\u00eb\u011b\u0119", E:"\u00c9\u00c8\u00ca\u00cb\u011a\u0118",i:"\u00ed\u00ec\u0130\u00ee\u00ef\u0131",I:"\u00cd\u00cc\u0130\u00ce\u00cf",o:"\u00f3\u00f2\u00f4\u00f5\u00f6",O:"\u00d3\u00d2\u00d4\u00d5\u00d6",ss:"\u00df",SS:"\u1e9e",u:"\u00fa\u00f9\u00fb\u00fc\u016f",U:"\u00da\u00d9\u00db\u00dc\u016e"};g.replaceAccents=function(b){var a,c="[",d=g.characterEquivalents;if(!g.characterRegex){g.characterRegexArray={};for(a in d)"string"===typeof a&&(c+=d[a],g.characterRegexArray[a]=new RegExp("["+d[a]+"]","g"));g.characterRegex= new RegExp(c+"]")}if(g.characterRegex.test(b))for(a in d)"string"===typeof a&&(b=b.replace(g.characterRegexArray[a],a));return b};g.isValueInArray=function(b,a){var c,d=a.length;for(c=0;c<d;c++)if(a[c][0]===b)return c;return-1};g.addParser=function(b){var a,c=g.parsers.length,d=!0;for(a=0;a<c;a++)g.parsers[a].id.toLowerCase()===b.id.toLowerCase()&&(d=!1);d&&g.parsers.push(b)};g.getParserById=function(b){if("false"==b)return!1;var a,c=g.parsers.length;for(a=0;a<c;a++)if(g.parsers[a].id.toLowerCase()=== b.toString().toLowerCase())return g.parsers[a];return!1};g.addWidget=function(b){g.widgets.push(b)};g.hasWidget=function(b,a){b=h(b);return b.length&&b[0].config&&b[0].config.widgetInit[a]||!1};g.getWidgetById=function(b){var a,c,d=g.widgets.length;for(a=0;a<d;a++)if((c=g.widgets[a])&&c.hasOwnProperty("id")&&c.id.toLowerCase()===b.toLowerCase())return c};g.applyWidget=function(b,a){b=h(b)[0];var c=b.config,d=c.widgetOptions,e=[],l,p,k;!1!==a&&b.hasInitialized&&(b.isApplyingWidgets||b.isUpdating)|| (c.debug&&(l=new Date),c.widgets.length&&(b.isApplyingWidgets=!0,c.widgets=h.grep(c.widgets,function(a,b){return h.inArray(a,c.widgets)===b}),h.each(c.widgets||[],function(a,b){(k=g.getWidgetById(b))&&k.id&&(k.priority||(k.priority=10),e[a]=k)}),e.sort(function(a,b){return a.priority<b.priority?-1:a.priority===b.priority?0:1}),h.each(e,function(e,g){if(g){if(a||!c.widgetInit[g.id])c.widgetInit[g.id]=!0,g.hasOwnProperty("options")&&(d=b.config.widgetOptions=h.extend(!0,{},g.options,d)),g.hasOwnProperty("init")&& g.init(b,g,c,d);!a&&g.hasOwnProperty("format")&&g.format(b,c,d,!1)}})),setTimeout(function(){b.isApplyingWidgets=!1},0),c.debug&&(p=c.widgets.length,q("Completed "+(!0===a?"initializing ":"applying ")+p+" widget"+(1!==p?"s":""),l)))};g.refreshWidgets=function(b,a,c){b=h(b)[0];var f,e=b.config,l=e.widgets,q=g.widgets,k=q.length;for(f=0;f<k;f++)q[f]&&q[f].id&&(a||0>h.inArray(q[f].id,l))&&(e.debug&&d('Refeshing widgets: Removing "'+q[f].id+'"'),q[f].hasOwnProperty("remove")&&e.widgetInit[q[f].id]&&(q[f].remove(b, e,e.widgetOptions),e.widgetInit[q[f].id]=!1));!0!==c&&g.applyWidget(b,a)};g.getData=function(b,a,c){var d="";b=h(b);var e,g;if(!b.length)return"";e=h.metadata?b.metadata():!1;g=" "+(b.attr("class")||"");"undefined"!==typeof b.data(c)||"undefined"!==typeof b.data(c.toLowerCase())?d+=b.data(c)||b.data(c.toLowerCase()):e&&"undefined"!==typeof e[c]?d+=e[c]:a&&"undefined"!==typeof a[c]?d+=a[c]:" "!==g&&g.match(" "+c+"-")&&(d=g.match(new RegExp("\\s"+c+"-([\\w-]+)"))[1]||"");return h.trim(d)};g.formatFloat= function(b,a){if("string"!==typeof b||""===b)return b;var c;b=(a&&a.config?!1!==a.config.usNumberFormat:"undefined"!==typeof a?a:1)?b.replace(/,/g,""):b.replace(/[\s|\.]/g,"").replace(/,/g,".");/^\s*\([.\d]+\)/.test(b)&&(b=b.replace(/^\s*\(([.\d]+)\)/,"-$1"));c=parseFloat(b);return isNaN(c)?h.trim(b):c};g.isDigit=function(b){return isNaN(b)?/^[\-+(]?\d+[)]?$/.test(b.toString().replace(/[,.'"\s]/g,"")):!0}}});var r=h.tablesorter;h.fn.extend({tablesorter:r.construct});r.addParser({id:"no-parser",is:function(){return!1}, format:function(){return""},type:"text"});r.addParser({id:"text",is:function(){return!0},format:function(d,q){var p=q.config;d&&(d=h.trim(p.ignoreCase?d.toLocaleLowerCase():d),d=p.sortLocaleCompare?r.replaceAccents(d):d);return d},type:"text"});r.addParser({id:"digit",is:function(d){return r.isDigit(d)},format:function(d,q){var p=r.formatFloat((d||"").replace(/[^\w,. \-()]/g,""),q);return d&&"number"===typeof p?p:d?h.trim(d&&q.config.ignoreCase?d.toLocaleLowerCase():d):d},type:"numeric"});r.addParser({id:"currency", is:function(d){return/^\(?\d+[\u00a3$\u20ac\u00a4\u00a5\u00a2?.]|[\u00a3$\u20ac\u00a4\u00a5\u00a2?.]\d+\)?$/.test((d||"").replace(/[+\-,. ]/g,""))},format:function(d,q){var p=r.formatFloat((d||"").replace(/[^\w,. \-()]/g,""),q);return d&&"number"===typeof p?p:d?h.trim(d&&q.config.ignoreCase?d.toLocaleLowerCase():d):d},type:"numeric"});r.addParser({id:"ipAddress",is:function(d){return/^\d{1,3}[\.]\d{1,3}[\.]\d{1,3}[\.]\d{1,3}$/.test(d)},format:function(d,h){var p,y=d?d.split("."):"",v="",w=y.length; for(p=0;p<w;p++)v+=("00"+y[p]).slice(-3);return d?r.formatFloat(v,h):d},type:"numeric"});r.addParser({id:"url",is:function(d){return/^(https?|ftp|file):\/\//.test(d)},format:function(d){return d?h.trim(d.replace(/(https?|ftp|file):\/\//,"")):d},parsed:!0,type:"text"});r.addParser({id:"isoDate",is:function(d){return/^\d{4}[\/\-]\d{1,2}[\/\-]\d{1,2}/.test(d)},format:function(d,h){return d?r.formatFloat(""!==d?(new Date(d.replace(/-/g,"/"))).getTime()||d:"",h):d},type:"numeric"});r.addParser({id:"percent", is:function(d){return/(\d\s*?%|%\s*?\d)/.test(d)&&15>d.length},format:function(d,h){return d?r.formatFloat(d.replace(/%/g,""),h):d},type:"numeric"});r.addParser({id:"usLongDate",is:function(d){return/^[A-Z]{3,10}\.?\s+\d{1,2},?\s+(\d{4})(\s+\d{1,2}:\d{2}(:\d{2})?(\s+[AP]M)?)?$/i.test(d)||/^\d{1,2}\s+[A-Z]{3,10}\s+\d{4}/i.test(d)},format:function(d,h){return d?r.formatFloat((new Date(d.replace(/(\S)([AP]M)$/i,"$1 $2"))).getTime()||d,h):d},type:"numeric"});r.addParser({id:"shortDate",is:function(d){return/(^\d{1,2}[\/\s]\d{1,2}[\/\s]\d{4})|(^\d{4}[\/\s]\d{1,2}[\/\s]\d{1,2})/.test((d|| "").replace(/\s+/g," ").replace(/[\-.,]/g,"/"))},format:function(d,h,p,y){if(d){p=h.config;var v=p.$headers.filter("[data-column="+y+"]:last");y=v.length&&v[0].dateFormat||r.getData(v,r.getColumnData(h,p.headers,y),"dateFormat")||p.dateFormat;d=d.replace(/\s+/g," ").replace(/[\-.,]/g,"/");"mmddyyyy"===y?d=d.replace(/(\d{1,2})[\/\s](\d{1,2})[\/\s](\d{4})/,"$3/$1/$2"):"ddmmyyyy"===y?d=d.replace(/(\d{1,2})[\/\s](\d{1,2})[\/\s](\d{4})/,"$3/$2/$1"):"yyyymmdd"===y&&(d=d.replace(/(\d{4})[\/\s](\d{1,2})[\/\s](\d{1,2})/, "$1/$2/$3"))}return d?r.formatFloat((new Date(d)).getTime()||d,h):d},type:"numeric"});r.addParser({id:"time",is:function(d){return/^(([0-2]?\d:[0-5]\d)|([0-1]?\d:[0-5]\d\s?([AP]M)))$/i.test(d)},format:function(d,h){return d?r.formatFloat((new Date("2000/01/01 "+d.replace(/(\S)([AP]M)$/i,"$1 $2"))).getTime()||d,h):d},type:"numeric"});r.addParser({id:"metadata",is:function(){return!1},format:function(d,q,p){d=q.config;d=d.parserMetadataName?d.parserMetadataName:"sortValue";return h(p).metadata()[d]}, type:"numeric"});r.addWidget({id:"zebra",priority:90,format:function(d,q,p){var y,v,w,z,D,E=new RegExp(q.cssChildRow,"i"),C=q.$tbodies;q.debug&&(D=new Date);for(d=0;d<C.length;d++)w=0,y=C.eq(d),y=y.children("tr:visible").not(q.selectorRemove),y.each(function(){v=h(this);E.test(this.className)||w++;z=0===w%2;v.removeClass(p.zebra[z?1:0]).addClass(p.zebra[z?0:1])});q.debug&&r.benchmark("Applying Zebra widget",D)},remove:function(d,q,p){var r;q=q.$tbodies;var v=(p.zebra||["even","odd"]).join(" ");for(p= 0;p<q.length;p++)r=h.tablesorter.processTbody(d,q.eq(p),!0),r.children().removeClass(v),h.tablesorter.processTbody(d,r,!1)}})}(jQuery);
/*!
 * ZUI: Standard edition - v1.9.2 - 2020-07-09
 * http://openzui.com
 * GitHub: https://github.com/easysoft/zui.git 
 * Copyright (c) 2020 cnezsoft.com; Licensed MIT
 */
/*! Some code copy from Bootstrap v3.0.0 by @fat and @mdo. (Copyright 2013 Twitter, Inc. Licensed under http://www.apache.org/licenses/)*/
!function(t,e,i){"use strict";if("undefined"==typeof t)throw new Error("ZUI requires jQuery");t.zui||(t.zui=function(e){t.isPlainObject(e)&&t.extend(t.zui,e)});var n={all:-1,left:0,middle:1,right:2},o=0;t.zui({uuid:function(t){var e=1e8*(Date.now()-1580890015292)+1e3*Math.floor(1e5*Math.random())+o++%1e3;return t?e:e.toString(36)},callEvent:function(e,n,o){if(t.isFunction(e)){o!==i&&(e=t.proxy(e,o));var a=e(n);return n&&(n.result=a),!(a!==i&&!a)}return 1},strCode:function(t){var e=0;if(t&&t.length)for(var i=0;i<t.length;++i)e+=i*t.charCodeAt(i);return e},getMouseButtonCode:function(t){return"number"!=typeof t&&(t=n[t]),t!==i&&null!==t||(t=-1),t},defaultLang:"en",clientLang:function(){var i,n=e.config;if("undefined"!=typeof n&&n.clientLang&&(i=n.clientLang),!i){var o=t("html").attr("lang");i=o?o:navigator.userLanguage||navigator.userLanguage||t.zui.defaultLang}return i.replace("-","_").toLowerCase()},langDataMap:{},addLangData:function(e,i,n){var o={};n&&i&&e?(o[i]={},o[i][e]=n):e&&i&&!n?(n=i,t.each(n,function(t){o[t]={},o[t][e]=n[t]})):!e||i||n||t.each(n,function(e){var i=n[e];t.each(i,function(t){o[t]||(o[t]={}),o[t][e]=i[t]})}),t.extend(!0,t.zui.langDataMap,o)},getLangData:function(e,i,n){if(!arguments.length)return t.extend({},t.zui.langDataMap);if(1===arguments.length)return t.extend({},t.zui.langDataMap[e]);if(2===arguments.length){var o=t.zui.langDataMap[e];return o?i?o[i]:o:{}}if(3===arguments.length){i=i||t.zui.clientLang();var o=t.zui.langDataMap[e],a=o?o[i]:{};return t.extend(!0,{},n[i]||n.en||n.zh_cn,a)}return null},lang:function(){return arguments.length&&t.isPlainObject(arguments[arguments.length-1])?t.zui.addLangData.apply(null,arguments):t.zui.getLangData.apply(null,arguments)},_scrollbarWidth:0,checkBodyScrollbar:function(){if(document.body.clientWidth>=e.innerWidth)return 0;if(!t.zui._scrollbarWidth){var i=document.createElement("div");i.className="scrollbar-measure",document.body.appendChild(i),t.zui._scrollbarWidth=i.offsetWidth-i.clientWidth,document.body.removeChild(i)}return t.zui._scrollbarWidth},fixBodyScrollbar:function(){if(t.zui.checkBodyScrollbar()){var e=t("body"),i=parseInt(e.css("padding-right")||0,10);return t.zui._scrollbarWidth&&e.css({paddingRight:i+t.zui._scrollbarWidth,overflowY:"hidden"}),!0}},resetBodyScrollbar:function(){t("body").css({paddingRight:"",overflowY:""})}}),t.fn.callEvent=function(e,n,o){var a=t(this),s=e.indexOf(".zui."),r=s<0?e:e.substring(0,s),l=t.Event(r,n);if(o===i&&s>0&&(o=a.data(e.substring(s+1))),o&&o.options){var d=o.options[r];t.isFunction(d)&&(l.result=t.zui.callEvent(d,l,o))}return a.trigger(l),l},t.fn.callComEvent=function(e,n,o){o===i||t.isArray(o)||(o=[o]);var a,s=this;s.trigger(n,o);var r=e.options[n];return r&&(a=r.apply(e,o)),a}}(jQuery,window,void 0),function(t){"use strict";t.fn.fixOlPd=function(e){return e=e||10,this.each(function(){var i=t(this);i.css("paddingLeft",Math.ceil(Math.log10(i.children().length))*e+10)})},t(function(){t(".ol-pd-fix,.article ol").fixOlPd()})}(jQuery),+function(t){"use strict";var e=function(i,n){this.$element=t(i),this.options=t.extend({},e.DEFAULTS,n),this.isLoading=!1};e.DEFAULTS={loadingText:"loading..."},e.prototype.setState=function(e){var i="disabled",n=this.$element,o=n.is("input")?"val":"html",a=n.data();e+="Text",a.resetText||n.data("resetText",n[o]()),n[o](a[e]||this.options[e]),setTimeout(t.proxy(function(){"loadingText"==e?(this.isLoading=!0,n.addClass(i).attr(i,i)):this.isLoading&&(this.isLoading=!1,n.removeClass(i).removeAttr(i))},this),0)},e.prototype.toggle=function(){var t=!0,e=this.$element.closest('[data-toggle="buttons"]');if(e.length){var i=this.$element.find("input");"radio"==i.prop("type")&&(i.prop("checked")&&this.$element.hasClass("active")?t=!1:e.find(".active").removeClass("active")),t&&i.prop("checked",!this.$element.hasClass("active")).trigger("change")}t&&this.$element.toggleClass("active")};var i=t.fn.button;t.fn.button=function(i){return this.each(function(){var n=t(this),o=n.data("zui.button"),a="object"==typeof i&&i;o||n.data("zui.button",o=new e(this,a)),"toggle"==i?o.toggle():i&&o.setState(i)})},t.fn.button.Constructor=e,t.fn.button.noConflict=function(){return t.fn.button=i,this},t(document).on("click.zui.button.data-api","[data-toggle^=button]",function(e){var i=t(e.target);i.hasClass("btn")||(i=i.closest(".btn")),i.button("toggle"),e.preventDefault()})}(jQuery),+function(t){"use strict";var e='[data-dismiss="alert"]',i="zui.alert",n=function(i){t(i).on("click",e,this.close)};n.prototype.close=function(e){function n(){s.trigger("closed."+i).remove()}var o=t(this),a=o.attr("data-target");a||(a=o.attr("href"),a=a&&a.replace(/.*(?=#[^\s]*$)/,""));var s=t(a);e&&e.preventDefault(),s.length||(s=o.hasClass("alert")?o:o.parent()),s.trigger(e=t.Event("close."+i)),e.isDefaultPrevented()||(s.removeClass("in"),t.support.transition&&s.hasClass("fade")?s.one(t.support.transition.end,n).emulateTransitionEnd(150):n())};var o=t.fn.alert;t.fn.alert=function(e){return this.each(function(){var o=t(this),a=o.data(i);a||o.data(i,a=new n(this)),"string"==typeof e&&a[e].call(o)})},t.fn.alert.Constructor=n,t.fn.alert.noConflict=function(){return t.fn.alert=o,this},t(document).on("click."+i+".data-api",e,n.prototype.close)}(window.jQuery),function(t,e){"use strict";var i="zui.pager",n={page:1,recTotal:0,recPerPage:10},o={zh_cn:{pageOfText:"第 {0} 页",prev:"上一页",next:"下一页",first:"第一页",last:"最后一页","goto":"跳转",pageOf:"第 <strong>{page}</strong> 页",totalPage:"共 <strong>{totalPage}</strong> 页",totalCount:"共 <strong>{recTotal}</strong> 项",pageSize:"每页 <strong>{recPerPage}</strong> 项",itemsRange:"第 <strong>{start}</strong> ~ <strong>{end}</strong> 项",pageOfTotal:"第 <strong>{page}</strong>/<strong>{totalPage}</strong> 页"},zh_tw:{pageOfText:"第 {0} 頁",prev:"上一頁",next:"下一頁",first:"第一頁",last:"最後一頁","goto":"跳轉",pageOf:"第 <strong>{page}</strong> 頁",totalPage:"共 <strong>{totalPage}</strong> 頁",totalCount:"共 <strong>{recTotal}</strong> 項",pageSize:"每頁 <strong>{recPerPage}</strong> 項",itemsRange:"第 <strong>{start}</strong> ~ <strong>{end}</strong> 項",pageOfTotal:"第 <strong>{page}</strong>/<strong>{totalPage}</strong> 頁"},en:{pageOfText:"Page {0}",prev:"Prev",next:"Next",first:"First",last:"Last","goto":"Goto",pageOf:"Page <strong>{page}</strong>",totalPage:"<strong>{totalPage}</strong> pages",totalCount:"Total: <strong>{recTotal}</strong> items",pageSize:"<strong>{recPerPage}</strong> per page",itemsRange:"From <strong>{start}</strong> to <strong>{end}</strong>",pageOfTotal:"Page <strong>{page}</strong> of <strong>{totalPage}</strong>"}},a=function(e,n){var s=this;s.name=i,s.$=t(e),n=s.options=t.extend({},a.DEFAULTS,this.$.data(),n),s.langName=n.lang||t.zui.clientLang(),s.lang=t.zui.getLangData(i,s.langName,o),s.state={},s.set(n.page,n.recTotal,n.recPerPage,!0),s.$.on("click",".pager-goto-btn",function(){var e=t(this).closest(".pager-goto"),i=parseInt(e.find(".pager-goto-input").val());NaN!==i&&s.set(i)}).on("click",".pager-item",function(){var e=t(this).data("page");"number"==typeof e&&e>0&&s.set(e)}).on("click",".pager-size-menu [data-size]",function(){var e=t(this).data("size");"number"==typeof e&&e>0&&s.set(-1,-1,e)})};a.prototype.set=function(e,i,o,a){var s=this;"object"==typeof e&&null!==e&&(o=e.recPerPage,i=e.recTotal,e=e.page);var r=s.state;r||(r=t.extend({},n));var l=t.extend({},r);return"number"==typeof o&&o>0&&(r.recPerPage=o),"number"==typeof i&&i>=0&&(r.recTotal=i),"number"==typeof e&&e>=0&&(r.page=e),r.totalPage=r.recTotal&&r.recPerPage?Math.ceil(r.recTotal/r.recPerPage):1,r.page=Math.max(0,Math.min(r.page,r.totalPage)),r.pageRecCount=r.recTotal,r.page&&r.recTotal&&(r.page<r.totalPage?r.pageRecCount=r.recPerPage:r.page>1&&(r.pageRecCount=r.recTotal-r.recPerPage*(r.page-1))),r.skip=r.page>1?(r.page-1)*r.recPerPage:0,r.start=r.skip+1,r.end=r.skip+r.pageRecCount,r.prev=r.page>1?r.page-1:0,r.next=r.page<r.totalPage?r.page+1:0,s.state=r,a||l.page===r.page&&l.recTotal===r.recTotal&&l.recPerPage===r.recPerPage||s.$.callComEvent(s,"onPageChange",[r,l]),s.render()},a.prototype.createLinkItem=function(i,n,o){var a=this;n===e&&(n=i);var s=t('<a title="'+a.lang.pageOfText.format(i)+'" class="pager-item" data-page="'+i+'"/>').attr("href",i?a.createLink(i,a.state):"###").html(n);return o||(s=t("<li />").append(s).toggleClass("active",i===a.state.page).toggleClass("disabled",!i||i===a.state.page)),s},a.prototype.createNavItems=function(t){var i=this,n=i.$,o=i.state,a=o.totalPage,s=o.page,r=function(t,o){if(t===!1)return void n.append(i.createLinkItem(0,o||i.options.navEllipsisItem));o===e&&(o=t);for(var a=t;a<=o;++a)n.append(i.createLinkItem(a))};t===e&&(t=i.options.maxNavCount||10),r(1),a>1&&(a<=t?r(2,a):s<t-2?(r(2,t-2),r(!1),r(a)):s>a-t+2?(r(!1),r(a-t+2,a)):(r(!1),r(s-Math.ceil((t-4)/2),s+Math.floor((t-4)/2)),r(!1),r(a)))},a.prototype.createGoto=function(){var e=this,i=this.state,n=t('<div class="input-group pager-goto" style="width: '+(35+9*(i.page+"").length+25+12*e.lang["goto"].length)+'px"><input value="'+i.page+'" type="number" min="1" max="'+i.totalPage+'" placeholder="'+i.page+'" class="form-control pager-goto-input"><span class="input-group-btn"><button class="btn pager-goto-btn" type="button">'+e.lang["goto"]+"</button></span></div>");return n},a.prototype.createSizeMenu=function(){var e=this,i=this.state,n=t('<ul class="dropdown-menu"></ul>'),o=e.options.pageSizeOptions;"string"==typeof o&&(o=o.split(","));for(var a=0;a<o.length;++a){var s=o[a];"string"==typeof s&&(s=parseInt(s));var r=t('<li><a href="###" data-size="'+s+'">'+s+"</a></li>").toggleClass("active",s===i.recPerPage);n.append(r)}return t('<div class="btn-group pager-size-menu"><button type="button" class="btn dropdown-toggle" data-toggle="dropdown">'+e.lang.pageSize.format(i)+' <span class="caret"></span></button></div>').addClass(e.options.menuDirection).append(n)},a.prototype.createElement=function(e,i,n){var o=this,a=t.proxy(o.createLinkItem,o),s=o.lang;switch(e){case"prev":return a(n.prev,s.prev);case"prev_icon":return a(n.prev,'<i class="icon '+o.options.prevIcon+'"></i>');case"next":return a(n.next,s.next);case"next_icon":return a(n.next,'<i class="icon '+o.options.nextIcon+'"></i>');case"first":return a(1,s.first);case"first_icon":return a(1,'<i class="icon '+o.options.firstIcon+'"></i>');case"last":return a(n.totalPage,s.last);case"last_icon":return a(n.totalPage,'<i class="icon '+o.options.lastIcon+'"></i>');case"space":case"|":return t('<li class="space" />');case"nav":case"pages":return void o.createNavItems();case"total_text":return t(('<div class="pager-label">'+s.totalCount+"</div>").format(n));case"page_text":return t(('<div class="pager-label">'+s.pageOf+"</div>").format(n));case"total_page_text":return t(('<div class="pager-label">'+s.totalPage+"</div>").format(n));case"page_of_total_text":return t(('<div class="pager-label">'+s.pageOfTotal+"</div>").format(n));case"page_size_text":return t(('<div class="pager-label">'+s.pageSize+"</div>").format(n));case"items_range_text":return t(('<div class="pager-label">'+s.itemsRange+"</div>").format(n));case"goto":return o.createGoto();case"size_menu":return o.createSizeMenu();default:return t("<li/>").html(e.format(n))}},a.prototype.createLink=function(i,n){i===e&&(i=this.state.page),n===e&&(n=this.state);var o=this.options.linkCreator;return"string"==typeof o?o.format(t.extend({},n,{page:i})):t.isFunction(o)?o(i,n):"#page="+i},a.prototype.render=function(e){var i=this,n=i.state,o=i.options.elementCreator||i.createElement,a=t.isPlainObject(o);e=e||i.elements||i.options.elements,"string"==typeof e&&(e=e.split(",")),i.elements=e,i.$.empty();for(var s=0;s<e.length;++s){var r=t.trim(e[s]),l=a?o[r]||o:o,d=l.call(i,r,i.$,n);d===!1&&(d=i.createElement(r,i.$,n)),d instanceof t&&("LI"!==d[0].tagName&&(d=t("<li/>").append(d)),i.$.append(d))}var c=null;return i.$.children("li").each(function(){var e=t(this),i=!!e.children(".pager-item").length;c?c.toggleClass("pager-item-right",!i):i&&e.addClass("pager-item-left"),c=i?e:null}),c&&c.addClass("pager-item-right"),i.$.callComEvent(i,"onRender",[n]),i},a.DEFAULTS=t.extend({elements:["first_icon","prev_icon","pages","next_icon","last_icon","page_of_total_text","items_range_text","total_text"],prevIcon:"icon-double-angle-left",nextIcon:"icon-double-angle-right",firstIcon:"icon-step-backward",lastIcon:"icon-step-forward",navEllipsisItem:'<i class="icon icon-ellipsis-h"></i>',maxNavCount:10,menuDirection:"dropdown",pageSizeOptions:[10,20,30,50,100]},n),t.fn.pager=function(e){return this.each(function(){var n=t(this),o=n.data(i),s="object"==typeof e&&e;o||n.data(i,o=new a(this,s)),"string"==typeof e&&o[e]()})},a.NAME=i,a.LANG=o,t.fn.pager.Constructor=a,t(function(){t('[data-ride="pager"]').pager()})}(jQuery,void 0),+function(t){"use strict";var e="zui.tab",i=function(e){this.element=t(e)};i.prototype.show=function(){var i=this.element,n=i.closest("ul:not(.dropdown-menu)"),o=i.attr("data-target")||i.attr("data-tab");if(o||(o=i.attr("href"),o=o&&o.replace(/.*(?=#[^\s]*$)/,"")),!i.parent("li").hasClass("active")){var a=n.find(".active:last a")[0],s=t.Event("show."+e,{relatedTarget:a});if(i.trigger(s),!s.isDefaultPrevented()){var r=t(o);this.activate(i.parent("li"),n),this.activate(r,r.parent(),function(){i.trigger({type:"shown."+e,relatedTarget:a})})}}},i.prototype.activate=function(e,i,n){function o(){a.removeClass("active").find("> .dropdown-menu > .active").removeClass("active"),e.addClass("active"),s?(e[0].offsetWidth,e.addClass("in")):e.removeClass("fade"),e.parent(".dropdown-menu")&&e.closest("li.dropdown").addClass("active"),n&&n()}var a=i.find("> .active"),s=n&&t.support.transition&&a.hasClass("fade");s?a.one(t.support.transition.end,o).emulateTransitionEnd(150):o(),a.removeClass("in")};var n=t.fn.tab;t.fn.tab=function(n){return this.each(function(){var o=t(this),a=o.data(e);a||o.data(e,a=new i(this)),"string"==typeof n&&a[n]()})},t.fn.tab.Constructor=i,t.fn.tab.noConflict=function(){return t.fn.tab=n,this},t(document).on("click.zui.tab.data-api",'[data-toggle="tab"], [data-tab]',function(e){e.preventDefault(),t(this).tab("show")})}(window.jQuery),+function(t){"use strict";function e(){var t=document.createElement("bootstrap"),e={WebkitTransition:"webkitTransitionEnd",MozTransition:"transitionend",OTransition:"oTransitionEnd otransitionend",transition:"transitionend"};for(var i in e)if(void 0!==t.style[i])return{end:e[i]};return!1}t.fn.emulateTransitionEnd=function(e){var i=!1,n=this;t(this).one("bsTransitionEnd",function(){i=!0});var o=function(){i||t(n).trigger(t.support.transition.end)};return setTimeout(o,e),this},t(function(){t.support.transition=e(),t.support.transition&&(t.event.special.bsTransitionEnd={bindType:t.support.transition.end,delegateType:t.support.transition.end,handle:function(e){if(t(e.target).is(this))return e.handleObj.handler.apply(this,arguments)}})})}(jQuery),+function(t){"use strict";var e="zui.collapse",i=function(e,n){this.$element=t(e),this.options=t.extend({},i.DEFAULTS,n),this.transitioning=null,this.options.parent&&(this.$parent=t(this.options.parent)),this.options.toggle&&this.toggle()};i.DEFAULTS={toggle:!0},i.prototype.dimension=function(){var t=this.$element.hasClass("width");return t?"width":"height"},i.prototype.show=function(){if(!this.transitioning&&!this.$element.hasClass("in")){var i=t.Event("show."+e);if(this.$element.trigger(i),!i.isDefaultPrevented()){var n=this.$parent&&this.$parent.find(".in");if(n&&n.length){var o=n.data(e);if(o&&o.transitioning)return;n.collapse("hide"),o||n.data(e,null)}var a=this.dimension();this.$element.removeClass("collapse").addClass("collapsing")[a](0),this.transitioning=1;var s=function(){this.$element.removeClass("collapsing").addClass("in")[a]("auto"),this.transitioning=0,this.$element.trigger("shown."+e)};if(!t.support.transition)return s.call(this);var r=t.camelCase(["scroll",a].join("-"));this.$element.one(t.support.transition.end,t.proxy(s,this)).emulateTransitionEnd(350)[a](this.$element[0][r])}}},i.prototype.hide=function(){if(!this.transitioning&&this.$element.hasClass("in")){var i=t.Event("hide."+e);if(this.$element.trigger(i),!i.isDefaultPrevented()){var n=this.dimension();this.$element[n](this.$element[n]())[0].offsetHeight,this.$element.addClass("collapsing").removeClass("collapse").removeClass("in"),this.transitioning=1;var o=function(){this.transitioning=0,this.$element.trigger("hidden."+e).removeClass("collapsing").addClass("collapse")};return t.support.transition?void this.$element[n](0).one(t.support.transition.end,t.proxy(o,this)).emulateTransitionEnd(350):o.call(this)}}},i.prototype.toggle=function(){this[this.$element.hasClass("in")?"hide":"show"]()};var n=t.fn.collapse;t.fn.collapse=function(n){return this.each(function(){var o=t(this),a=o.data(e),s=t.extend({},i.DEFAULTS,o.data(),"object"==typeof n&&n);a||o.data(e,a=new i(this,s)),"string"==typeof n&&a[n]()})},t.fn.collapse.Constructor=i,t.fn.collapse.noConflict=function(){return t.fn.collapse=n,this},t(document).on("click."+e+".data-api","[data-toggle=collapse]",function(i){var n,o=t(this),a=o.attr("data-target")||i.preventDefault()||(n=o.attr("href"))&&n.replace(/.*(?=#[^\s]+$)/,""),s=t(a),r=s.data(e),l=r?"toggle":o.data(),d=o.attr("data-parent"),c=d&&t(d);r&&r.transitioning||(c&&c.find('[data-toggle=collapse][data-parent="'+d+'"]').not(o).addClass("collapsed"),o[s.hasClass("in")?"addClass":"removeClass"]("collapsed")),s.collapse(l)})}(window.jQuery),function(t,e){"use strict";var i=1200,n=992,o=768,a=e(t),s=function(){var t=a.width();e("html").toggleClass("screen-desktop",t>=n&&t<i).toggleClass("screen-desktop-wide",t>=i).toggleClass("screen-tablet",t>=o&&t<n).toggleClass("screen-phone",t<o).toggleClass("device-mobile",t<n).toggleClass("device-desktop",t>=n)},r="",l=navigator.userAgent;l.match(/(iPad|iPhone|iPod)/i)?r+=" os-ios":l.match(/android/i)?r+=" os-android":l.match(/Win/i)?r+=" os-windows":l.match(/Mac/i)?r+=" os-mac":l.match(/Linux/i)?r+=" os-linux":l.match(/X11/i)&&(r+=" os-unix"),"ontouchstart"in document.documentElement&&(r+=" is-touchable"),e("html").addClass(r),a.resize(s),s()}(window,jQuery),function(t){"use strict";var e={zh_cn:'您的浏览器版本过低，无法体验所有功能，建议升级或者更换浏览器。 <a href="https://browsehappy.com/" target="_blank" class="alert-link">了解更多...</a>',zh_tw:'您的瀏覽器版本過低，無法體驗所有功能，建議升級或者更换瀏覽器。<a href="https://browsehappy.com/" target="_blank" class="alert-link">了解更多...</a>',en:'Your browser is too old, it has been unable to experience the colorful internet. We strongly recommend that you upgrade a better one. <a href="https://browsehappy.com/" target="_blank" class="alert-link">Learn more...</a>'},i=function(){for(var t=!1,e=11;e>5;e--)if(this.isIE(e)){t=e;break}this.ie=t,this.cssHelper()};i.prototype.cssHelper=function(){var e=this.ie,i=t("html");i.toggleClass("ie",e).removeClass("ie-6 ie-7 ie-8 ie-9 ie-10"),e&&i.addClass("ie-"+e).toggleClass("gt-ie-7 gte-ie-8 support-ie",e>=8).toggleClass("lte-ie-7 lt-ie-8 outdated-ie",e<8).toggleClass("gt-ie-8 gte-ie-9",e>=9).toggleClass("lte-ie-8 lt-ie-9",e<9).toggleClass("gt-ie-9 gte-ie-10",e>=10).toggleClass("lte-ie-9 lt-ie-10",e<10).toggleClass("gt-ie-10 gte-ie-11",e>=11).toggleClass("lte-ie-10 lt-ie-11",e<11)},i.prototype.tip=function(i){var n=t("#browseHappyTip");n.length||(n=t('<div id="browseHappyTip" class="alert alert-dismissable alert-danger-inverse alert-block" style="position: relative; z-index: 99999"><button type="button" class="close" data-dismiss="alert" aria-hidden="true">×</button><div class="container"><div class="content text-center"></div></div></div>'),n.prependTo("body")),i||(i=t.zui.getLangData("zui.browser",t.zui.clientLang(),e),"object"==typeof i&&(i=i.tip)),n.find(".content").html(i)},i.prototype.isIE=function(t){if(11===t)return this.isIE11();if(10===t)return this.isIE10();if(!t&&(this.isIE11()||this.isIE10()))return!0;var e=document.createElement("b");return e.innerHTML="<!--[if IE "+(t||"")+"]><i></i><![endif]-->",1===e.getElementsByTagName("i").length},i.prototype.isIE10=function(){return navigator.appVersion.indexOf("MSIE 10")!==-1},i.prototype.isIE11=function(){var t=navigator.userAgent;return t.indexOf("Trident")!==-1&&t.indexOf("rv:11")!==-1},t.zui({browser:new i}),t(function(){t("body").hasClass("disabled-browser-tip")||t.zui.browser.ie&&t.zui.browser.ie<8&&t.zui.browser.tip()})}(jQuery),function(t){"use strict";const e=864e5,i=function(t){return t instanceof Date||("number"==typeof t&&t<1e10&&(t*=1e3),t=new Date(t)),t},n=function(t){return i(t).getTime()},o=function(t,e){t=i(t),void 0===e&&(e="yyyy-MM-dd hh:mm:ss");var n={"M+":t.getMonth()+1,"d+":t.getDate(),"h+":t.getHours(),"m+":t.getMinutes(),"s+":t.getSeconds(),"q+":Math.floor((t.getMonth()+3)/3),"S+":t.getMilliseconds()};/(y+)/i.test(e)&&(e=e.replace(RegExp.$1,(t.getFullYear()+"").substr(4-RegExp.$1.length)));for(var o in n)new RegExp("("+o+")").test(e)&&(e=e.replace(RegExp.$1,1==RegExp.$1.length?n[o]:("00"+n[o]).substr((""+n[o]).length)));return e},a=function(t,e){return t.setTime(t.getTime()+e),t},s=function(t,i){return a(t,i*e)},r=function(t){return new Date(i(t).getTime())},l=function(t){return t%4===0&&t%100!==0||t%400===0},d=function(t,e){return[31,l(t)?29:28,31,30,31,30,31,31,30,31,30,31][e]},c=function(t){return d(t.getFullYear(),t.getMonth())},p=function(t){return t.setHours(0),t.setMinutes(0),t.setSeconds(0),t.setMilliseconds(0),t},u=function(t,e){var i=t.getDate();return t.setDate(1),t.setMonth(t.getMonth()+e),t.setDate(Math.min(i,c(t))),t},h=function(t,e){e=e||1;for(var i=new Date(t.getTime());i.getDay()!=e;)i=s(i,-1);return p(i)},f=function(t,e){return t.toDateString()===e.toDateString()},g=function(t,e){var i=h(t),n=s(r(i),7);return e>=i&&e<n},m=function(t,e){return t.getFullYear()===e.getFullYear()},v={formatDate:o,createDate:i,date:{ONEDAY_TICKS:e,create:i,getTimestamp:n,format:o,addMilliseconds:a,addDays:s,cloneDate:r,isLeapYear:l,getDaysInMonth:d,getDaysOfThisMonth:c,clearTime:p,addMonths:u,getLastWeekday:h,isSameDay:f,isSameWeek:g,isSameYear:m}};t.$&&t.$.zui?$.zui(v):t.dateHelper=v.date,t.noDatePrototypeHelper||(Date.ONEDAY_TICKS=e,Date.prototype.format||(Date.prototype.format=function(t){return o(this,t)}),Date.prototype.addMilliseconds||(Date.prototype.addMilliseconds=function(t){return a(this,t)}),Date.prototype.addDays||(Date.prototype.addDays=function(t){return s(this,t)}),Date.prototype.clone||(Date.prototype.clone=function(){return r(this)}),Date.isLeapYear||(Date.isLeapYear=function(t){return l(t)}),Date.getDaysInMonth||(Date.getDaysInMonth=function(t,e){return d(t,e)}),Date.prototype.isLeapYear||(Date.prototype.isLeapYear=function(){return l(this.getFullYear())}),Date.prototype.clearTime||(Date.prototype.clearTime=function(){return p(this)}),Date.prototype.getDaysInMonth||(Date.prototype.getDaysInMonth=function(){return c(this)}),Date.prototype.addMonths||(Date.prototype.addMonths=function(t){return u(this,t)}),Date.prototype.getLastWeekday||(Date.prototype.getLastWeekday=function(t){return h(this,t)}),Date.prototype.isSameDay||(Date.prototype.isSameDay=function(t){return f(t,this)}),Date.prototype.isSameWeek||(Date.prototype.isSameWeek=function(t){return g(t,this)}),Date.prototype.isSameYear||(Date.prototype.isSameYear=function(t){return m(this,t)}),Date.create||(Date.create=function(t){return i(t)}),Date.timestamp||(Date.timestamp=function(t){return n(t)}))}(window),function(){"use strict";const t=function(t,e){if(arguments.length>1){var i;if(2==arguments.length&&"object"==typeof e)for(var n in e)void 0!==e[n]&&(i=new RegExp("({"+n+"})","g"),t=t.replace(i,e[n]));else for(var o=1;o<arguments.length;o++)void 0!==arguments[o]&&(i=new RegExp("({["+(o-1)+"]})","g"),t=t.replace(i,arguments[o]))}return t},e=function(t){if(null!==t){var e,i;return i=/\d*/i,e=t.match(i),e==t}return!1},i={formatString:t,string:{format:t,isNum:e}};window.$&&window.$.zui?$.zui(i):window.stringHelper=i.string,window.noStringPrototypeHelper||(String.prototype.format||(String.prototype.format=function(){var e=[].slice.call(arguments);return e.unshift(this),t.apply(this,e)}),String.prototype.isNum||(String.prototype.isNum=function(){return e(this)}),String.prototype.endsWith||(String.prototype.endsWith=function(t,e){return(void 0===e||e>this.length)&&(e=this.length),this.substring(e-t.length,e)===t}),String.prototype.startsWith||Object.defineProperty(String.prototype,"startsWith",{value:function(t,e){return e=!e||e<0?0:+e,this.substring(e,e+t.length)===t}}),String.prototype.includes||(String.prototype.includes=function(){return String.prototype.indexOf.apply(this,arguments)!==-1}))}(),/*!
 * jQuery resize event - v1.1
 * http://benalman.com/projects/jquery-resize-plugin/
 * Copyright (c) 2010 "Cowboy" Ben Alman
 * MIT & GPL http://benalman.com/about/license/
 */
function(t,e,i){"$:nomunge";function n(){o=e[r](function(){a.each(function(){var e=t(this),i=e.width(),n=e.height(),o=t.data(this,d);i===o.w&&n===o.h||e.trigger(l,[o.w=i,o.h=n])}),n()},s[c])}var o,a=t([]),s=t.resize=t.extend(t.resize,{}),r="setTimeout",l="resize",d=l+"-special-event",c="delay",p="throttleWindow";s[c]=250,s[p]=!0,t.event.special[l]={setup:function(){if(!s[p]&&this[r])return!1;var e=t(this);a=a.add(e),t.data(this,d,{w:e.width(),h:e.height()}),1===a.length&&n()},teardown:function(){if(!s[p]&&this[r])return!1;var e=t(this);a=a.not(e),e.removeData(d),a.length||clearTimeout(o)},add:function(e){function n(e,n,a){var s=t(this),r=t.data(this,d)||{};r.w=n!==i?n:s.width(),r.h=a!==i?a:s.height(),o.apply(this,arguments)}if(!s[p]&&this[r])return!1;var o;return t.isFunction(e)?(o=e,n):(o=e.handler,void(e.handler=n))}}}(jQuery,this),+function(t){"use strict";function e(n,o){var a,s=t.proxy(this.process,this);this.$element=t(t(n).is("body")?window:n),this.$body=t("body"),this.$scrollElement=this.$element.on("scroll."+i+".data-api",s),this.options=t.extend({},e.DEFAULTS,o),this.selector||(this.selector=(this.options.target||(a=t(n).attr("href"))&&a.replace(/.*(?=#[^\s]+$)/,"")||"")+" .nav li > a"),this.offsets=t([]),this.targets=t([]),this.activeTarget=null,this.refresh(),this.process()}var i="zui.scrollspy";e.DEFAULTS={offset:10},e.prototype.refresh=function(){var e=this.$element[0]==window?"offset":"position";this.offsets=t([]),this.targets=t([]);var i=this;this.$body.find(this.selector).map(function(){var n=t(this),o=n.data("target")||n.attr("href"),a=/^#./.test(o)&&t(o);return a&&a.length&&a.is(":visible")&&[[a[e]().top+(!t.isWindow(i.$scrollElement.get(0))&&i.$scrollElement.scrollTop()),o]]||null}).sort(function(t,e){return t[0]-e[0]}).each(function(){i.offsets.push(this[0]),i.targets.push(this[1])})},e.prototype.process=function(){var t,e=this.$scrollElement.scrollTop()+this.options.offset,i=this.$scrollElement[0].scrollHeight||this.$body[0].scrollHeight,n=i-this.$scrollElement.height(),o=this.offsets,a=this.targets,s=this.activeTarget;if(e>=n)return s!=(t=a.last()[0])&&this.activate(t);if(s&&e<=o[0])return s!=(t=a[0])&&this.activate(t);for(t=o.length;t--;)s!=a[t]&&e>=o[t]&&(!o[t+1]||e<=o[t+1])&&this.activate(a[t])},e.prototype.activate=function(e){this.activeTarget=e,t(this.selector).parentsUntil(this.options.target,".active").removeClass("active");var n=this.selector+'[data-target="'+e+'"],'+this.selector+'[href="'+e+'"]',o=t(n).parents("li").addClass("active");o.parent(".dropdown-menu").length&&(o=o.closest("li.dropdown").addClass("active")),o.trigger("activate."+i)};var n=t.fn.scrollspy;t.fn.scrollspy=function(n){return this.each(function(){var o=t(this),a=o.data(i),s="object"==typeof n&&n;a||o.data(i,a=new e(this,s)),"string"==typeof n&&a[n]()})},t.fn.scrollspy.Constructor=e,t.fn.scrollspy.noConflict=function(){return t.fn.scrollspy=n,this},t(window).on("load",function(){t('[data-spy="scroll"]').each(function(){var e=t(this);e.scrollspy(e.data())})})}(jQuery),function(t,e){"use strict";var i,n,o="localStorage",a="page_"+t.location.pathname+t.location.search,s=function(){this.silence=!0;try{o in t&&t[o]&&t[o].setItem&&(this.enable=!0,i=t[o])}catch(s){}this.enable||(n={},i={getLength:function(){var t=0;return e.each(n,function(){t++}),t},key:function(t){var i,o=0;return e.each(n,function(e){return o===t?(i=e,!1):void o++}),i},removeItem:function(t){delete n[t]},getItem:function(t){return n[t]},setItem:function(t,e){n[t]=e},clear:function(){n={}}}),this.storage=i,this.page=this.get(a,{})};s.prototype.pageSave=function(){if(e.isEmptyObject(this.page))this.remove(a);else{var t,i=[];for(t in this.page){var n=this.page[t];null===n&&i.push(t)}for(t=i.length-1;t>=0;t--)delete this.page[i[t]];this.set(a,this.page)}},s.prototype.pageRemove=function(t){"undefined"!=typeof this.page[t]&&(this.page[t]=null,this.pageSave())},s.prototype.pageClear=function(){this.page={},this.pageSave()},s.prototype.pageGet=function(t,e){var i=this.page[t];return void 0===e||null!==i&&void 0!==i?i:e},s.prototype.pageSet=function(t,i){e.isPlainObject(t)?e.extend(!0,this.page,t):this.page[this.serialize(t)]=i,this.pageSave()},s.prototype.check=function(){if(!this.enable&&!this.silence)throw new Error("Browser not support localStorage or enable status been set true.");return this.enable},s.prototype.length=function(){return this.check()?i.getLength?i.getLength():i.length:0},s.prototype.removeItem=function(t){return i.removeItem(t),this},s.prototype.remove=function(t){return this.removeItem(t)},s.prototype.getItem=function(t){return i.getItem(t)},s.prototype.get=function(t,e){var i=this.deserialize(this.getItem(t));return"undefined"!=typeof i&&null!==i||"undefined"==typeof e?i:e},s.prototype.key=function(t){return i.key(t)},s.prototype.setItem=function(t,e){return i.setItem(t,e),this},s.prototype.set=function(t,e){return void 0===e?this.remove(t):(this.setItem(t,this.serialize(e)),this)},s.prototype.clear=function(){return i.clear(),this},s.prototype.forEach=function(t){for(var e=this.length(),n=e-1;n>=0;n--){var o=i.key(n);t(o,this.get(o))}return this},s.prototype.getAll=function(){var t={};return this.forEach(function(e,i){t[e]=i}),t},s.prototype.serialize=function(t){return"string"==typeof t?t:JSON.stringify(t)},s.prototype.deserialize=function(t){if("string"==typeof t)try{return JSON.parse(t)}catch(e){return t||void 0}},e.zui({store:new s})}(window,jQuery),function(t){"use strict";var e="zui.searchBox",i=function(e,n){var o=this;o.name=name,o.$=t(e),o.options=n=t.extend({},i.DEFAULTS,o.$.data(),n);var a=o.$.is(n.inputSelector)?o.$:o.$.find(n.inputSelector);if(a.length){var s=function(){o.changeTimer&&(clearTimeout(o.changeTimer),o.changeTimer=null)},r=function(){s();var t=o.getSearch();if(t!==o.lastValue){var e=""===t;a.toggleClass("empty",e),o.$.callComEvent(o,"onSearchChange",[t,e]),o.lastValue=t}};o.$input=a=a.first(),a.on(n.listenEvent,function(t){o.changeTimer=setTimeout(function(){r()},n.changeDelay)}).on("focus",function(t){a.addClass("focus"),o.$.callComEvent(o,"onFocus",[t])}).on("blur",function(t){a.removeClass("focus"),o.$.callComEvent(o,"onBlur",[t])}).on("keydown",function(t){var e=0,i=t.which;27===i&&n.escToClear?(this.setSearch("",!0),r(),e=1):13===i&&n.onPressEnter&&(r(),o.$.callComEvent(o,"onPressEnter",[t]));var a=o.$.callComEvent(o,"onKeyDown",[t]);a===!1&&(e=1),e&&t.preventDefault()}),o.$.on("click",".search-clear-btn",function(t){o.setSearch("",!0),r(),o.focus(),t.preventDefault()}),r()}else console.error("ZUI: search box init error, cannot find search box input element.")};i.DEFAULTS={inputSelector:'input[type="search"],input[type="text"]',listenEvent:"change input paste",changeDelay:500},i.prototype.getSearch=function(){return this.$input&&t.trim(this.$input.val())},i.prototype.setSearch=function(t,e){var i=this.$input;i&&(i.val(t),e||i.trigger("change"))},i.prototype.focus=function(){this.$input&&this.$input.focus()},t.fn.searchBox=function(n){return this.each(function(){var o=t(this),a=o.data(e),s="object"==typeof n&&n;a||o.data(e,a=new i(this,s)),"string"==typeof n&&a[n]()})},i.NAME=e,t.fn.searchBox.Constructor=i}(jQuery),function(t,e){"use strict";var i="zui.draggable",n={container:"body",move:!0},o=0,a=function(e,i){var a=this;a.$=t(e),a.id=o++,a.options=t.extend({},n,a.$.data(),i),a.init()};a.DEFAULTS=n,a.NAME=i,a.prototype.init=function(){var n,o,a,s,r,l=this,d=l.$,c="before",p="drag",u="finish",h="."+i+"."+l.id,f="mousedown"+h,g="mouseup"+h,m="mousemove"+h,v=l.options,y=v.selector,b=v.handle,w=d,C=t.isFunction(v.move),x=function(t){var e=t.pageX,i=t.pageY;r=!0;var o={left:e-a.x,top:i-a.y};w.removeClass("drag-ready").addClass("dragging"),v.move&&(C?v.move(o,w):w.css(o)),v[p]&&v[p]({event:t,element:w,startOffset:a,pos:o,offset:{x:e-n.x,y:i-n.y},smallOffset:{x:e-s.x,y:i-s.y}}),s.x=e,s.y=i,v.stopPropagation&&t.stopPropagation()},$=function(i){if(t(e).off(h),!r)return void w.removeClass("drag-ready");var o={left:i.pageX-a.x,top:i.pageY-a.y};w.removeClass("drag-ready dragging"),v.move&&(C?v.move(o,w):w.css(o)),v[u]&&v[u]({event:i,element:w,startOffset:a,pos:o,offset:{x:i.pageX-n.x,y:i.pageY-n.y},smallOffset:{x:i.pageX-s.x,y:i.pageY-s.y}}),i.preventDefault(),v.stopPropagation&&i.stopPropagation()},T=function(i){var l=t.zui.getMouseButtonCode(v.mouseButton);if(!(l>-1&&i.button!==l)){var d=t(this);if(y&&(w=b?d.closest(y):d),v[c]){var p=v[c]({event:i,element:w});if(p===!1)return}var u=t(v.container),h=w.offset();o=u.offset(),n={x:i.pageX,y:i.pageY},a={x:i.pageX-h.left+o.left,y:i.pageY-h.top+o.top},s=t.extend({},n),r=!1,w.addClass("drag-ready"),i.preventDefault(),v.stopPropagation&&i.stopPropagation(),t(e).on(m,x).on(g,$)}};b?d.on(f,b,T):y?d.on(f,y,T):d.on(f,T)},a.prototype.destroy=function(){var n="."+i+"."+this.id;this.$.off(n),t(e).off(n),this.$.data(i,null)},t.fn.draggable=function(e){return this.each(function(){var n=t(this),o=n.data(i),s="object"==typeof e&&e;o||n.data(i,o=new a(this,s)),"string"==typeof e&&o[e]()})},t.fn.draggable.Constructor=a}(jQuery,document),function(t,e,i){"use strict";var n="zui.droppable",o={target:".droppable-target",deviation:5,sensorOffsetX:0,sensorOffsetY:0,dropToClass:"drop-to"},a=0,s=function(e,i){var n=this;n.id=a++,n.$=t(e),n.options=t.extend({},o,n.$.data(),i),n.init()};s.DEFAULTS=o,s.NAME=n,s.prototype.trigger=function(e,i){return t.zui.callEvent(this.options[e],i,this)},s.prototype.init=function(){var o,a,s,r,l,d,c,p,u,h,f,g,m,v=this,y=v.$,b=v.options,w=b.deviation,C="."+n+"."+v.id,x="mousedown"+C,$="mouseup"+C,T="mousemove"+C,D=b.selector,S=b.handle,k=b.flex,z=b.container,E=b.canMoveHere,P=b.dropToClass,I=y,M=!1,O=z?t(b.container).first():D?y:t("body"),L=function(e){if(M&&(f={left:e.pageX,top:e.pageY},!(i.abs(f.left-p.left)<w&&i.abs(f.top-p.top)<w))){if(null===s){var n=O.css("position");"absolute"!=n&&"relative"!=n&&"fixed"!=n&&(d=n,O.css("position","relative")),s=I.clone().removeClass("drag-from").addClass("drag-shadow").css({position:"absolute",width:I.outerWidth(),transition:"none"}).appendTo(O),I.addClass("dragging"),v.trigger("start",{event:e,element:I,shadowElement:s,targets:o})}var c={left:f.left-h.left,top:f.top-h.top},m={left:c.left-u.left,top:c.top-u.top};s.css(m),t.extend(g,f);var y=!1;r=!1,k||o.removeClass(P);var C=null;if(o.each(function(){var e=t(this),i=e.offset(),n=e.outerWidth(),o=e.outerHeight(),a=i.left+b.sensorOffsetX,s=i.top+b.sensorOffsetY;if(f.left>a&&f.top>s&&f.left<a+n&&f.top<s+o&&(C&&C.removeClass(P),C=e,!b.nested))return!1}),C){r=!0;var x=C.data("id");I.data("id")!=x&&(l=!1),(null===a||a.data("id")!==x&&!l)&&(y=!0),a=C,k&&o.removeClass(P),a.addClass(P)}k?null!==a&&a.length&&(r=!0):(I.toggleClass("drop-in",r),s.toggleClass("drop-in",r)),E&&E(I,a)===!1||v.trigger("drag",{event:e,isIn:r,target:a,element:I,isNew:y,selfTarget:l,clickOffset:h,offset:c,position:{left:c.left-u.left,top:c.top-u.top},mouseOffset:f}),e.preventDefault()}},j=function(i){if(t(e).off(C),clearTimeout(m),M){if(M=!1,d&&O.css("position",d),null===s)return I.removeClass("drag-from"),void v.trigger("always",{event:i,cancel:!0});r||(a=null);var n=!0;f=i?{left:i.pageX,top:i.pageY}:g;var c={left:f.left-h.left,top:f.top-h.top},p={left:f.left-g.left,top:f.top-g.top};g.left=f.left,g.top=f.top;var y={event:i,isIn:r,target:a,element:I,isNew:!l&&null!==a,selfTarget:l,offset:c,mouseOffset:f,position:{left:c.left-u.left,top:c.top-u.top},lastMouseOffset:g,moveOffset:p};n=v.trigger("beforeDrop",y),n&&r&&v.trigger("drop",y),o.removeClass(P),I.removeClass("dragging").removeClass("drag-from"),s.remove(),s=null,v.trigger("finish",y),v.trigger("always",y),i&&i.preventDefault()}},A=function(i){var n=t.zui.getMouseButtonCode(b.mouseButton);if(!(n>-1&&i.button!==n)){var f=t(this);D&&(I=S?f.closest(D):f),I.hasClass("drag-shadow")||b.before&&b.before({event:i,element:I})===!1||(M=!0,o=t.isFunction(b.target)?b.target(I,y):O.find(b.target),a=null,s=null,r=!1,l=!0,d=null,c=I.offset(),u=O.offset(),u.top=u.top-O.scrollTop(),u.left=u.left-O.scrollLeft(),p={left:i.pageX,top:i.pageY},g=t.extend({},p),h={left:p.left-c.left,top:p.top-c.top},I.addClass("drag-from"),t(e).on(T,L).on($,j),m=setTimeout(function(){t(e).on(x,j)},10),i.preventDefault(),b.stopPropagation&&i.stopPropagation())}};S?y.on(x,S,A):D?y.on(x,D,A):y.on(x,A)},s.prototype.destroy=function(){var i="."+n+"."+this.id;this.$.off(i),t(e).off(i),this.$.data(n,null)},s.prototype.reset=function(){this.destroy(),this.init()},t.fn.droppable=function(e){return this.each(function(){var i=t(this),o=i.data(n),a="object"==typeof e&&e;o||i.data(n,o=new s(this,a)),"string"==typeof e&&o[e]()})},t.fn.droppable.Constructor=s}(jQuery,document,Math),+function(t,e){"use strict";function i(e,i,a){return this.each(function(){var s=t(this),r=s.data(n),l=t.extend({},o.DEFAULTS,s.data(),"object"==typeof e&&e);r||s.data(n,r=new o(this,l)),"string"==typeof e?r[e](i,a):l.show&&r.show(i,a)})}var n="zui.modal",o=function(i,o){var a=this;a.options=o,a.$body=t(document.body),a.$element=t(i),a.$backdrop=a.isShown=null,a.scrollbarWidth=0,o.moveable===e&&(a.options.moveable=a.$element.hasClass("modal-moveable")),o.remote&&a.$element.find(".modal-content").load(o.remote,function(){a.$element.trigger("loaded."+n)}),o.scrollInside&&t(window).on("resize."+n,function(){a.isShown&&a.adjustPosition()})};o.VERSION="3.2.0",o.TRANSITION_DURATION=300,o.BACKDROP_TRANSITION_DURATION=150,o.DEFAULTS={backdrop:!0,keyboard:!0,show:!0,position:"fit"};var a=function(e,i){var n=t(window);i.left=Math.max(0,Math.min(i.left,n.width()-e.outerWidth())),i.top=Math.max(0,Math.min(i.top,n.height()-e.outerHeight())),e.css(i)};o.prototype.toggle=function(t,e){return this.isShown?this.hide():this.show(t,e)},o.prototype.adjustPosition=function(i){var o=this,s=o.options;if(i===e&&(i=s.position),i!==e&&null!==i){t.isFunction(i)&&(i=i(o));var r=o.$element.find(".modal-dialog"),l=t(window).height(),d={maxHeight:"initial",overflow:"visible"},c=r.find(".modal-body").css(d);if(s.scrollInside&&c.length){var p=s.headerHeight,u=s.footerHeight,h=r.find(".modal-header"),f=r.find(".modal-footer");p="number"!=typeof p&&h.length?h.outerHeight():t.isFunction(p)?p(h):0,u="number"!=typeof u&&f.length?f.outerHeight():t.isFunction(u)?u(f):0,d.maxHeight=l-p-u,d.overflow=c[0].scrollHeight>d.maxHeight?"auto":"visible",c.css(d)}var g=Math.max(0,(l-r.outerHeight())/2);if("fit"===i?i={top:g>50?Math.floor(2*g/3):g}:"center"===i?i={top:g}:t.isPlainObject(i)||(i={top:i}),r.hasClass("modal-moveable")){var m=null,v=s.rememberPos;v&&(v===!0?m=o.$element.data("modal-pos"):t.zui.store&&(m=t.zui.store.pageGet(n+".rememberPos."+v))),i=t.extend(i,{left:Math.max(0,(t(window).width()-r.outerWidth())/2)},m),"inside"===s.moveable?a(r,i):r.css(i)}else r.css(i)}},o.prototype.setMoveable=function(){t.fn.draggable||console.error("Moveable modal requires draggable.js.");var e=this,i=e.options,o=e.$element.find(".modal-dialog").removeClass("modal-dragged");o.toggleClass("modal-moveable",!!i.moveable),e.$element.data("modal-moveable-setup")||o.draggable({container:e.$element,handle:".modal-header",before:function(){var t=o.css("margin-top");t&&"0px"!==t&&o.css("top",t).css("margin-top","").addClass("modal-dragged")},finish:function(o){var a=i.rememberPos;a&&(e.$element.data("modal-pos",o.pos),t.zui.store&&a!==!0&&t.zui.store.pageSet(n+".rememberPos."+a,o.pos))},move:"inside"!==i.moveable||function(t){a(o,t)}})},o.prototype.show=function(e,i){var a=this,s=t.Event("show."+n,{relatedTarget:e});a.$element.trigger(s),a.$element.toggleClass("modal-scroll-inside",!!a.options.scrollInside),a.isShown||s.isDefaultPrevented()||(a.isShown=!0,a.options.moveable&&a.setMoveable(),a.options.backdrop!==!1&&(a.$body.addClass("modal-open"),a.setScrollbar()),a.escape(),a.$element.on("click.dismiss."+n,'[data-dismiss="modal"]',function(t){a.hide(),t.stopPropagation()}),a.backdrop(function(){var s=t.support.transition&&a.$element.hasClass("fade");a.$element.parent().length||a.$element.appendTo(a.$body),a.$element.show().scrollTop(0),s&&a.$element[0].offsetWidth,a.$element.addClass("in").attr("aria-hidden",!1),a.adjustPosition(i),a.enforceFocus();var r=t.Event("shown."+n,{relatedTarget:e});s?a.$element.find(".modal-dialog").one("bsTransitionEnd",function(){a.$element.trigger("focus").trigger(r)}).emulateTransitionEnd(o.TRANSITION_DURATION):a.$element.trigger("focus").trigger(r)}))},o.prototype.hide=function(e){e&&e.preventDefault&&e.preventDefault();var i=this;e=t.Event("hide."+n),i.$element.trigger(e),i.isShown&&!e.isDefaultPrevented()&&(i.isShown=!1,i.options.backdrop!==!1&&(i.$body.removeClass("modal-open"),i.resetScrollbar()),i.escape(),t(document).off("focusin."+n),i.$element.removeClass("in").attr("aria-hidden",!0).off("click.dismiss."+n),t.support.transition&&i.$element.hasClass("fade")?i.$element.one("bsTransitionEnd",t.proxy(i.hideModal,i)).emulateTransitionEnd(o.TRANSITION_DURATION):i.hideModal())},o.prototype.enforceFocus=function(){t(document).off("focusin."+n).on("focusin."+n,t.proxy(function(t){this.$element[0]===t.target||this.$element.has(t.target).length||this.$element.trigger("focus")},this))},o.prototype.escape=function(){this.isShown&&this.options.keyboard?t(document).on("keydown.dismiss."+n,t.proxy(function(i){if(27==i.which){var o=t.Event("escaping."+n),a=this.$element.triggerHandler(o,"esc");if(a!=e&&!a)return;this.hide()}},this)):this.isShown||t(document).off("keydown.dismiss."+n)},o.prototype.hideModal=function(){var t=this;this.$element.hide(),this.backdrop(function(){t.$element.trigger("hidden."+n)})},o.prototype.removeBackdrop=function(){this.$backdrop&&this.$backdrop.remove(),this.$backdrop=null},o.prototype.backdrop=function(e){var i=this,a=this.$element.hasClass("fade")?"fade":"";if(this.isShown&&this.options.backdrop){var s=t.support.transition&&a;if(this.$backdrop=t('<div class="modal-backdrop '+a+'" />').appendTo(this.$body),this.$element.on("mousedown.dismiss."+n,t.proxy(function(t){t.target===t.currentTarget&&("static"==this.options.backdrop?this.$element[0].focus.call(this.$element[0]):this.hide.call(this))},this)),s&&this.$backdrop[0].offsetWidth,this.$backdrop.addClass("in"),!e)return;s?this.$backdrop.one("bsTransitionEnd",e).emulateTransitionEnd(o.BACKDROP_TRANSITION_DURATION):e()}else if(!this.isShown&&this.$backdrop){this.$backdrop.removeClass("in");var r=function(){i.removeBackdrop(),e&&e()};t.support.transition&&this.$element.hasClass("fade")?this.$backdrop.one("bsTransitionEnd",r).emulateTransitionEnd(o.BACKDROP_TRANSITION_DURATION):r()}else e&&e()},o.prototype.setScrollbar=function(){t.zui.fixBodyScrollbar()&&this.options.onSetScrollbar&&this.options.onSetScrollbar(paddingRight)},o.prototype.resetScrollbar=function(){t.zui.resetBodyScrollbar(),this.options.onSetScrollbar&&this.options.onSetScrollbar("")},o.prototype.measureScrollbar=function(){var t=document.createElement("div");t.className="modal-scrollbar-measure",this.$body.append(t);var e=t.offsetWidth-t.clientWidth;return this.$body[0].removeChild(t),e};var s=t.fn.modal;t.fn.modal=i,t.fn.modal.Constructor=o,t.fn.modal.noConflict=function(){return t.fn.modal=s,this},t(document).on("click."+n+".data-api",'[data-toggle="modal"]',function(e){var o=t(this),a=o.attr("href"),s=null;try{s=t(o.attr("data-target")||a&&a.replace(/.*(?=#[^\s]+$)/,""))}catch(r){return}if(s.length){var l=s.data(n)?"toggle":t.extend({remote:!/#/.test(a)&&a},s.data(),o.data());o.is("a")&&e.preventDefault(),s.one("show."+n,function(t){t.isDefaultPrevented()||s.one("hidden."+n,function(){o.is(":visible")&&o.trigger("focus")})}),i.call(s,l,this,o.data("position"))}})}(jQuery,void 0),function(t,e,i){"use strict";if(!t.fn.modal)throw new Error("Modal trigger requires modal.js");var n="zui.modaltrigger",o="ajax",a=".zui.modal",s="string",r=function(e,i){e=t.extend({},r.DEFAULTS,t.ModalTriggerDefaults,i?i.data():null,e),this.isShown,this.$trigger=i,this.options=e,this.id=t.zui.uuid()};r.DEFAULTS={type:"custom",height:"auto",name:"triggerModal",fade:!0,position:"fit",showHeader:!0,delay:0,backdrop:!0,keyboard:!0,waittime:0,loadingIcon:"icon-spinner-indicator",scrollInside:!1},r.prototype.initOptions=function(i){if(i.url&&(!i.type||i.type!=o&&"iframe"!=i.type)&&(i.type=o),i.remote)i.type=o,typeof i.remote===s&&(i.url=i.remote);else if(i.iframe)i.type="iframe",typeof i.iframe===s&&(i.url=i.iframe);else if(i.custom&&(i.type="custom",typeof i.custom===s)){var n;try{n=t(i.custom)}catch(a){}n&&n.length?i.custom=n:t.isFunction(e[i.custom])&&(i.custom=e[i.custom])}return i},r.prototype.init=function(e){var i=this,o=t("#"+e.name);o.length&&(i.isShown||o.off(a),o.remove()),o=t('<div id="'+e.name+'" class="modal modal-trigger '+(e.className||"")+'">'+("string"==typeof e.loadingIcon&&0===e.loadingIcon.indexOf("icon-")?'<div class="icon icon-spin loader '+e.loadingIcon+'"></div>':e.loadingIcon)+'<div class="modal-dialog"><div class="modal-content"><div class="modal-header"><button class="close" data-dismiss="modal">×</button><h4 class="modal-title"><i class="modal-icon"></i> <span class="modal-title-name"></span></h4></div><div class="modal-body"></div></div></div></div>').appendTo("body").data(n,i);var s=function(i,n){var s=e[i];t.isFunction(s)&&o.on(n+a,s)};s("onShow","show"),s("shown","shown"),s("onHide","hide"),s("hidden","hidden"),s("loaded","loaded"),o.on("shown"+a,function(){i.isShown=!0}).on("hidden"+a,function(){i.isShown=!1}),this.$modal=o,this.$dialog=o.find(".modal-dialog"),e.mergeOptions&&(this.options=e)},r.prototype.show=function(i){var a=this,l=t.extend({},r.DEFAULTS,a.options,{url:a.$trigger?a.$trigger.attr("href")||a.$trigger.attr("data-url")||a.$trigger.data("url"):a.options.url},i),d=a.isShown;l=a.initOptions(l),d||a.init(l);var c=a.$modal,p=c.find(".modal-dialog"),u=l.custom,h=p.find(".modal-body").css("padding","").toggleClass("load-indicator loading",!!d),f=p.find(".modal-header"),g=p.find(".modal-content");c.toggleClass("fade",l.fade).addClass(l.className).toggleClass("modal-loading",!d).toggleClass("modal-scroll-inside",!!l.scrollInside),p.toggleClass("modal-md","md"===l.size).toggleClass("modal-sm","sm"===l.size).toggleClass("modal-lg","lg"===l.size).toggleClass("modal-fullscreen","fullscreen"===l.size),f.toggle(l.showHeader),f.find(".modal-icon").attr("class","modal-icon icon-"+l.icon),f.find(".modal-title-name").text(l.title||""),l.size&&"fullscreen"===l.size&&(l.width="",l.height="");var m=function(){clearTimeout(this.resizeTask),this.resizeTask=setTimeout(function(){a.adjustPosition(l.position)},100)},v=function(t,e){return"undefined"==typeof t&&(t=l.delay),setTimeout(function(){p=c.find(".modal-dialog"),l.width&&"auto"!=l.width&&p.css("width",l.width),l.height&&"auto"!=l.height&&(p.css("height",l.height),"iframe"===l.type&&h.css("height",p.height()-f.outerHeight())),a.adjustPosition(l.position),c.removeClass("modal-loading").removeClass("modal-updating"),d&&h.removeClass("loading"),"iframe"!=l.type&&(h=p.off("resize."+n).find(".modal-body").off("resize."+n),(h.length?h:p).on("resize."+n,m)),e&&e()},t)};if("custom"===l.type&&u)if(t.isFunction(u)){var y=u({modal:c,options:l,modalTrigger:a,ready:v});typeof y===s&&(h.html(y),v())}else u instanceof t?(h.html(t("<div>").append(u.clone()).html()),v()):(h.html(u),v());else if(l.url){var b=function(){var t=c.callComEvent(a,"broken");"string"==typeof t&&h.html(t),v()};if(c.attr("ref",l.url),"iframe"===l.type){c.addClass("modal-iframe"),this.firstLoad=!0;var w="iframe-"+l.name;f.detach(),h.detach(),g.empty().append(f).append(h),h.css("padding",0).html('<iframe id="'+w+'" name="'+w+'" src="'+l.url+'" frameborder="no"  allowfullscreen="true" mozallowfullscreen="true" webkitallowfullscreen="true"  allowtransparency="true" scrolling="auto" style="width: 100%; height: 100%; left: 0px;"></iframe>'),l.waittime>0&&(a.waitTimeout=v(l.waittime,b));var C=document.getElementById(w);C.onload=C.onreadystatechange=function(i){var o=!!l.scrollInside;if(a.firstLoad&&c.addClass("modal-loading"),!this.readyState||"complete"==this.readyState){a.firstLoad=!1,l.waittime>0&&clearTimeout(a.waitTimeout);try{c.attr("ref",C.contentWindow.location.href);var s=e.frames[w].$;if(s&&"auto"===l.height&&"fullscreen"!=l.size){var r=s("body").addClass("body-modal").toggleClass("body-modal-scroll-inside",o);l.iframeBodyClass&&r.addClass(l.iframeBodyClass);var d=[],p=function(i){c.removeClass("fade");var n=r.outerHeight();if(i===!0&&l.onlyIncreaseHeight&&(n=Math.max(n,h.data("minModalHeight")||0),h.data("minModalHeight",n)),o){var a=l.headerHeight;"number"!=typeof a?a=f.outerHeight():t.isFunction(a)&&(a=a(f));var s=t(e).height();n=Math.min(n,s-a)}for(d.length>1&&n===d[0]&&(n=Math.max(n,d[1])),d.push(n);d.length>2;)d.shift();h.css("height",n),l.fade&&c.addClass("fade"),v()};c.callComEvent(a,"loaded",{modalType:"iframe",jQuery:s}),setTimeout(p,100),r.off("resize."+n).on("resize."+n,p),o&&t(e).off("resize."+n).on("resize."+n,p)}else v();var u=l.handleLinkInIframe;u&&s("body").on("click","string"==typeof u?u:"a[href]",function(){t(this).is('[data-toggle="modal"]')||c.addClass("modal-updating")}),l.iframeStyle&&s("head").append("<style>"+l.iframeStyle+"</style>")}catch(i){v()}}}}else t.ajax(t.extend({url:l.url,success:function(i){try{var s=t(i);s.filter(".modal-dialog").length?p.parent().empty().append(s):s.filter(".modal-content").length?p.find(".modal-content").replaceWith(s):h.wrapInner(s)}catch(r){e.console&&e.console.warn&&console.warn("ZUI: Cannot recogernize remote content.",{error:r,data:i}),c.html(i)}c.callComEvent(a,"loaded",{modalType:o}),v(),l.scrollInside&&t(e).off("resize."+n).on("resize."+n,m)},error:b},l.ajaxOptions))}d||c.modal({show:"show",backdrop:l.backdrop,moveable:l.moveable,rememberPos:l.rememberPos,keyboard:l.keyboard,scrollInside:l.scrollInside})},r.prototype.close=function(i,n){var o=this;(i||n)&&o.$modal.on("hidden"+a,function(){t.isFunction(i)&&i(),typeof n===s&&n.length&&!o.$modal.data("cancel-reload")&&("this"===n?e.location.reload():e.location=n)}),o.$modal.modal("hide")},r.prototype.toggle=function(t){this.isShown?this.close():this.show(t)},r.prototype.adjustPosition=function(e){e=e===i?this.options.position:e,t.isFunction(e)&&(e=e(this)),this.$modal.modal("adjustPosition",e)},t.zui({ModalTrigger:r,modalTrigger:new r}),t.fn.modalTrigger=function(e,i){return t(this).each(function(){var o=t(this),a=o.data(n),l=t.extend({title:o.attr("title")||o.text(),url:o.attr("href"),type:o.hasClass("iframe")?"iframe":""},o.data(),t.isPlainObject(e)&&e);a||o.data(n,a=new r(l,o)),typeof e==s?a[e](i):l.show&&a.show(i),o.on((l.trigger||"click")+".toggle."+n,function(e){l=t.extend(l,{url:o.attr("href")||o.attr("data-url")||o.data("url")||l.url}),a.toggle(l),o.is("a")&&e.preventDefault()})})};var l=t.fn.modal;t.fn.modal=function(e,i){return t(this).each(function(){var n=t(this);n.hasClass("modal")?l.call(n,e,i):n.modalTrigger(e,i)})},t.fn.modal.bs=l;var d=function(e){return e?e=t(e):(e=t(".modal.modal-trigger"),!e.length),e&&e instanceof t?e:null},c=function(i,o,a){var s=i;if(t.isFunction(i)){var r=a;a=o,o=i,i=r}i=d(i),i&&i.length?i.each(function(){t(this).data(n).close(o,a)}):t("body").hasClass("modal-open")||t(".modal.in").length||t("body").hasClass("body-modal")&&e.parent.$.zui.closeModal(s,o,a)},p=function(t,e){e=d(e),e&&e.length&&e.modal("adjustPosition",t)},u=function(e,i){"string"==typeof e&&(e={url:e});var o=d(i);o&&o.length&&o.each(function(){t(this).data(n).show(e)})};t.zui({reloadModal:u,closeModal:c,ajustModalPosition:p,adjustModalPosition:p}),t(document).on("click."+n+".data-api",'[data-toggle="modal"]',function(e){var i=t(this),o=i.attr("href"),a=null;try{a=t(i.attr("data-target")||o&&o.replace(/.*(?=#[^\s]+$)/,""))}catch(s){}a&&a.length||(i.data(n)?i.trigger(".toggle."+n):i.modalTrigger({show:!0})),i.is("a")&&e.preventDefault()}).on("click."+n+".data-api",'[data-dismiss="modal"]',function(){t.zui.closeModal()})}(window.jQuery,window,void 0),+function(t){"use strict";var e=function(t,e){this.type=null,this.options=null,this.enabled=null,this.timeout=null,this.hoverState=null,this.$element=null,this.init("tooltip",t,e)};e.DEFAULTS={animation:!0,placement:"top",selector:!1,template:'<div class="tooltip"><div class="tooltip-arrow"></div><div class="tooltip-inner"></div></div>',trigger:"hover focus",title:"",delay:0,html:!1,container:!1},e.prototype.init=function(e,i,n){this.enabled=!0,this.type=e,this.$element=t(i),this.options=this.getOptions(n);for(var o=this.options.trigger.split(" "),a=o.length;a--;){var s=o[a];if("click"==s)this.$element.on("click."+this.type,this.options.selector,t.proxy(this.toggle,this));else if("manual"!=s){var r="hover"==s?"mouseenter":"focus",l="hover"==s?"mouseleave":"blur";this.$element.on(r+"."+this.type,this.options.selector,t.proxy(this.enter,this)),this.$element.on(l+"."+this.type,this.options.selector,t.proxy(this.leave,this))}}this.options.selector?this._options=t.extend({},this.options,{trigger:"manual",selector:""}):this.fixTitle()},e.prototype.getDefaults=function(){return e.DEFAULTS},e.prototype.getOptions=function(e){return e=t.extend({},this.getDefaults(),this.$element.data(),e),e.delay&&"number"==typeof e.delay&&(e.delay={show:e.delay,hide:e.delay}),e},e.prototype.getDelegateOptions=function(){var e={},i=this.getDefaults();return this._options&&t.each(this._options,function(t,n){i[t]!=n&&(e[t]=n)}),e},e.prototype.enter=function(e){var i=e instanceof this.constructor?e:t(e.currentTarget)[this.type](this.getDelegateOptions()).data("zui."+this.type);return clearTimeout(i.timeout),i.hoverState="in",i.options.delay&&i.options.delay.show?void(i.timeout=setTimeout(function(){"in"==i.hoverState&&i.show()},i.options.delay.show)):i.show()},e.prototype.leave=function(e){var i=e instanceof this.constructor?e:t(e.currentTarget)[this.type](this.getDelegateOptions()).data("zui."+this.type);return clearTimeout(i.timeout),i.hoverState="out",i.options.delay&&i.options.delay.hide?void(i.timeout=setTimeout(function(){"out"==i.hoverState&&i.hide()},i.options.delay.hide)):i.hide()},e.prototype.show=function(e){var i=t.Event("show.zui."+this.type);if((e||this.hasContent())&&this.enabled){var n=this;if(n.$element.trigger(i),i.isDefaultPrevented())return;var o=n.tip();n.setContent(e),n.options.animation&&o.addClass("fade");var a="function"==typeof n.options.placement?n.options.placement.call(n,o[0],n.$element[0]):n.options.placement,s=/\s?auto?\s?/i,r=s.test(a);r&&(a=a.replace(s,"")||"top"),o.detach().css({top:0,left:0,display:"block"}).addClass(a),n.options.container?o.appendTo(n.options.container):o.insertAfter(n.$element);var l=n.getPosition(),d=o[0].offsetWidth,c=o[0].offsetHeight;if(r){var p=n.$element.parent(),u=a,h=document.documentElement.scrollTop||document.body.scrollTop,f="body"==n.options.container?window.innerWidth:p.outerWidth(),g="body"==n.options.container?window.innerHeight:p.outerHeight(),m="body"==n.options.container?0:p.offset().left;a="bottom"==a&&l.top+l.height+c-h>g?"top":"top"==a&&l.top-h-c<0?"bottom":"right"==a&&l.right+d>f?"left":"left"==a&&l.left-d<m?"right":a,o.removeClass(u).addClass(a)}var v=n.getCalculatedOffset(a,l,d,c);n.applyPlacement(v,a);var y=function(){var t=n.hoverState;n.$element.trigger("shown.zui."+n.type),n.hoverState=null,"out"==t&&n.leave(n)};t.support.transition&&n.$tip.hasClass("fade")?o.one("bsTransitionEnd",y).emulateTransitionEnd(150):y()}},e.prototype.applyPlacement=function(t,e){var i,n=this.tip(),o=n[0].offsetWidth,a=n[0].offsetHeight,s=parseInt(n.css("margin-top"),10),r=parseInt(n.css("margin-left"),10);isNaN(s)&&(s=0),isNaN(r)&&(r=0),t.top=t.top+s,t.left=t.left+r,n.offset(t).addClass("in");var l=n[0].offsetWidth,d=n[0].offsetHeight;if("top"==e&&d!=a&&(i=!0,t.top=t.top+a-d),/bottom|top/.test(e)){var c=0;t.left<0&&(c=t.left*-2,t.left=0,n.offset(t),l=n[0].offsetWidth,d=n[0].offsetHeight),this.replaceArrow(c-o+l,l,"left")}else this.replaceArrow(d-a,d,"top");i&&n.offset(t)},e.prototype.replaceArrow=function(t,e,i){this.arrow().css(i,t?50*(1-t/e)+"%":"")},e.prototype.setContent=function(t){var e=this.tip(),i=t||this.getTitle();this.options.tipId&&e.attr("id",this.options.tipId),this.options.tipClass&&e.addClass(this.options.tipClass),
e.find(".tooltip-inner")[this.options.html?"html":"text"](i),e.removeClass("fade in top bottom left right")},e.prototype.hide=function(){function e(){"in"!=i.hoverState&&n.detach()}var i=this,n=this.tip(),o=t.Event("hide.zui."+this.type);if(this.$element.trigger(o),!o.isDefaultPrevented())return n.removeClass("in"),t.support.transition&&this.$tip.hasClass("fade")?n.one(t.support.transition.end,e).emulateTransitionEnd(150):e(),this.$element.trigger("hidden.zui."+this.type),this},e.prototype.fixTitle=function(){var t=this.$element;(t.attr("title")||"string"!=typeof t.attr("data-original-title"))&&t.attr("data-original-title",t.attr("title")||"").attr("title","")},e.prototype.hasContent=function(){return this.getTitle()},e.prototype.getPosition=function(){var e=this.$element[0];return t.extend({},"function"==typeof e.getBoundingClientRect?e.getBoundingClientRect():{width:e.offsetWidth,height:e.offsetHeight},this.$element.offset())},e.prototype.getCalculatedOffset=function(t,e,i,n){return"bottom"==t?{top:e.top+e.height,left:e.left+e.width/2-i/2}:"top"==t?{top:e.top-n,left:e.left+e.width/2-i/2}:"left"==t?{top:e.top+e.height/2-n/2,left:e.left-i}:{top:e.top+e.height/2-n/2,left:e.left+e.width}},e.prototype.getTitle=function(){var t,e=this.$element,i=this.options;return t=e.attr("data-original-title")||("function"==typeof i.title?i.title.call(e[0]):i.title)},e.prototype.tip=function(){return this.$tip=this.$tip||t(this.options.template)},e.prototype.arrow=function(){return this.$arrow=this.$arrow||this.tip().find(".tooltip-arrow")},e.prototype.validate=function(){this.$element[0].parentNode||(this.hide(),this.$element=null,this.options=null)},e.prototype.enable=function(){this.enabled=!0},e.prototype.disable=function(){this.enabled=!1},e.prototype.toggleEnabled=function(){this.enabled=!this.enabled},e.prototype.toggle=function(e){var i=e?t(e.currentTarget)[this.type](this.getDelegateOptions()).data("zui."+this.type):this;i.tip().hasClass("in")?i.leave(i):i.enter(i)},e.prototype.destroy=function(){this.hide().$element.off("."+this.type).removeData("zui."+this.type)};var i=t.fn.tooltip;t.fn.tooltip=function(i,n){return this.each(function(){var o=t(this),a=o.data("zui.tooltip"),s="object"==typeof i&&i;a||o.data("zui.tooltip",a=new e(this,s)),"string"==typeof i&&a[i](n)})},t.fn.tooltip.Constructor=e,t.fn.tooltip.noConflict=function(){return t.fn.tooltip=i,this}}(window.jQuery),+function(t){"use strict";var e=function(t,e){this.init("popover",t,e)};if(!t.fn.tooltip)throw new Error("Popover requires tooltip.js");e.DEFAULTS=t.extend({},t.fn.tooltip.Constructor.DEFAULTS,{placement:"right",trigger:"click",content:"",template:'<div class="popover"><div class="arrow"></div><h3 class="popover-title"></h3><div class="popover-content"></div></div>'}),e.prototype=t.extend({},t.fn.tooltip.Constructor.prototype),e.prototype.constructor=e,e.prototype.getDefaults=function(){return e.DEFAULTS},e.prototype.setContent=function(){var t=this.tip(),e=this.getTarget();if(e)return e.find(".arrow").length<1&&t.addClass("no-arrow"),void t.html(e.html());var i=this.getTitle(),n=this.getContent();t.find(".popover-title")[this.options.html?"html":"text"](i),t.find(".popover-content")[this.options.html?"html":"text"](n),t.removeClass("fade top bottom left right in"),this.options.tipId&&t.attr("id",this.options.tipId),this.options.tipClass&&t.addClass(this.options.tipClass),t.find(".popover-title").html()||t.find(".popover-title").hide()},e.prototype.hasContent=function(){return this.getTarget()||this.getTitle()||this.getContent()},e.prototype.getContent=function(){var t=this.$element,e=this.options;return t.attr("data-content")||("function"==typeof e.content?e.content.call(t[0]):e.content)},e.prototype.getTarget=function(){var e=this.$element,i=this.options,n=e.attr("data-target")||("function"==typeof i.target?i.target.call(e[0]):i.target);return!!n&&("$next"==n?e.next(".popover"):t(n))},e.prototype.arrow=function(){return this.$arrow=this.$arrow||this.tip().find(".arrow")},e.prototype.tip=function(){return this.$tip||(this.$tip=t(this.options.template)),this.$tip};var i=t.fn.popover;t.fn.popover=function(i){return this.each(function(){var n=t(this),o=n.data("zui.popover"),a="object"==typeof i&&i;o||n.data("zui.popover",o=new e(this,a)),"string"==typeof i&&o[i]()})},t.fn.popover.Constructor=e,t.fn.popover.noConflict=function(){return t.fn.popover=i,this}}(window.jQuery),+function(t){"use strict";function e(){t(o).remove(),t(a).each(function(e){var o=i(t(this));o.hasClass("open")&&(o.trigger(e=t.Event("hide."+n)),e.isDefaultPrevented()||o.removeClass("open").trigger("hidden."+n))})}function i(e){var i=e.attr("data-target");i||(i=e.attr("href"),i=i&&/#/.test(i)&&i.replace(/.*(?=#[^\s]*$)/,""));var n;try{n=i&&t(i)}catch(o){}return n&&n.length?n:e.parent()}var n="zui.dropdown",o=".dropdown-backdrop",a="[data-toggle=dropdown]",s=function(e){t(e).on("click."+n,this.toggle)};s.prototype.toggle=function(o){var a=t(this);if(!a.is(".disabled, :disabled")){var s=i(a),r=s.hasClass("open");if(e(),!r){if("ontouchstart"in document.documentElement&&!s.closest(".navbar-nav").length&&t('<div class="dropdown-backdrop"/>').insertAfter(t(this)).on("click",e),s.trigger(o=t.Event("show."+n)),o.isDefaultPrevented())return;s.toggleClass("open").trigger("shown."+n),a.focus()}return!1}},s.prototype.keydown=function(e){if(/(38|40|27)/.test(e.keyCode)){var n=t(this);if(e.preventDefault(),e.stopPropagation(),!n.is(".disabled, :disabled")){var o=i(n),s=o.hasClass("open");if(!s||s&&27==e.keyCode)return 27==e.which&&o.find(a).focus(),n.click();var r=t("[role=menu] li:not(.divider):visible a",o);if(r.length){var l=r.index(r.filter(":focus"));38==e.keyCode&&l>0&&l--,40==e.keyCode&&l<r.length-1&&l++,~l||(l=0),r.eq(l).focus()}}}};var r=t.fn.dropdown;t.fn.dropdown=function(e){return this.each(function(){var i=t(this),n=i.data("dropdown");n||i.data("dropdown",n=new s(this)),"string"==typeof e&&n[e].call(i)})},t.fn.dropdown.Constructor=s,t.fn.dropdown.noConflict=function(){return t.fn.dropdown=r,this};var l=n+".data-api";t(document).on("click."+l,e).on("click."+l,".dropdown form",function(t){t.stopPropagation()}).on("click."+l,a,s.prototype.toggle).on("keydown."+l,a+", [role=menu]",s.prototype.keydown)}(window.jQuery),function(t,e){"use strict";var i="zui.contextmenu",n={duration:200},o=!1,a={},s="zui-contextmenu-"+t.zui.uuid(),r=0,l=0,d=function(){return t(document).off("mousemove."+i).on("mousemove."+i,function(t){r=t.clientX,l=t.clientY}),a},c=function(e,i){if("string"==typeof e&&(e="seperator"===e||"divider"===e||"-"===e||"|"===e?{type:"seperator"}:{label:e,id:i}),"seperator"===e.type||"divider"===e.type)return t('<li class="divider"></li>');var n=t("<a/>").attr({href:e.url||"###","class":e.className,style:e.style}).data("item",e);return e.html?e.html===!0?n.html(e.label||e.text):n=t(e.html):n.text(e.label||e.text),e.onClick&&n.on("click",e.onClick),t("<li />").toggleClass("disabled",e.disabled===!0).append(n)},p=null,u=function(e,i){"function"==typeof e&&(i=e,e=null),p&&(clearTimeout(p),p=null);var n=t("#"+s);if(n.length){var o=n.data("options");if(!e||o.id===e){var r=function(){n.hide(),o.onHidden&&o.onHidden(),i&&i()};o.onHide&&o.onHide();var l=o.animation;n.removeClass("in"),l?p=setTimeout(r,o.duration):r()}}return a},h=function(d,h,f){t.isPlainObject(d)&&(f=h,h=d,d=h.items),o=!0,h=t.extend({},n,h);var g=h.x,m=h.y;g===e&&(g=(h.event||h).clientX),g===e&&(g=r),m===e&&(m=(h.event||h).clientY),m===e&&(m=l);var v=t("#"+s);v.length||(v=t('<div style="display: none; position: fixed; z-index: 2000;" class="contextmenu" id="'+s+'"><ul class="dropdown-menu contextmenu-menu"></ul></div>').appendTo("body"));var y=v.find(".contextmenu-menu").off("click."+i).on("click."+i,"a",function(e){var i=t(this),n=h.onClickItem&&h.onClickItem(i.data("item"),i,e);n!==!1&&u()}).empty();y.attr("class","dropdown-menu contextmenu-menu"+(h.className?" "+h.className:"")),v.hide().attr("class","contextmenu");var b=h.itemCreator||c,w=typeof d;"string"===w?d=d.split(","):"function"===w&&(d=d(h)),t.each(d,function(t,e){y.append(b(e,t,h))});var C=h.animation,x=h.duration;C===!0&&(h.animation=C="fade"),p&&(clearTimeout(p),p=null);var $=function(){v.addClass("in"),h.onShown&&h.onShown(),f&&f()};h.onShow&&h.onShow(),v.data("options",{animation:C,onHide:h.onHide,onHidden:h.onHidden,id:h.id,duration:x});var T=t(window);return g=Math.max(0,Math.min(g,T.width()-y.outerWidth())),m=Math.max(0,Math.min(m,T.height()-y.outerHeight())),v.css({left:g,top:m}),C?(v.addClass("open").addClass(C).show(),p=setTimeout(function(){$(),o=!1},h.duration)):(v.addClass("open").show(),$(),p=setTimeout(function(){o=!1},200)),a};t(document).on("click",function(e){o||t(e.target).closest(".contextmenu").length||u()}),t.extend(a,{NAME:i,DEFAULTS:n,show:h,hide:u,listenMouse:d}),t.zui({ContextMenu:a});var f=function(e,n){var o=this;o.name=i,o.$=t(e),n=o.options=t.extend({trigger:"contextmenu"},a.DEFAULTS,this.$.data(),n);var s=n.trigger;o.id=t.zui.uuid();var r=function(t){if("mousedown"!==t.type||2===t.button){var e={x:t.clientX,y:t.clientY,event:t};return n.itemsCreator&&(e.items=n.itemsCreator.call(this,t)),o.show(e),t.preventDefault(),t.returnValue=!1,!1}},l=s+"."+i;n.selector?o.$.on(l,n.selector,r):o.$.on(l,r)};f.prototype.destory=function(){that.$.off("."+i)},f.prototype.hide=function(t){a.hide(this.id,t)},f.prototype.show=function(e,i){e=t.extend({},this.options,e),a.show(e,i)},t.fn.contextmenu=function(e){return this.each(function(){var n=t(this),o=n.data(i),a="object"==typeof e&&e;o||n.data(i,o=new f(this,a)),"string"==typeof e&&o[e]()})},t.fn.contextmenu.Constructor=f}(jQuery,void 0),+function(t){"use strict";var e=function(e,i){this.$element=t(e),this.$indicators=this.$element.find(".carousel-indicators"),this.options=i,this.paused=this.sliding=this.interval=this.$active=this.$items=null,"hover"==this.options.pause&&this.$element.on("mouseenter",t.proxy(this.pause,this)).on("mouseleave",t.proxy(this.cycle,this))};e.DEFAULTS={interval:5e3,pause:"hover",wrap:!0,touchable:!0},e.prototype.touchable=function(){function e(e){var e=e||window.event;e.originalEvent&&(e=e.originalEvent);var a=t(this);switch(e.type){case"touchstart":n=e.touches[0].pageX,o=e.touches[0].pageY;break;case"touchend":var s=e.changedTouches[0].pageX-n,r=e.changedTouches[0].pageY-o;if(Math.abs(s)>Math.abs(r))i(a,s),Math.abs(s)>10&&e.preventDefault();else{var l=t(window);t("body,html").animate({scrollTop:l.scrollTop()-r},400)}}}function i(t,e){e>10?a.prev():e<-10&&a.next()}if(this.options.touchable){this.$element.on("touchstart touchmove touchend",e);var n,o,a=this}},e.prototype.cycle=function(e){return e||(this.paused=!1),this.interval&&clearInterval(this.interval),this.options.interval&&!this.paused&&(this.interval=setInterval(t.proxy(this.next,this),this.options.interval)),this},e.prototype.getActiveIndex=function(){return this.$active=this.$element.find(".item.active"),this.$items=this.$active.parent().children(),this.$items.index(this.$active)},e.prototype.to=function(e){var i=this,n=this.getActiveIndex();if(!(e>this.$items.length-1||e<0))return this.sliding?this.$element.one("slid",function(){i.to(e)}):n==e?this.pause().cycle():this.slide(e>n?"next":"prev",t(this.$items[e]))},e.prototype.pause=function(e){return e||(this.paused=!0),this.$element.find(".next, .prev").length&&t.support.transition.end&&(this.$element.trigger(t.support.transition.end),this.cycle(!0)),this.interval=clearInterval(this.interval),this},e.prototype.next=function(){if(!this.sliding)return this.slide("next")},e.prototype.prev=function(){if(!this.sliding)return this.slide("prev")},e.prototype.slide=function(e,i){var n=this.$element.find(".item.active"),o=i||n[e](),a=this.interval,s="next"==e?"left":"right",r="next"==e?"first":"last",l=this;if(!o.length){if(!this.options.wrap)return;o=this.$element.find(".item")[r]()}this.sliding=!0,a&&this.pause();var d=t.Event("slide.zui.carousel",{relatedTarget:o[0],direction:s});if(!o.hasClass("active")){if(this.$indicators.length&&(this.$indicators.find(".active").removeClass("active"),this.$element.one("slid",function(){var e=t(l.$indicators.children()[l.getActiveIndex()]);e&&e.addClass("active")})),t.support.transition&&this.$element.hasClass("slide")){if(this.$element.trigger(d),d.isDefaultPrevented())return;o.addClass(e),o[0].offsetWidth,n.addClass(s),o.addClass(s),n.one(t.support.transition.end,function(){o.removeClass([e,s].join(" ")).addClass("active"),n.removeClass(["active",s].join(" ")),l.sliding=!1,setTimeout(function(){l.$element.trigger("slid")},0)}).emulateTransitionEnd(600)}else{if(this.$element.trigger(d),d.isDefaultPrevented())return;n.removeClass("active"),o.addClass("active"),this.sliding=!1,this.$element.trigger("slid")}return a&&this.cycle(),this}};var i=t.fn.carousel;t.fn.carousel=function(i){return this.each(function(){var n=t(this),o=n.data("zui.carousel"),a=t.extend({},e.DEFAULTS,n.data(),"object"==typeof i&&i),s="string"==typeof i?i:a.slide;o||n.data("zui.carousel",o=new e(this,a)),"number"==typeof i?o.to(i):s?o[s]():a.interval&&o.pause().cycle(),a.touchable&&o.touchable()})},t.fn.carousel.Constructor=e,t.fn.carousel.noConflict=function(){return t.fn.carousel=i,this},t(document).on("click.zui.carousel.data-api","[data-slide], [data-slide-to]",function(e){var i,n=t(this),o=t(n.attr("data-target")||(i=n.attr("href"))&&i.replace(/.*(?=#[^\s]+$)/,"")),a=t.extend({},o.data(),n.data()),s=n.attr("data-slide-to");s&&(a.interval=!1),o.carousel(a),(s=n.attr("data-slide-to"))&&o.data("zui.carousel").to(s),e.preventDefault()}),t(window).on("load",function(){t('[data-ride="carousel"]').each(function(){var e=t(this);e.carousel(e.data())})})}(window.jQuery),/*! TangBin: image.ready.js http://www.planeart.cn/?p=1121 */
function(t){"use strict";t.zui.imgReady=function(){var t=[],e=null,i=function(){for(var e=0;e<t.length;e++)t[e].end?t.splice(e--,1):t[e]();!t.length&&n()},n=function(){clearInterval(e),e=null};return function(n,o,a,s){var r,l,d,c,p,u=new Image;return u.src=n,u.complete?(o.call(u),void(a&&a.call(u))):(l=u.width,d=u.height,u.onerror=function(){s&&s.call(u),r.end=!0,u=u.onload=u.onerror=null},r=function(){c=u.width,p=u.height,(c!==l||p!==d||c*p>1024)&&(o.call(u),r.end=!0)},r(),u.onload=function(){!r.end&&r(),a&&a.call(u),u=u.onload=u.onerror=null},void(r.end||(t.push(r),null===e&&(e=setInterval(i,40)))))}}()}(jQuery),function(t,e,i){"use strict";if(!t.fn.modalTrigger)throw new Error("modal & modalTrigger requires for lightbox");if(!t.zui.imgReady)throw new Error("imgReady requires for lightbox");var n=function(e,i){this.$=t(e),this.options=this.getOptions(i),this.init()};n.DEFAULTS={modalTeamplate:'<div class="icon-spinner icon-spin loader"></div><div class="modal-dialog"><button class="close" data-dismiss="modal" aria-hidden="true"><i class="icon-remove"></i></button><button class="controller prev"><i class="icon icon-chevron-left"></i></button><button class="controller next"><i class="icon icon-chevron-right"></i></button><img class="lightbox-img" src="{image}" alt="" data-dismiss="modal" /><div class="caption"><div class="content">{caption}<div></div></div>'},n.prototype.getOptions=function(e){var i="image";return e=t.extend({},n.DEFAULTS,this.$.data(),e),e[i]||(e[i]=this.$.attr("src")||this.$.attr("href")||this.$.find("img").attr("src"),this.$.data(i,e[i])),e},n.prototype.init=function(){this.bindEvents()},n.prototype.initGroups=function(){var e=this.$.data("groups");e||(e=t('[data-toggle="lightbox"][data-group="'+this.options.group+'"], [data-lightbox-group="'+this.options.group+'"]'),this.$.data("groups",e),e.each(function(e){t(this).attr("data-group-index",e)})),this.groups=e,this.groupIndex=parseInt(this.$.data("group-index"))},n.prototype.setImage=function(t,e){void 0!==t&&(this.options.image=t),void 0!==e&&(this.options.caption=e)},n.prototype.show=function(t,e){this.setImage(t,e),this.$.triggerHandler("click")},n.prototype.bindEvents=function(){var n=this.$,o=this,a=this.options;return!!a.image&&void n.modalTrigger({type:"custom",name:"lightboxModal",position:"center",custom:function(n){o.initGroups();var s=n.modal,r=o.groups,l=o.groupIndex;s.addClass("modal-lightbox").html(a.modalTeamplate.format(a)).toggleClass("lightbox-with-caption","string"==typeof a.caption).removeClass("lightbox-full").data("group-index",l);var d=s.find(".modal-dialog"),c=t(e).width();t.zui.imgReady(a.image,function(){d.css({width:i.min(c,this.width)}),c<this.width+30&&s.addClass("lightbox-full"),n.ready(200)}),s.find(".prev").toggleClass("show",r.filter('[data-group-index="'+(l-1)+'"]').length>0),s.find(".next").toggleClass("show",r.filter('[data-group-index="'+(l+1)+'"]').length>0),s.find(".controller").click(function(){var o=t(this),a=s.data("group-index")+(o.hasClass("prev")?-1:1),l=r.filter('[data-group-index="'+a+'"]');if(l.length){var p=l.data("image"),u=l.data("caption");s.addClass("modal-loading").data("group-index",a).toggleClass("lightbox-with-caption","string"==typeof u).removeClass("lightbox-full"),s.find(".lightbox-img").attr("src",p),s.find(".caption > .content").text(u),c=t(e).width(),t.zui.imgReady(p,function(){d.css({width:i.min(c,this.width)}),c<this.width+30&&s.addClass("lightbox-full"),n.ready()})}return s.find(".prev").toggleClass("show",r.filter('[data-group-index="'+(a-1)+'"]').length>0),s.find(".next").toggleClass("show",r.filter('[data-group-index="'+(a+1)+'"]').length>0),!1})}})},t.fn.lightbox=function(e){var i="group"+(new Date).getTime();return this.each(function(){var o=t(this),a="object"==typeof e&&e;"object"==typeof a&&a.group?o.attr("data-lightbox-group",a.group):o.data("group")?o.attr("data-lightbox-group",o.data("group")):o.attr("data-lightbox-group",i),o.data("group",o.data("lightbox-group"));var s=o.data("zui.lightbox");s||o.data("zui.lightbox",s=new n(this,a)),"string"==typeof e&&s[e]()})},t.fn.lightbox.Constructor=n,t(function(){t('[data-toggle="lightbox"]').lightbox()})}(jQuery,window,Math),function(t,e,i){"use strict";var n=0,o='<div class="messager messager-{type} {placement}" style="display: none"><div class="messager-content"></div><div class="messager-actions"></div></div>',a={icons:{},type:"default",placement:"top",time:4e3,parent:"body",close:!0,fade:!0,scale:!0},s={},r=function(e,r){t.isPlainObject(e)?r=t.extend({},r,e):e&&(r?r.content=e:r={content:e});var l=this;r=l.options=t.extend({},a,r),l.id=r.id||n++;var d=s[l.id];d&&d.destroy(),s[l.id]=l,l.$=t(o.format(r)).toggleClass("fade",r.fade).toggleClass("scale",r.scale).attr("id","messager-"+l.id),r.cssClass&&l.$.addClass(r.cssClass);var c=!1,p=l.$.find(".messager-actions"),u=function(e){var n=t('<button type="button" class="action action-'+e.name+'"/>');"close"===e.name&&n.addClass("close"),e.html!==i&&n.html(e.html),e.icon!==i&&n.append('<i class="action-icon icon-'+e.icon+'"/>'),e.text!==i&&n.append('<span class="action-text">'+e.text+"</span>"),e.tooltip!==i&&n.attr("title",e.tooltip).tooltip(),n.data("action",e),p.append(n)};r.actions&&t.each(r.actions,function(t,e){e.name===i&&(e.name=t),"close"==e.name&&(c=!0),u(e)}),!c&&r.close&&u({name:"close",html:"&times;"}),l.$.on("click",".action",function(e){var i,n=t(this).data("action");r.onAction&&(i=r.onAction.call(this,n.name,n,l),i===!1)||t.isFunction(n.action)&&(i=n.action.call(this,l),i===!1)||(l.hide(),e.stopPropagation())}),l.$.on("click",function(t){if(r.onAction){var e=r.onAction.call(this,"content",null,l);e===!0&&l.hide()}}),l.$.data("zui.messager",l),r.show&&l.message!==i&&l.show()};r.prototype.update=function(e,i){t.isPlainObject(e)?i=e:e&&(i?i.content=e:i={content:e});var n=this,o=n.options;n.$.removeClass("messager-"+o.type);var a=n.$.find(".messager-content");o.contentClass&&a.removeClass(o.contentClass),i&&(o=t.extend(o,i)),n.$.addClass("messager-"+o.type).toggleClass("messager-notification",!!o.notification),o.contentClass&&a.addClass(o.contentClass);var s=o.title,r=o.icon;if(e=o.content,a.empty(),s){var l=t('<div class="messager-title"></div>');l[o.html?"html":"text"](s),a.append(l)}if(e){var d=t('<div class="messager-text"></div>');d[o.html?"html":"text"](e),a.append(d)}var c=n.$.find(".messager-icon");if(r){var p=t.isPlainObject(r)?r.html:'<i class="icon-'+r+' icon"></i>';c.length?c.html(p):a.before('<div class="messager-icon">'+p+"<div>")}else c.remove();n.$.toggleClass("messager-has-icon",!!r),n.updateTime||o.onUpdate&&o.onUpdate.call(n,o),n.updateTime=Date.now()},r.prototype.show=function(n,o){var a=this,s=this.options;if(t.isFunction(n)){var r=o;o=n,r!==i&&(n=r)}if(a.isShow)return void a.hide(function(){a.show(n,o)});a.hiding&&(clearTimeout(a.hiding),a.hiding=null),a.update(n);var l=s.placement,d=t(s.parent),c=d.children(".messagers-holder."+l);if(c.length||(c=t("<div/>").attr("class","messagers-holder "+l).appendTo(d)),c.append(a.$),"center"===l){var p=t(e).height()-c.height();c.css("top",Math.max(-p,p/2))}return a.$.show().addClass("in"),s.time&&(a.hiding=setTimeout(function(){a.hide()},s.time)),a.isShow=!0,o&&o(),s.onShow&&s.onShow.call(a,s),a},r.prototype.hide=function(t,e){t===!0&&(e=!0,t=null);var i=this,n=i.options;if(i.$.hasClass("in")){i.$.removeClass("in");var o=function(){var o=i.$.parent();i.$.detach(),o.children().length||o.remove(),t&&t(!0),n.onHide&&n.onHide.call(i,e)};e?o():setTimeout(o,200)}else t&&t(!1),n.onHide&&n.onHide.call(i,e);i.isShow=!1},r.prototype.destroy=function(){var t=this;t.hide(function(){t.$.remove(),t.$=null},!0),delete s[t.id]};var l=function(e){if(e===i)t(".messager").each(function(){var e=t(this).data("zui.messager");e&&e.hide&&e.hide(!0)});else{var n=t("#messager-"+e).data("zui.messager");n&&n.hide&&n.hide()}},d=function(e,n){"string"==typeof n&&(n={type:n}),t.isPlainObject(e)&&(n=t.extend({},n,e),e=null),n=t.extend({},n),n.id===i&&l();var o=s[n.id]||new r(e,n);return o.show(),o},c={notification:!0,placement:"bottom-right",time:0,icon:"bell icon-2x"},p=function(e,i,n){var o=t.extend({id:t.zui.uuid()},c),a="string"==typeof e,s="string"==typeof i;return a&&s?n=t.extend(o,n,{title:e,content:i}):a&&t.isPlainObject(i)?n=t.extend(o,n,i,{title:e}):t.isPlainObject(e)?n=t.extend(o,n,i,e):a&&(n=t.extend(o,n,{title:e})),d(n)},u=function(t){return"string"==typeof t?{placement:t}:t},h={show:d,hide:l};r.all=s,r.DEFAULTS=a,r.NOTIFICATION_DEFAULTS=c,t.each({primary:0,success:"ok-sign",info:"info-sign",warning:"warning-sign",danger:"exclamation-sign",important:0,special:0},function(e,i){h[e]=function(n,o){return d(n,t.extend({type:e,icon:r.DEFAULTS.icons[e]||i||null},u(o)))}}),t.zui({Messager:r,showMessager:d,showNotification:p,messager:h})}(jQuery,window,void 0),function(t,e,i,n){"use strict";function o(t){if(t=t.toLowerCase(),t&&c.test(t)){var e;if(4===t.length){var i="#";for(e=1;e<4;e+=1)i+=t.slice(e,e+1).concat(t.slice(e,e+1));t=i}var n=[];for(e=1;e<7;e+=2)n.push(b("0x"+t.slice(e,e+2)));return{r:n[0],g:n[1],b:n[2],a:1}}throw new Error("Wrong hex string! (hex: "+t+")")}function a(e){return typeof e===f&&("transparent"===e.toLowerCase()||m[e.toLowerCase()]||c.test(t.trim(e.toLowerCase())))}function s(t){function e(t){return t=t<0?t+1:t>1?t-1:t,6*t<1?r+(s-r)*t*6:2*t<1?s:3*t<2?r+(s-r)*(2/3-t)*6:r}var i=t.h,n=t.s,o=t.l,a=t.a;i=d(i)%u/u,n=l(d(n)),o=l(d(o)),a=l(d(a));var s=o<=.5?o*(n+1):o+n-o*n,r=2*o-s,c={r:e(i+1/3)*p,g:e(i)*p,b:e(i-1/3)*p,a:a};return c}function r(t,i,n){return v(n)&&(n=0),v(i)&&(i=p),e.min(e.max(t,n),i)}function l(t,e){return r(t,e)}function d(t){return"number"==typeof t?t:parseFloat(t)}var c=/^#([0-9a-fA-f]{3}|[0-9a-fA-f]{6})$/,p=255,u=360,h=100,f="string",g="object",m={aliceblue:"#f0f8ff",antiquewhite:"#faebd7",aqua:"#00ffff",aquamarine:"#7fffd4",azure:"#f0ffff",beige:"#f5f5dc",bisque:"#ffe4c4",black:"#000000",blanchedalmond:"#ffebcd",blue:"#0000ff",blueviolet:"#8a2be2",brown:"#a52a2a",burlywood:"#deb887",cadetblue:"#5f9ea0",chartreuse:"#7fff00",chocolate:"#d2691e",coral:"#ff7f50",cornflowerblue:"#6495ed",cornsilk:"#fff8dc",crimson:"#dc143c",cyan:"#00ffff",darkblue:"#00008b",darkcyan:"#008b8b",darkgoldenrod:"#b8860b",darkgray:"#a9a9a9",darkgreen:"#006400",darkkhaki:"#bdb76b",darkmagenta:"#8b008b",darkolivegreen:"#556b2f",darkorange:"#ff8c00",darkorchid:"#9932cc",darkred:"#8b0000",darksalmon:"#e9967a",darkseagreen:"#8fbc8f",darkslateblue:"#483d8b",darkslategray:"#2f4f4f",darkturquoise:"#00ced1",darkviolet:"#9400d3",deeppink:"#ff1493",deepskyblue:"#00bfff",dimgray:"#696969",dodgerblue:"#1e90ff",firebrick:"#b22222",floralwhite:"#fffaf0",forestgreen:"#228b22",fuchsia:"#ff00ff",gainsboro:"#dcdcdc",ghostwhite:"#f8f8ff",gold:"#ffd700",goldenrod:"#daa520",gray:"#808080",green:"#008000",greenyellow:"#adff2f",honeydew:"#f0fff0",hotpink:"#ff69b4",indianred:"#cd5c5c",indigo:"#4b0082",ivory:"#fffff0",khaki:"#f0e68c",lavender:"#e6e6fa",lavenderblush:"#fff0f5",lawngreen:"#7cfc00",lemonchiffon:"#fffacd",lightblue:"#add8e6",lightcoral:"#f08080",lightcyan:"#e0ffff",lightgoldenrodyellow:"#fafad2",lightgray:"#d3d3d3",lightgreen:"#90ee90",lightpink:"#ffb6c1",lightsalmon:"#ffa07a",lightseagreen:"#20b2aa",lightskyblue:"#87cefa",lightslategray:"#778899",lightsteelblue:"#b0c4de",lightyellow:"#ffffe0",lime:"#00ff00",limegreen:"#32cd32",linen:"#faf0e6",magenta:"#ff00ff",maroon:"#800000",mediumaquamarine:"#66cdaa",mediumblue:"#0000cd",mediumorchid:"#ba55d3",mediumpurple:"#9370db",mediumseagreen:"#3cb371",mediumslateblue:"#7b68ee",mediumspringgreen:"#00fa9a",mediumturquoise:"#48d1cc",mediumvioletred:"#c71585",midnightblue:"#191970",mintcream:"#f5fffa",mistyrose:"#ffe4e1",moccasin:"#ffe4b5",navajowhite:"#ffdead",navy:"#000080",oldlace:"#fdf5e6",olive:"#808000",olivedrab:"#6b8e23",orange:"#ffa500",orangered:"#ff4500",orchid:"#da70d6",palegoldenrod:"#eee8aa",palegreen:"#98fb98",paleturquoise:"#afeeee",palevioletred:"#db7093",papayawhip:"#ffefd5",peachpuff:"#ffdab9",peru:"#cd853f",pink:"#ffc0cb",plum:"#dda0dd",powderblue:"#b0e0e6",purple:"#800080",red:"#ff0000",rosybrown:"#bc8f8f",royalblue:"#4169e1",saddlebrown:"#8b4513",salmon:"#fa8072",sandybrown:"#f4a460",seagreen:"#2e8b57",seashell:"#fff5ee",sienna:"#a0522d",silver:"#c0c0c0",skyblue:"#87ceeb",slateblue:"#6a5acd",slategray:"#708090",snow:"#fffafa",springgreen:"#00ff7f",steelblue:"#4682b4",tan:"#d2b48c",teal:"#008080",thistle:"#d8bfd8",tomato:"#ff6347",turquoise:"#40e0d0",violet:"#ee82ee",wheat:"#f5deb3",white:"#ffffff",whitesmoke:"#f5f5f5",yellow:"#ffff00",yellowgreen:"#9acd32"},v=function(t){return t===n},y=function(t){return!v(t)},b=function(t){return parseInt(t)},w=function(t){return b(l(d(t),p))},C=function(t,e,i,n){var a=this;if(a.r=a.g=a.b=0,a.a=1,y(n)&&(a.a=l(d(n),1)),y(t)&&y(e)&&y(i))a.r=w(t),a.g=w(e),a.b=w(i);else if(y(t)){var r=typeof t;if(r==f)if(t=t.toLowerCase(),"transparent"===t)a.a=0;else if(m[t])a.rgb(o(m[t]));else if(0===t.indexOf("rgb")){var c=t.substring(t.indexOf("(")+1,t.lastIndexOf(")")).split(",",4);a.rgb({r:c[0],g:c[1],b:c[2],a:c[3]})}else a.rgb(o(t));else if("number"==r&&v(e))a.r=a.g=a.b=w(t);else if(r==g&&y(t.r))a.r=w(t.r),y(t.g)&&(a.g=w(t.g)),y(t.b)&&(a.b=w(t.b)),y(t.a)&&(a.a=l(d(t.a),1));else if(r==g&&y(t.h)){var p={h:l(d(t.h),u),s:1,l:1,a:1};y(t.s)&&(p.s=l(d(t.s),1)),y(t.l)&&(p.l=l(d(t.l),1)),y(t.a)&&(p.a=l(d(t.a),1)),a.rgb(s(p))}}};C.prototype.rgb=function(t){var e=this;if(y(t)){if(typeof t==g)y(t.r)&&(e.r=w(t.r)),y(t.g)&&(e.g=w(t.g)),y(t.b)&&(e.b=w(t.b)),y(t.a)&&(e.a=l(d(t.a),1));else{var i=b(d(t));e.r=i,e.g=i,e.b=i}return e}return{r:e.r,g:e.g,b:e.b,a:e.a}},C.prototype.hue=function(t){var e=this,i=e.toHsl();return v(t)?i.h:(i.h=l(d(t),u),e.rgb(s(i)),e)},C.prototype.darken=function(t){var e=this,i=e.toHsl();return i.l-=t/h,i.l=l(i.l,1),e.rgb(s(i)),e},C.prototype.clone=function(){var t=this;return new C(t.r,t.g,t.b,t.a)},C.prototype.lighten=function(t){return this.darken(-t)},C.prototype.fade=function(t){return this.a=l(t/h,1),this},C.prototype.spin=function(t){var e=this.toHsl(),i=(e.h+t)%u;return e.h=i<0?u+i:i,this.rgb(s(e))},C.prototype.toHsl=function(){var t,i,n=this,o=n.r/p,a=n.g/p,s=n.b/p,r=n.a,l=e.max(o,a,s),d=e.min(o,a,s),c=(l+d)/2,h=l-d;if(l===d)t=i=0;else{switch(i=c>.5?h/(2-l-d):h/(l+d),l){case o:t=(a-s)/h+(a<s?6:0);break;case a:t=(s-o)/h+2;break;case s:t=(o-a)/h+4}t/=6}return{h:t*u,s:i,l:c,a:r}},C.prototype.luma=function(){var t=this.r/p,i=this.g/p,n=this.b/p;return t=t<=.03928?t/12.92:e.pow((t+.055)/1.055,2.4),i=i<=.03928?i/12.92:e.pow((i+.055)/1.055,2.4),n=n<=.03928?n/12.92:e.pow((n+.055)/1.055,2.4),.2126*t+.7152*i+.0722*n},C.prototype.saturate=function(t){var e=this.toHsl();return e.s+=t/h,e.s=l(e.s),this.rgb(s(e))},C.prototype.desaturate=function(t){return this.saturate(-t)},C.prototype.contrast=function(t,e,i){if(e=v(e)?new C(p,p,p,1):new C(e),t=v(t)?new C(0,0,0,1):new C(t),t.luma()>e.luma()){var n=e;e=t,t=n}return this.a<.5?t:(i=v(i)?.43:d(i),this.luma()<i?e:t)},C.prototype.hexStr=function(){var t=this.r.toString(16),e=this.g.toString(16),i=this.b.toString(16);return 1==t.length&&(t="0"+t),1==e.length&&(e="0"+e),1==i.length&&(i="0"+i),"#"+t+e+i},C.prototype.toCssStr=function(){var t=this;return t.a>0?t.a<1?"rgba("+t.r+","+t.g+","+t.b+","+t.a+")":t.hexStr():"transparent"},C.isColor=a,C.names=m,C.get=function(t){return new C(t)},t.zui({Color:C})}(jQuery,Math,window,void 0),function(t){"use strict";function e(e,i){if(e===!1)return e;if(!e)return i;e===!0?e={add:!0,"delete":!0,edit:!0,sort:!0}:"string"==typeof e&&(e=e.split(","));var n;return t.isArray(e)&&(n={},t.each(e,function(e,i){t.isPlainObject(i)?n[i.action]=i:n[i]=!0}),e=n),t.isPlainObject(e)&&(n={},t.each(e,function(e,i){i?n[e]=t.extend({type:e},s[e],t.isPlainObject(i)?i:null):n[e]=!1}),e=n),i?t.extend(!0,{},i,e):e}function i(e,i,n){return i=i||e.type,t(n||e.template).addClass("tree-action").attr(t.extend({"data-type":i,title:e.title||""},e.attr)).data("action",e)}var n="zui.tree",o=0,a=function(e,i){this.name=n,this.$=t(e),this.getOptions(i),this._init()},s={sort:{template:'<a class="sort-handler" href="javascript:;"><i class="icon icon-move"></i></a>'},add:{template:'<a href="javascript:;"><i class="icon icon-plus"></i></a>'},edit:{template:'<a href="javascript:;"><i class="icon icon-pencil"></i></a>'},"delete":{template:'<a href="javascript:;"><i class="icon icon-trash"></i></a>'}};a.DEFAULTS={animate:null,initialState:"normal",toggleTemplate:'<i class="list-toggle icon"></i>'},a.prototype.add=function(e,i,n,o,a){var s,r=t(e),l=this.options;if(r.is("li")?(s=r.children("ul"),s.length||(s=t("<ul/>"),r.append(s),this._initList(s,r))):s=r,s){var d=this;t.isArray(i)||(i=[i]),t.each(i,function(e,i){var n=t("<li/>").data(i).appendTo(s);void 0!==i.id&&n.attr("data-id",i.id);var o=l.itemWrapper?t(l.itemWrapper===!0?'<div class="tree-item-wrapper"/>':l.itemWrapper).appendTo(n):n;if(i.html)o.html(i.html);else if(t.isFunction(d.options.itemCreator)){var a=d.options.itemCreator(n,i);a!==!0&&a!==!1&&o.html(a)}else i.url?o.append(t("<a/>",{href:i.url}).text(i.title||i.name)):o.append(t("<span/>").text(i.title||i.name));d._initItem(n,i.idx||e,s,i),i.children&&i.children.length&&d.add(n,i.children)}),this._initList(s),n&&!s.hasClass("tree")&&d.expand(s.parent("li"),o,a)}},a.prototype.reload=function(e){var i=this;e&&(i.$.empty(),i.add(i.$,e)),i.isPreserve&&i.store.time&&i.$.find("li:not(.tree-action-item)").each(function(){var e=t(this);i[i.store[e.data("id")]?"expand":"collapse"](e,!0,!0)})},a.prototype._initList=function(n,o,a,s){var r=this;n.hasClass("tree")?(a=0,o=null):(o=(o||n.closest("li")).addClass("has-list"),o.find(".list-toggle").length||o.prepend(this.options.toggleTemplate),a=a||o.data("idx")),n.removeClass("has-active-item");var l=n.attr("data-idx",a||0).children("li:not(.tree-action-item)").each(function(e){r._initItem(t(this),e+1,n)});1!==l.length||l.find("ul").length||l.addClass("tree-single-item"),s=s||(o?o.data():null);var d=e(s?s.actions:null,this.actions);if(d){if(d.add&&d.add.templateInList!==!1){var c=n.children("li.tree-action-item");c.length?c.detach().appendTo(n):t('<li class="tree-action-item"/>').append(i(d.add,"add",d.add.templateInList)).appendTo(n)}d.sort&&n.sortable(t.extend({dragCssClass:"tree-drag-holder",trigger:".sort-handler",selector:"li:not(.tree-action-item)",finish:function(t){r.callEvent("action",{action:d.sort,$list:n,target:t.target,item:s})}},d.sort.options,t.isPlainObject(this.options.sortable)?this.options.sortable:null))}o&&(o.hasClass("open")||s&&s.open)&&o.addClass("open in")},a.prototype._initItem=function(n,o,a,s){if(void 0===o){var r=n.prev("li");o=r.length?r.data("idx")+1:1}if(a=a||n.closest("ul"),n.attr("data-idx",o).removeClass("tree-single-item"),!n.data("id")){var l=o;a.hasClass("tree")||(l=a.parent("li").data("id")+"-"+l),n.attr("data-id",l)}n.hasClass("active")&&a.parent("li").addClass("has-active-item"),s=s||n.data();var d=e(s.actions,this.actions);if(d){var c=n.find(".tree-actions");c.length||(c=t('<div class="tree-actions"/>').appendTo(this.options.itemWrapper?n.find(".tree-item-wrapper"):n),t.each(d,function(t,e){e&&c.append(i(e,t))}))}var p=n.children("ul");p.length&&this._initList(p,n,o,s)},a.prototype._init=function(){var i=this.options,a=this;this.actions=e(i.actions),this.$.addClass("tree"),i.animate&&this.$.addClass("tree-animate"),this._initList(this.$);var s=i.initialState,r=t.zui&&t.zui.store&&t.zui.store.enable;r&&(this.selector=n+"::"+(i.name||"")+"#"+(this.$.attr("id")||o++),this.store=t.zui.store[i.name?"get":"pageGet"](this.selector,{})),"preserve"===s&&(r?this.isPreserve=!0:this.options.initialState=s="normal"),this.reload(i.data),r&&(this.isPreserve=!0),"expand"===s?this.expand():"collapse"===s?this.collapse():"active"===s&&this.expandSelect(".active"),this.$.on("click",'.list-toggle,a[href="#"],.tree-toggle',function(e){var i=t(this),n=i.parent("li");a.callEvent("hit",{target:n,item:n.data()}),a.toggle(n),i.is("a")&&e.preventDefault()}).on("click",".tree-action",function(){var e=t(this),i=e.data();if(i.action&&(i=i.action),"sort"!==i.type){var n=e.closest("li:not(.tree-action-item)");a.callEvent("action",{action:i,target:this,$item:n,item:n.data()})}})},a.prototype.preserve=function(e,i,n){if(this.isPreserve)if(e)i=i||e.data("id"),n=void 0===n&&e.hasClass("open"),n?this.store[i]=n:delete this.store[i],this.store.time=(new Date).getTime(),t.zui.store[this.options.name?"set":"pageSet"](this.selector,this.store);else{var o=this;this.store={},this.$.find("li").each(function(){o.preserve(t(this))})}},a.prototype.expandSelect=function(t){this.show(t,!0)},a.prototype.expand=function(t,e,i){t?(t.addClass("open"),!e&&this.options.animate?setTimeout(function(){t.addClass("in")},10):t.addClass("in")):t=this.$.find("li.has-list").addClass("open in"),i||this.preserve(t),this.callEvent("expand",t,this)},a.prototype.show=function(e,i,n){var o=this;e instanceof t||(e=o.$.find("li").filter(e)),e.each(function(){var e=t(this);if(o.expand(e,i,n),e)for(var a=e.parent("ul");a&&a.length&&!a.hasClass("tree");){var s=a.parent("li");s.length?(o.expand(s,i,n),a=s.parent("ul")):a=!1}})},a.prototype.collapse=function(t,e,i){t?!e&&this.options.animate?(t.removeClass("in"),setTimeout(function(){t.removeClass("open")},300)):t.removeClass("open in"):t=this.$.find("li.has-list").removeClass("open in"),i||this.preserve(t),this.callEvent("collapse",t,this)},a.prototype.toggle=function(t){var e=t&&t.hasClass("open")||t===!1||void 0===t&&this.$.find("li.has-list.open").length;this[e?"collapse":"expand"](t)},a.prototype.getOptions=function(e){this.options=t.extend({},a.DEFAULTS,this.$.data(),e),null===this.options.animate&&this.$.hasClass("tree-animate")&&(this.options.animate=!0)},a.prototype.toData=function(e,i){t.isFunction(e)&&(i=e,e=null),e=e||this.$;var n=this;return e.children("li:not(.tree-action-item)").map(function(){var e=t(this),o=e.data();delete o["zui.droppable"];var a=e.children("ul");return a.length&&(o.children=n.toData(a)),t.isFunction(i)?i(o,e):o}).get()},a.prototype.callEvent=function(e,i){var n;return t.isFunction(this.options[e])&&(n=this.options[e](i,this)),this.$.trigger(t.Event(e+"."+this.name,i)),n},t.fn.tree=function(e,i){return this.each(function(){var o=t(this),s=o.data(n),r="object"==typeof e&&e;s||o.data(n,s=new a(this,r)),"string"==typeof e&&s[e](i)})},t.fn.tree.Constructor=a,t(function(){t('[data-ride="tree"]').tree()})}(jQuery),
/*!
 * Chart.js 1.0.2
 * Copyright 2015 Nick Downie
 * Released under the MIT license
 * http://chartjs.org/
 */

function(V) {
	var ak = V && V.zui ? V.zui : this,
		ag = (ak.Chart, function(b) {
			this.canvas = b.canvas, this.ctx = b;
			var c = function(f, g) {
					return f["offset" + g] ? f["offset" + g] : document.defaultView.getComputedStyle(f).getPropertyValue(g)
				},
				a = this.width = c(b.canvas, "Width"),
				d = this.height = c(b.canvas, "Height");
			b.canvas.width = a, b.canvas.height = d;
			var a = this.width = b.canvas.width,
				d = this.height = b.canvas.height;
			return this.aspectRatio = this.width / this.height, ac.retinaScale(this), this
		});
	ag.defaults = {
		global: {
			animation: !0,
			animationSteps: 60,
			animationEasing: "easeOutQuart",
			showScale: !0,
			scaleOverride: !1,
			scaleSteps: null,
			scaleStepWidth: null,
			scaleStartValue: null,
			scaleLineColor: "rgba(0,0,0,.1)",
			scaleLineWidth: 1,
			scaleShowLabels: !0,
			scaleLabel: "<%=value%>",
			scaleIntegersOnly: !0,
			scaleBeginAtZero: !1,
			scaleFontFamily: "'Helvetica Neue', 'Helvetica', 'Arial', sans-serif",
			scaleFontSize: 12,
			scaleFontStyle: "normal",
			scaleFontColor: "#666",
			responsive: !1,
			maintainAspectRatio: !0,
			showTooltips: !0,
			customTooltips: !1,
			tooltipEvents: ["mousemove", "touchstart", "touchmove", "mouseout"],
			tooltipFillColor: "rgba(0,0,0,0.8)",
			tooltipFontFamily: "'Helvetica Neue', 'Helvetica', 'Arial', sans-serif",
			tooltipFontSize: 14,
			tooltipFontStyle: "normal",
			tooltipFontColor: "#fff",
			tooltipTitleFontFamily: "'Helvetica Neue', 'Helvetica', 'Arial', sans-serif",
			tooltipTitleFontSize: 14,
			tooltipTitleFontStyle: "bold",
			tooltipTitleFontColor: "#fff",
			tooltipYPadding: 6,
			tooltipXPadding: 6,
			tooltipCaretSize: 8,
			tooltipCornerRadius: 6,
			tooltipXOffset: 10,
			tooltipTemplate: "<%if (label){%><%=label%>: <%}%><%= value %>",
			multiTooltipTemplate: "<%if (datasetLabel){%><%=datasetLabel%>: <%}%><%= value %>",
			multiTooltipTitleTemplate: "<%= label %>",
			multiTooltipKeyBackground: "#fff",
			onAnimationProgress: function() {},
			onAnimationComplete: function() {}
		}
	}, ag.types = {};
	var ac = ag.helpers = {},
		ab = ac.each = function(b, d, a) {
			var g = Array.prototype.slice.call(arguments, 3);
			if (b) {
				if (b.length === +b.length) {
					var f;
					for (f = 0; f < b.length; f++) {
						d.apply(a, [b[f], f].concat(g))
					}
				} else {
					for (var c in b) {
						d.apply(a, [b[c], c].concat(g))
					}
				}
			}
		},
		X = ac.clone = function(a) {
			var b = {};
			return ab(a, function(c, d) {
				a.hasOwnProperty(d) && (b[d] = c)
			}), b
		},
		ao = ac.extend = function(a) {
			return ab(Array.prototype.slice.call(arguments, 1), function(b) {
				ab(b, function(c, d) {
					b.hasOwnProperty(d) && (a[d] = c)
				})
			}), a
		},
		Z = ac.merge = function(b, c) {
			var a = Array.prototype.slice.call(arguments, 0);
			return a.unshift({}), ao.apply(null, a)
		},
		ae = ac.indexOf = function(b, c) {
			if (Array.prototype.indexOf) {
				return b.indexOf(c)
			}
			for (var a = 0; a < b.length; a++) {
				if (b[a] === c) {
					return a
				}
			}
			return -1
		},
		ah = (ac.where = function(b, c) {
			var a = [];
			return ac.each(b, function(d) {
				c(d) && a.push(d)
			}), a
		}, ac.findNextWhere = function(b, c, a) {
			a || (a = -1);
			for (var f = a + 1; f < b.length; f++) {
				var d = b[f];
				if (c(d)) {
					return d
				}
			}
		}, ac.findPreviousWhere = function(b, c, a) {
			a || (a = b.length);
			for (var f = a - 1; f >= 0; f--) {
				var d = b[f];
				if (c(d)) {
					return d
				}
			}
		}, ac.inherits = function(b) {
			var c = this,
				a = b && b.hasOwnProperty("constructor") ? b.constructor : function() {
					return c.apply(this, arguments)
				},
				d = function() {
					this.constructor = a
				};
			return d.prototype = c.prototype, a.prototype = new d, a.extend = ah, b && ao(a.prototype, b), a.__super__ = c.prototype, a
		}),
		am = ac.noop = function() {},
		al = ac.uid = function() {
			var a = 0;
			return function() {
				return "chart-" + a++
			}
		}(),
		U = ac.warn = function(a) {
			window.console && "function" == typeof window.console.warn && console.warn(a)
		},
		aa = ac.amd = "function" == typeof define && define.amd,
		aj = ac.isNumber = function(a) {
			return !isNaN(parseFloat(a)) && isFinite(a)
		},
		ai = ac.max = function(a) {
			return Math.max.apply(Math, a)
		},
		ad = ac.min = function(a) {
			return Math.min.apply(Math, a)
		},
		Q = (ac.cap = function(b, c, a) {
			if (aj(c)) {
				if (b > c) {
					return c
				}
			} else {
				if (aj(a) && b < a) {
					return a
				}
			}
			return b
		}, ac.getDecimalPlaces = function(a) {
			return a % 1 !== 0 && aj(a) ? a.toString().split(".")[1].length : 0
		}),
		J = ac.radians = function(a) {
			return a * (Math.PI / 180)
		},
		an = (ac.getAngleFromPoint = function(b, d) {
			var a = d.x - b.x,
				g = d.y - b.y,
				f = Math.sqrt(a * a + g * g),
				c = 2 * Math.PI + Math.atan2(g, a);
			return a < 0 && g < 0 && (c += 2 * Math.PI), {
				angle: c,
				distance: f
			}
		}, ac.aliasPixel = function(a) {
			return a % 2 === 0 ? 0 : 0.5
		}),
		O = (ac.splineCurve = function(d, h, c, l) {
			var k = Math.sqrt(Math.pow(h.x - d.x, 2) + Math.pow(h.y - d.y, 2)),
				f = Math.sqrt(Math.pow(c.x - h.x, 2) + Math.pow(c.y - h.y, 2)),
				b = l * k / (k + f),
				g = l * f / (k + f);
			return {
				inner: {
					x: h.x - b * (c.x - d.x),
					y: h.y - b * (c.y - d.y)
				},
				outer: {
					x: h.x + g * (c.x - d.x),
					y: h.y + g * (c.y - d.y)
				}
			}
		}, ac.calculateOrderOfMagnitude = function(a) {
			return Math.floor(Math.log(a) / Math.LN10)
		}),
		N = (ac.calculateScaleRange = function(T, D, x, m, k) {
			var aq = 2,
				M = Math.floor(D / (1.5 * x)),
				b = aq >= M,
				w = ai(T),
				z = ad(T);
			w === z && (w += 0.5, z >= 0.5 && !m ? z -= 0.5 : w += 0.5);
			for (var I = Math.abs(w - z), F = O(I), S = Math.ceil(w / (1 * Math.pow(10, F))) * Math.pow(10, F), g = m ? 0 : Math.floor(z / (1 * Math.pow(10, F))) * Math.pow(10, F), C = S - g, P = Math.pow(10, F), L = Math.round(C / P);
			(L > M || 2 * L < M) && !b;) {
				if (L > M) {
					P *= 2, L = Math.round(C / P), L % 1 !== 0 && (b = !0)
				} else {
					if (k && F >= 0) {
						if (P / 2 % 1 !== 0) {
							break
						}
						P /= 2, L = Math.round(C / P)
					} else {
						P /= 2, L = Math.round(C / P)
					}
				}
			}
			return b && (L = aq, P = C / L), {
				steps: L,
				stepValue: P,
				min: g,
				max: g + L * P
			}
		}, ac.template = function(b, c) {
			function a(g, h) {
				var f = /\W/.test(g) ? new Function("obj", "var p=[],print=function(){p.push.apply(p,arguments);};with(obj){p.push('" + g.replace(/[\r\t\n]/g, " ").split("<%").join("\t").replace(/((^|%>)[^\t]*)'/g, "$1\r").replace(/\t=(.*?)%>/g, "',$1,'").split("\t").join("');").split("%>").join("p.push('").split("\r").join("\\'") + "');}return p.join('');") : d[g] = d[g];
				return h ? f(h) : f
			}
			if (b instanceof Function) {
				return b(c)
			}
			var d = {};
			return a(b, c)
		}),
		Y = (ac.generateLabels = function(b, d, a, f) {
			var c = new Array(d);
			return labelTemplateString && ab(c, function(g, h) {
				c[h] = N(b, {
					value: a + f * (h + 1)
				})
			}), c
		}, ac.easingEffects = {
			linear: function(a) {
				return a
			},
			easeInQuad: function(a) {
				return a * a
			},
			easeOutQuad: function(a) {
				return -1 * a * (a - 2)
			},
			easeInOutQuad: function(a) {
				return (a /= 0.5) < 1 ? 0.5 * a * a : -0.5 * (--a * (a - 2) - 1)
			},
			easeInCubic: function(a) {
				return a * a * a
			},
			easeOutCubic: function(a) {
				return 1 * ((a = a / 1 - 1) * a * a + 1)
			},
			easeInOutCubic: function(a) {
				return (a /= 0.5) < 1 ? 0.5 * a * a * a : 0.5 * ((a -= 2) * a * a + 2)
			},
			easeInQuart: function(a) {
				return a * a * a * a
			},
			easeOutQuart: function(a) {
				return -1 * ((a = a / 1 - 1) * a * a * a - 1)
			},
			easeInOutQuart: function(a) {
				return (a /= 0.5) < 1 ? 0.5 * a * a * a * a : -0.5 * ((a -= 2) * a * a * a - 2)
			},
			easeInQuint: function(a) {
				return 1 * (a /= 1) * a * a * a * a
			},
			easeOutQuint: function(a) {
				return 1 * ((a = a / 1 - 1) * a * a * a * a + 1)
			},
			easeInOutQuint: function(a) {
				return (a /= 0.5) < 1 ? 0.5 * a * a * a * a * a : 0.5 * ((a -= 2) * a * a * a * a + 2)
			},
			easeInSine: function(a) {
				return -1 * Math.cos(a / 1 * (Math.PI / 2)) + 1
			},
			easeOutSine: function(a) {
				return 1 * Math.sin(a / 1 * (Math.PI / 2))
			},
			easeInOutSine: function(a) {
				return -0.5 * (Math.cos(Math.PI * a / 1) - 1)
			},
			easeInExpo: function(a) {
				return 0 === a ? 1 : 1 * Math.pow(2, 10 * (a / 1 - 1))
			},
			easeOutExpo: function(a) {
				return 1 === a ? 1 : 1 * (-Math.pow(2, -10 * a / 1) + 1)
			},
			easeInOutExpo: function(a) {
				return 0 === a ? 0 : 1 === a ? 1 : (a /= 0.5) < 1 ? 0.5 * Math.pow(2, 10 * (a - 1)) : 0.5 * (-Math.pow(2, -10 * --a) + 2)
			},
			easeInCirc: function(a) {
				return a >= 1 ? a : -1 * (Math.sqrt(1 - (a /= 1) * a) - 1)
			},
			easeOutCirc: function(a) {
				return 1 * Math.sqrt(1 - (a = a / 1 - 1) * a)
			},
			easeInOutCirc: function(a) {
				return (a /= 0.5) < 1 ? -0.5 * (Math.sqrt(1 - a * a) - 1) : 0.5 * (Math.sqrt(1 - (a -= 2) * a) + 1)
			},
			easeInElastic: function(b) {
				var c = 1.70158,
					a = 0,
					d = 1;
				return 0 === b ? 0 : 1 == (b /= 1) ? 1 : (a || (a = 0.3), d < Math.abs(1) ? (d = 1, c = a / 4) : c = a / (2 * Math.PI) * Math.asin(1 / d), -(d * Math.pow(2, 10 * (b -= 1)) * Math.sin((1 * b - c) * (2 * Math.PI) / a)))
			},
			easeOutElastic: function(b) {
				var c = 1.70158,
					a = 0,
					d = 1;
				return 0 === b ? 0 : 1 == (b /= 1) ? 1 : (a || (a = 0.3), d < Math.abs(1) ? (d = 1, c = a / 4) : c = a / (2 * Math.PI) * Math.asin(1 / d), d * Math.pow(2, -10 * b) * Math.sin((1 * b - c) * (2 * Math.PI) / a) + 1)
			},
			easeInOutElastic: function(b) {
				var c = 1.70158,
					a = 0,
					d = 1;
				return 0 === b ? 0 : 2 == (b /= 0.5) ? 1 : (a || (a = 1 * (0.3 * 1.5)), d < Math.abs(1) ? (d = 1, c = a / 4) : c = a / (2 * Math.PI) * Math.asin(1 / d), b < 1 ? -0.5 * (d * Math.pow(2, 10 * (b -= 1)) * Math.sin((1 * b - c) * (2 * Math.PI) / a)) : d * Math.pow(2, -10 * (b -= 1)) * Math.sin((1 * b - c) * (2 * Math.PI) / a) * 0.5 + 1)
			},
			easeInBack: function(a) {
				var b = 1.70158;
				return 1 * (a /= 1) * a * ((b + 1) * a - b)
			},
			easeOutBack: function(a) {
				var b = 1.70158;
				return 1 * ((a = a / 1 - 1) * a * ((b + 1) * a + b) + 1)
			},
			easeInOutBack: function(a) {
				var b = 1.70158;
				return (a /= 0.5) < 1 ? 0.5 * (a * a * (((b *= 1.525) + 1) * a - b)) : 0.5 * ((a -= 2) * a * (((b *= 1.525) + 1) * a + b) + 2)
			},
			easeInBounce: function(a) {
				return 1 - Y.easeOutBounce(1 - a)
			},
			easeOutBounce: function(a) {
				return (a /= 1) < 1 / 2.75 ? 1 * (7.5625 * a * a) : a < 2 / 2.75 ? 1 * (7.5625 * (a -= 1.5 / 2.75) * a + 0.75) : a < 2.5 / 2.75 ? 1 * (7.5625 * (a -= 2.25 / 2.75) * a + 0.9375) : 1 * (7.5625 * (a -= 2.625 / 2.75) * a + 0.984375)
			},
			easeInOutBounce: function(a) {
				return a < 0.5 ? 0.5 * Y.easeInBounce(2 * a) : 0.5 * Y.easeOutBounce(2 * a - 1) + 0.5
			}
		}),
		ap = ac.requestAnimFrame = function() {
			return window.requestAnimationFrame || window.webkitRequestAnimationFrame || window.mozRequestAnimationFrame || window.oRequestAnimationFrame || window.msRequestAnimationFrame ||
			function(a) {
				return window.setTimeout(a, 1000 / 60)
			}
		}(),
		af = ac.cancelAnimFrame = function() {
			return window.cancelAnimationFrame || window.webkitCancelAnimationFrame || window.mozCancelAnimationFrame || window.oCancelAnimationFrame || window.msCancelAnimationFrame ||
			function(a) {
				return window.clearTimeout(a, 1000 / 60)
			}
		}(),
		j = (ac.animationLoop = function(m, h, g, d, c, p) {
			var k = 0,
				b = Y[g] || Y.linear,
				f = function() {
					k++;
					var a = k / h,
						e = b(a);
					m.call(p, e, a, k), d.call(p, e, a), k < h ? p.animationFrame = ap(f) : c.apply(p)
				};
			ap(f)
		}, ac.getRelativePosition = function(b) {
			var d, a, g = b.originalEvent || b,
				f = b.currentTarget || b.srcElement,
				c = f.getBoundingClientRect();
			return g.touches ? (d = g.touches[0].clientX - c.left, a = g.touches[0].clientY - c.top) : (d = g.clientX - c.left, a = g.clientY - c.top), {
				x: d,
				y: a
			}
		}, ac.addEvent = function(b, c, a) {
			b.addEventListener ? b.addEventListener(c, a) : b.attachEvent ? b.attachEvent("on" + c, a) : b["on" + c] = a
		}),
		A = ac.removeEvent = function(b, c, a) {
			b.removeEventListener ? b.removeEventListener(c, a, !1) : b.detachEvent ? b.detachEvent("on" + c, a) : b["on" + c] = am
		},
		W = (ac.bindEvents = function(b, c, a) {
			b.events || (b.events = {}), ab(c, function(d) {
				b.events[d] = function() {
					a.apply(b, arguments)
				}, j(b.chart.canvas, d, b.events[d])
			})
		}, ac.unbindEvents = function(a, b) {
			ab(b, function(d, c) {
				A(a.chart.canvas, c, d)
			})
		}),
		E = ac.getMaximumWidth = function(a) {
			var b = a.parentNode;
			return b.clientWidth
		},
		B = ac.getMaximumHeight = function(a) {
			var b = a.parentNode;
			return b.clientHeight
		},
		R = (ac.getMaximumSize = ac.getMaximumWidth, ac.retinaScale = function(b) {
			var c = b.ctx,
				a = b.canvas.width,
				d = b.canvas.height;
			window.devicePixelRatio && (c.canvas.style.width = a + "px", c.canvas.style.height = d + "px", c.canvas.height = d * window.devicePixelRatio, c.canvas.width = a * window.devicePixelRatio, c.scale(window.devicePixelRatio, window.devicePixelRatio))
		}),
		G = ac.clear = function(a) {
			a.ctx.clearRect(0, 0, a.width, a.height)
		},
		H = ac.fontString = function(b, c, a) {
			return c + " " + b + "px " + a
		},
		K = ac.longestText = function(b, c, a) {
			b.font = c;
			var d = 0;
			return ab(a, function(g) {
				var f = b.measureText(g).width;
				d = f > d ? f : d
			}), d
		},
		q = ac.drawRoundedRectangle = function(b, d, a, g, f, c) {
			b.beginPath(), b.moveTo(d + c, a), b.lineTo(d + g - c, a), b.quadraticCurveTo(d + g, a, d + g, a + c), b.lineTo(d + g, a + f - c), b.quadraticCurveTo(d + g, a + f, d + g - c, a + f), b.lineTo(d + c, a + f), b.quadraticCurveTo(d, a + f, d, a + f - c), b.lineTo(d, a + c), b.quadraticCurveTo(d, a, d + c, a), b.closePath()
		};
	ag.instances = {}, ag.Type = function(a, b, c) {
		this.options = b, this.chart = c, this.id = al(), ag.instances[this.id] = this, b.responsive && this.resize(), this.initialize.call(this, a)
	}, ao(ag.Type.prototype, {
		initialize: function() {
			return this
		},
		clear: function() {
			return G(this.chart), this
		},
		stop: function() {
			return af(this.animationFrame), this
		},
		resize: function(b) {
			this.stop();
			var c = this.chart.canvas,
				a = E(this.chart.canvas),
				d = this.options.maintainAspectRatio ? a / this.chart.aspectRatio : B(this.chart.canvas);
			return c.width = this.chart.width = a, c.height = this.chart.height = d, R(this.chart), "function" == typeof b && b.apply(this, Array.prototype.slice.call(arguments, 1)), this
		},
		reflow: am,
		render: function(a) {
			return a && this.reflow(), this.options.animation && !a ? ac.animationLoop(this.draw, this.options.animationSteps, this.options.animationEasing, this.options.onAnimationProgress, this.options.onAnimationComplete, this) : (this.draw(), this.options.onAnimationComplete.call(this)), this
		},
		generateLegend: function() {
			return N(this.options.legendTemplate, this)
		},
		destroy: function() {
			this.clear(), W(this, this.events);
			var a = this.chart.canvas;
			a.width = this.chart.width, a.height = this.chart.height, a.style.removeProperty ? (a.style.removeProperty("width"), a.style.removeProperty("height")) : (a.style.removeAttribute("width"), a.style.removeAttribute("height")), delete ag.instances[this.id]
		},
		showTooltip: function(o, g) {
			"undefined" == typeof this.activeElements && (this.activeElements = []);
			var p = function(a) {
					var c = !1;
					return a.length !== this.activeElements.length ? c = !0 : (ab(a, function(e, d) {
						e !== this.activeElements[d] && (c = !0)
					}, this), c)
				}.call(this, o);
			if (p || g) {
				if (this.activeElements = o, this.draw(), this.options.customTooltips && this.options.customTooltips(!1), o.length > 0) {
					if (this.datasets && this.datasets.length > 1) {
						for (var m, b, f = this.datasets.length - 1; f >= 0 && (m = this.datasets[f].points || this.datasets[f].bars || this.datasets[f].segments, b = ae(m, o[0]), b === -1); f--) {}
						var l = [],
							k = [],
							n = function(z) {
								var w, r, c, C, x, d = [],
									v = [],
									y = [];
								return ac.each(this.datasets, function(a) {
									a.showTooltips !== !1 && (w = a.points || a.bars || a.segments, w[b] && w[b].hasValue() && d.push(w[b]))
								}), ac.each(d, function(a) {
									v.push(a.x), y.push(a.y), l.push(ac.template(this.options.multiTooltipTemplate, a)), k.push({
										fill: a._saved.fillColor || a.fillColor,
										stroke: a._saved.strokeColor || a.strokeColor
									})
								}, this), x = ad(y), c = ai(y), C = ad(v), r = ai(v), {
									x: C > this.chart.width / 2 ? C : r,
									y: (x + c) / 2
								}
							}.call(this, b);
						new ag.MultiTooltip({
							x: n.x,
							y: n.y,
							xPadding: this.options.tooltipXPadding,
							yPadding: this.options.tooltipYPadding,
							xOffset: this.options.tooltipXOffset,
							fillColor: this.options.tooltipFillColor,
							textColor: this.options.tooltipFontColor,
							fontFamily: this.options.tooltipFontFamily,
							fontStyle: this.options.tooltipFontStyle,
							fontSize: this.options.tooltipFontSize,
							titleTextColor: this.options.tooltipTitleFontColor,
							titleFontFamily: this.options.tooltipTitleFontFamily,
							titleFontStyle: this.options.tooltipTitleFontStyle,
							titleFontSize: this.options.tooltipTitleFontSize,
							cornerRadius: this.options.tooltipCornerRadius,
							labels: l,
							legendColors: k,
							legendColorBackground: this.options.multiTooltipKeyBackground,
							title: N(this.options.multiTooltipTitleTemplate, o[0]),
							chart: this.chart,
							ctx: this.chart.ctx,
							custom: this.options.customTooltips
						}).draw()
					} else {
						ab(o, function(a) {
							var c = a.tooltipPosition();
							new ag.Tooltip({
								x: Math.round(c.x),
								y: Math.round(c.y),
								xPadding: this.options.tooltipXPadding,
								yPadding: this.options.tooltipYPadding,
								fillColor: this.options.tooltipFillColor,
								textColor: this.options.tooltipFontColor,
								fontFamily: this.options.tooltipFontFamily,
								fontStyle: this.options.tooltipFontStyle,
								fontSize: this.options.tooltipFontSize,
								caretHeight: this.options.tooltipCaretSize,
								cornerRadius: this.options.tooltipCornerRadius,
								text: N(this.options.tooltipTemplate, a),
								chart: this.chart,
								custom: this.options.customTooltips
							}).draw()
						}, this)
					}
				}
				return this
			}
		},
		toBase64Image: function() {
			return this.chart.canvas.toDataURL.apply(this.chart.canvas, arguments)
		}
	}), ag.Type.extend = function(b) {
		var c = this,
			f = function() {
				return c.apply(this, arguments)
			};
		if (f.prototype = X(c.prototype), ao(f.prototype, b), f.extend = ag.Type.extend, b.name || c.prototype.name) {
			var d = b.name || c.prototype.name,
				a = ag.defaults[c.prototype.name] ? X(ag.defaults[c.prototype.name]) : {};
			ag.defaults[d] = ao(a, b.defaults), ag.types[d] = f, ag.prototype[d] = function(g, k) {
				var h = Z(ag.defaults.global, ag.defaults[d], k || {});
				return new f(g, h, this)
			}
		} else {
			U("Name not provided for this chart, so it hasn't been registered")
		}
		return c
	}, ag.Element = function(a) {
		ao(this, a), this.initialize.apply(this, arguments), this.save()
	}, ao(ag.Element.prototype, {
		initialize: function() {},
		restore: function(a) {
			return a ? ab(a, function(b) {
				this[b] = this._saved[b]
			}, this) : ao(this, this._saved), this
		},
		save: function() {
			return this._saved = X(this), delete this._saved._saved, this
		},
		update: function(a) {
			return ab(a, function(b, c) {
				this._saved[c] = this[c], this[c] = b
			}, this), this
		},
		transition: function(a, b) {
			return ab(a, function(d, c) {
				this[c] = (d - this._saved[c]) * b + this._saved[c]
			}, this), this
		},
		tooltipPosition: function() {
			return {
				x: this.x,
				y: this.y
			}
		},
		hasValue: function() {
			return aj(this.value)
		}
	}), ag.Element.extend = ah, ag.Point = ag.Element.extend({
		display: !0,
		inRange: function(b, c) {
			var a = this.hitDetectionRadius + this.radius;
			return Math.pow(b - this.x, 2) + Math.pow(c - this.y, 2) < Math.pow(a, 2)
		},
		draw: function() {
			if (this.display) {
				var a = this.ctx;
				a.beginPath(), a.arc(this.x, this.y, this.radius, 0, 2 * Math.PI), a.closePath(), a.strokeStyle = this.strokeColor, a.lineWidth = this.strokeWidth, a.fillStyle = this.fillColor, a.fill(), a.stroke()
			}
		}
	}), ag.Arc = ag.Element.extend({
		inRange: function(b, d) {
			var a = ac.getAngleFromPoint(this, {
				x: b,
				y: d
			}),
				f = a.angle >= this.startAngle && a.angle <= this.endAngle,
				c = a.distance >= this.innerRadius && a.distance <= this.outerRadius;
			return f && c
		},
		tooltipPosition: function() {
			var a = this.startAngle + (this.endAngle - this.startAngle) / 2,
				b = (this.outerRadius - this.innerRadius) / 2 + this.innerRadius;
			return {
				x: this.x + Math.cos(a) * b,
				y: this.y + Math.sin(a) * b
			}
		},
		draw: function(b) {
			var c = this.ctx;
			if (c.beginPath(), c.arc(this.x, this.y, this.outerRadius, this.startAngle, this.endAngle), c.arc(this.x, this.y, this.innerRadius, this.endAngle, this.startAngle, !0), c.closePath(), c.strokeStyle = this.strokeColor, c.lineWidth = this.strokeWidth, c.fillStyle = this.fillColor, c.fill(), c.lineJoin = "bevel", this.showStroke && c.stroke(), this.circleBeginEnd) {
				var a = (this.outerRadius + this.innerRadius) / 2,
					d = (this.outerRadius - this.innerRadius) / 2;
				c.beginPath(), c.arc(this.x + Math.cos(this.startAngle) * a, this.y + Math.sin(this.startAngle) * a, d, 0, 2 * Math.PI), c.closePath(), c.fill(), c.beginPath(), c.arc(this.x + Math.cos(this.endAngle) * a, this.y + Math.sin(this.endAngle) * a, d, 0, 2 * Math.PI), c.closePath(), c.fill()
			}
		}
	}), ag.Rectangle = ag.Element.extend({
		draw: function() {
			var b = this.ctx,
				d = this.width / 2,
				a = this.x - d,
				g = this.x + d,
				f = this.base - (this.base - this.y),
				c = this.strokeWidth / 2;
			this.showStroke && (a += c, g -= c, f += c), b.beginPath(), b.fillStyle = this.fillColor, b.strokeStyle = this.strokeColor, b.lineWidth = this.strokeWidth, b.moveTo(a, this.base), b.lineTo(a, f), b.lineTo(g, f), b.lineTo(g, this.base), b.fill(), this.showStroke && b.stroke()
		},
		height: function() {
			return this.base - this.y
		},
		inRange: function(a, b) {
			return a >= this.x - this.width / 2 && a <= this.x + this.width / 2 && b >= this.y && b <= this.base
		}
	}), ag.Tooltip = ag.Element.extend({
		draw: function() {
			var d = this.chart.ctx;
			d.font = H(this.fontSize, this.fontStyle, this.fontFamily), this.xAlign = "center", this.yAlign = "above";
			var g = this.caretPadding = 2,
				c = d.measureText(this.text).width + 2 * this.xPadding,
				k = this.fontSize + 2 * this.yPadding,
				h = k + this.caretHeight + g;
			this.x + c / 2 > this.chart.width ? this.xAlign = "left" : this.x - c / 2 < 0 && (this.xAlign = "right"), this.y - h < 0 && (this.yAlign = "below");
			var f = this.x - c / 2,
				b = this.y - h;
			if (d.fillStyle = this.fillColor, this.custom) {
				this.custom(this)
			} else {
				switch (this.yAlign) {
				case "above":
					d.beginPath(), d.moveTo(this.x, this.y - g), d.lineTo(this.x + this.caretHeight, this.y - (g + this.caretHeight)), d.lineTo(this.x - this.caretHeight, this.y - (g + this.caretHeight)), d.closePath(), d.fill();
					break;
				case "below":
					b = this.y + g + this.caretHeight, d.beginPath(), d.moveTo(this.x, this.y + g), d.lineTo(this.x + this.caretHeight, this.y + g + this.caretHeight), d.lineTo(this.x - this.caretHeight, this.y + g + this.caretHeight), d.closePath(), d.fill()
				}
				switch (this.xAlign) {
				case "left":
					f = this.x - c + (this.cornerRadius + this.caretHeight);
					break;
				case "right":
					f = this.x - (this.cornerRadius + this.caretHeight)
				}
				q(d, f, b, c, k, this.cornerRadius), d.fill(), d.fillStyle = this.textColor, d.textAlign = "center", d.textBaseline = "middle", d.fillText(this.text, f + c / 2, b + k / 2)
			}
		}
	}), ag.MultiTooltip = ag.Element.extend({
		initialize: function() {
			this.font = H(this.fontSize, this.fontStyle, this.fontFamily), this.titleFont = H(this.titleFontSize, this.titleFontStyle, this.titleFontFamily), this.height = this.labels.length * this.fontSize + (this.labels.length - 1) * (this.fontSize / 2) + 2 * this.yPadding + 1.5 * this.titleFontSize, this.ctx.font = this.titleFont;
			var b = this.ctx.measureText(this.title).width,
				c = K(this.ctx, this.font, this.labels) + this.fontSize + 3,
				a = ai([c, b]);
			this.width = a + 2 * this.xPadding;
			var d = this.height / 2;
			this.y - d < 0 ? this.y = d : this.y + d > this.chart.height && (this.y = this.chart.height - d), this.x > this.chart.width / 2 ? this.x -= this.xOffset + this.width : this.x += this.xOffset
		},
		getLineHeight: function(b) {
			var c = this.y - this.height / 2 + this.yPadding,
				a = b - 1;
			return 0 === b ? c + this.titleFontSize / 2 : c + (1.5 * this.fontSize * a + this.fontSize / 2) + 1.5 * this.titleFontSize
		},
		draw: function() {
			if (this.custom) {
				this.custom(this)
			} else {
				q(this.ctx, this.x, this.y - this.height / 2, this.width, this.height, this.cornerRadius);
				var a = this.ctx;
				a.fillStyle = this.fillColor, a.fill(), a.closePath(), a.textAlign = "left", a.textBaseline = "middle", a.fillStyle = this.titleTextColor, a.font = this.titleFont, a.fillText(this.title, this.x + this.xPadding, this.getLineHeight(0)), a.font = this.font, ac.each(this.labels, function(c, b) {
					a.fillStyle = this.textColor, a.fillText(c, this.x + this.xPadding + this.fontSize + 3, this.getLineHeight(b + 1)), a.fillStyle = this.legendColorBackground, a.fillRect(this.x + this.xPadding, this.getLineHeight(b + 1) - this.fontSize / 2, this.fontSize, this.fontSize), a.fillStyle = this.legendColors[b].fill, a.fillRect(this.x + this.xPadding, this.getLineHeight(b + 1) - this.fontSize / 2, this.fontSize, this.fontSize)
				}, this)
			}
		}
	}), ag.Scale = ag.Element.extend({
		initialize: function() {
			this.fit()
		},
		buildYLabels: function() {
			this.yLabels = [];
			for (var a = Q(this.stepValue), b = 0; b <= this.steps; b++) {
				this.yLabels.push(N(this.templateString, {
					value: (this.min + b * this.stepValue).toFixed(a)
				}))
			}
			this.yLabelWidth = this.display && this.showLabels ? K(this.ctx, this.font, this.yLabels) : 0
		},
		addXLabel: function(a) {
			this.xLabels.push(a), this.valuesCount++, this.fit()
		},
		removeXLabel: function() {
			this.xLabels.shift(), this.valuesCount--, this.fit()
		},
		fit: function() {
			this.startPoint = this.display ? this.fontSize : 0, this.endPoint = this.display ? this.height - 1.5 * this.fontSize - 5 : this.height, this.startPoint += this.padding, this.endPoint -= this.padding;
			var a, b = this.endPoint - this.startPoint;
			for (this.calculateYRange(b), this.buildYLabels(), this.calculateXLabelRotation(); b > this.endPoint - this.startPoint;) {
				b = this.endPoint - this.startPoint, a = this.yLabelWidth, this.calculateYRange(b), this.buildYLabels(), a < this.yLabelWidth && this.calculateXLabelRotation()
			}
		},
		calculateXLabelRotation: function() {
			this.ctx.font = this.font;
			var d, g, c = this.ctx.measureText(this.xLabels[0]).width,
				k = this.ctx.measureText(this.xLabels[this.xLabels.length - 1]).width;
			if (this.xScalePaddingRight = k / 2 + 3, this.xScalePaddingLeft = c / 2 > this.yLabelWidth + 10 ? c / 2 : this.yLabelWidth + 10, this.xLabelRotation = 0, this.display) {
				var h, f = K(this.ctx, this.font, this.xLabels);
				this.xLabelWidth = f;
				for (var b = Math.floor(this.calculateX(1) - this.calculateX(0)) - 6; this.xLabelWidth > b && 0 === this.xLabelRotation || this.xLabelWidth > b && this.xLabelRotation <= 90 && this.xLabelRotation > 0;) {
					h = Math.cos(J(this.xLabelRotation)), d = h * c, g = h * k, d + this.fontSize / 2 > this.yLabelWidth + 8 && (this.xScalePaddingLeft = d + this.fontSize / 2), this.xScalePaddingRight = this.fontSize / 2, this.xLabelRotation++, this.xLabelWidth = h * f
				}
				this.xLabelRotation > 0 && (this.endPoint -= Math.sin(J(this.xLabelRotation)) * f + 3)
			} else {
				this.xLabelWidth = 0, this.xScalePaddingRight = this.padding, this.xScalePaddingLeft = this.padding
			}
		},
		calculateYRange: am,
		drawingArea: function() {
			return this.startPoint - this.endPoint
		},
		calculateY: function(a) {
			var b = this.drawingArea() / (this.min - this.max);
			return this.endPoint - b * (a - this.min)
		},
		calculateX: function(b) {
			var c = (this.xLabelRotation > 0, this.width - (this.xScalePaddingLeft + this.xScalePaddingRight)),
				a = c / Math.max(this.valuesCount - (this.offsetGridLines ? 0 : 1), 1),
				d = a * b + this.xScalePaddingLeft;
			return this.offsetGridLines && (d += a / 2), Math.round(d)
		},
		update: function(a) {
			ac.extend(this, a), this.fit()
		},
		draw: function() {
			var b = this.ctx,
				d = (this.endPoint - this.startPoint) / this.steps,
				a = Math.round(this.xScalePaddingLeft);
			if (this.display) {
				b.fillStyle = this.textColor, b.font = this.font;
				var c = this.showBeyondLine ? 5 : 0;
				ab(this.yLabels, function(m, f) {
					var k = this.endPoint - d * f,
						e = Math.round(k),
						g = this.showHorizontalLines;
					b.textAlign = "right", b.textBaseline = "middle", this.showLabels && b.fillText(m, a - 10, k), 0 !== f || g || (g = !0), g && b.beginPath(), f > 0 ? (b.lineWidth = this.gridLineWidth, b.strokeStyle = this.gridLineColor) : (b.lineWidth = this.lineWidth, b.strokeStyle = this.lineColor), e += ac.aliasPixel(b.lineWidth), g && (b.moveTo(a, e), b.lineTo(this.width, e), b.stroke(), b.closePath()), b.lineWidth = this.lineWidth, b.strokeStyle = this.lineColor, b.beginPath(), b.moveTo(a - c, e), b.lineTo(a, e), b.stroke(), b.closePath()
				}, this), ab(this.xLabels, function(k, g) {
					var m = this.calculateX(g) + an(this.lineWidth),
						l = this.calculateX(g - (this.offsetGridLines ? 0.5 : 0)) + an(this.lineWidth),
						f = this.xLabelRotation > 0,
						h = this.showVerticalLines;
					0 !== g || h || (h = !0), h && b.beginPath(), g > 0 ? (b.lineWidth = this.gridLineWidth, b.strokeStyle = this.gridLineColor) : (b.lineWidth = this.lineWidth, b.strokeStyle = this.lineColor), h && (b.moveTo(l, this.endPoint), b.lineTo(l, this.startPoint - 3), b.stroke(), b.closePath()), b.lineWidth = this.lineWidth, b.strokeStyle = this.lineColor, b.beginPath(), b.moveTo(l, this.endPoint), b.lineTo(l, this.endPoint + c), b.stroke(), b.closePath(), b.save(), b.translate(m, f ? this.endPoint + 12 : this.endPoint + 8), b.rotate(J(this.xLabelRotation) * -1), b.font = this.font, b.textAlign = f ? "right" : "center", b.textBaseline = f ? "middle" : "top", b.fillText(k, 0, 0), b.restore()
				}, this)
			}
		}
	}), ag.RadialScale = ag.Element.extend({
		initialize: function() {
			this.size = ad([this.height, this.width]), this.drawingArea = this.display ? this.size / 2 - (this.fontSize / 2 + this.backdropPaddingY) : this.size / 2
		},
		calculateCenterOffset: function(a) {
			var b = this.drawingArea / (this.max - this.min);
			return (a - this.min) * b
		},
		update: function() {
			this.lineArc ? this.drawingArea = this.display ? this.size / 2 - (this.fontSize / 2 + this.backdropPaddingY) : this.size / 2 : this.setScaleSize(), this.buildYLabels()
		},
		buildYLabels: function() {
			this.yLabels = [];
			for (var a = Q(this.stepValue), b = 0; b <= this.steps; b++) {
				this.yLabels.push(N(this.templateString, {
					value: (this.min + b * this.stepValue).toFixed(a)
				}))
			}
		},
		getCircumference: function() {
			return 2 * Math.PI / this.valuesCount
		},
		setScaleSize: function() {
			var L, z, w, m, k, M, F, b, v, x, D, C, I = ad([this.height / 2 - this.pointLabelFontSize - 5, this.width / 2]),
				f = this.width,
				y = 0;
			for (this.ctx.font = H(this.pointLabelFontSize, this.pointLabelFontStyle, this.pointLabelFontFamily), z = 0; z < this.valuesCount; z++) {
				L = this.getPointPosition(z, I), w = this.ctx.measureText(N(this.templateString, {
					value: this.labels[z]
				})).width + 5, 0 === z || z === this.valuesCount / 2 ? (m = w / 2, L.x + m > f && (f = L.x + m, k = z), L.x - m < y && (y = L.x - m, F = z)) : z < this.valuesCount / 2 ? L.x + w > f && (f = L.x + w, k = z) : z > this.valuesCount / 2 && L.x - w < y && (y = L.x - w, F = z)
			}
			v = y, x = Math.ceil(f - this.width), M = this.getIndexAngle(k), b = this.getIndexAngle(F), D = x / Math.sin(M + Math.PI / 2), C = v / Math.sin(b + Math.PI / 2), D = aj(D) ? D : 0, C = aj(C) ? C : 0, this.drawingArea = I - (C + D) / 2, this.setCenterPoint(C, D)
		},
		setCenterPoint: function(b, c) {
			var a = this.width - c - this.drawingArea,
				d = b + this.drawingArea;
			this.xCenter = (d + a) / 2, this.yCenter = this.height / 2
		},
		getIndexAngle: function(a) {
			var b = 2 * Math.PI / this.valuesCount;
			return a * b - Math.PI / 2
		},
		getPointPosition: function(b, c) {
			var a = this.getIndexAngle(b);
			return {
				x: Math.cos(a) * c + this.xCenter,
				y: Math.sin(a) * c + this.yCenter
			}
		},
		draw: function() {
			if (this.display) {
				var o = this.ctx;
				if (ab(this.yLabels, function(w, l) {
					if (l > 0) {
						var y, x = l * (this.drawingArea / this.steps),
							u = this.yCenter - x;
						if (this.lineWidth > 0) {
							if (o.strokeStyle = this.lineColor, o.lineWidth = this.lineWidth, this.lineArc) {
								o.beginPath(), o.arc(this.xCenter, this.yCenter, x, 0, 2 * Math.PI), o.closePath(), o.stroke()
							} else {
								o.beginPath();
								for (var h = 0; h < this.valuesCount; h++) {
									y = this.getPointPosition(h, this.calculateCenterOffset(this.min + l * this.stepValue)), 0 === h ? o.moveTo(y.x, y.y) : o.lineTo(y.x, y.y)
								}
								o.closePath(), o.stroke()
							}
						}
						if (this.showLabels) {
							if (o.font = H(this.fontSize, this.fontStyle, this.fontFamily), this.showLabelBackdrop) {
								var v = o.measureText(w).width;
								o.fillStyle = this.backdropColor, o.fillRect(this.xCenter - v / 2 - this.backdropPaddingX, u - this.fontSize / 2 - this.backdropPaddingY, v + 2 * this.backdropPaddingX, this.fontSize + 2 * this.backdropPaddingY)
							}
							o.textAlign = "center", o.textBaseline = "middle", o.fillStyle = this.fontColor, o.fillText(w, this.xCenter, u)
						}
					}
				}, this), !this.lineArc) {
					o.lineWidth = this.angleLineWidth, o.strokeStyle = this.angleLineColor;
					for (var k = this.valuesCount - 1; k >= 0; k--) {
						if (this.angleLineWidth > 0) {
							var f = this.getPointPosition(k, this.calculateCenterOffset(this.max));
							o.beginPath(), o.moveTo(this.xCenter, this.yCenter), o.lineTo(f.x, f.y), o.stroke(), o.closePath()
						}
						var c = this.getPointPosition(k, this.calculateCenterOffset(this.max) + 5);
						o.font = H(this.pointLabelFontSize, this.pointLabelFontStyle, this.pointLabelFontFamily), o.fillStyle = this.pointLabelFontColor;
						var p = this.labels.length,
							m = this.labels.length / 2,
							b = m / 2,
							d = k < b || k > p - b,
							g = k === b || k === p - b;
						0 === k ? o.textAlign = "center" : k === m ? o.textAlign = "center" : k < m ? o.textAlign = "left" : o.textAlign = "right", g ? o.textBaseline = "middle" : d ? o.textBaseline = "bottom" : o.textBaseline = "top", o.fillText(this.labels[k], c.x, c.y)
					}
				}
			}
		}
	}), ac.addEvent(window, "resize", function() {
		var a;
		return function() {
			clearTimeout(a), a = setTimeout(function() {
				ab(ag.instances, function(b) {
					b.options.responsive && b.resize(b.render, !0)
				})
			}, 50)
		}
	}()), aa ? define(function() {
		return ag
	}) : "object" == typeof module && module.exports && (module.exports = ag), ak.Chart = ag, V.fn.chart = function() {
		var a = [];
		return this.each(function() {
			a.push(new ag(this.getContext("2d")))
		}), 1 === a.length ? a[0] : a
	}
}.call(this, jQuery), function(b) {
	var c = b && b.zui ? b.zui : this,
		a = c.Chart,
		f = a.helpers,
		d = {
			scaleShowGridLines: !0,
			scaleGridLineColor: "rgba(0,0,0,.05)",
			scaleGridLineWidth: 1,
			scaleShowHorizontalLines: !0,
			scaleShowBeyondLine: !0,
			scaleShowVerticalLines: !0,
			bezierCurve: !0,
			bezierCurveTension: 0.4,
			pointDot: !0,
			pointDotRadius: 4,
			pointDotStrokeWidth: 1,
			pointHitDetectionRadius: 20,
			datasetStroke: !0,
			datasetStrokeWidth: 2,
			datasetFill: !0,
			legendTemplate: '<ul class="<%=name.toLowerCase()%>-legend"><% for (var i=0; i<datasets.length; i++){%><li><span style="background-color:<%=datasets[i].strokeColor%>"></span><%if(datasets[i].label){%><%=datasets[i].label%><%}%></li><%}%></ul>'
		};
	a.Type.extend({
		name: "Line",
		defaults: d,
		initialize: function(g) {
			this.PointClass = a.Point.extend({
				strokeWidth: this.options.pointDotStrokeWidth,
				radius: this.options.pointDotRadius,
				display: this.options.pointDot,
				hitDetectionRadius: this.options.pointHitDetectionRadius,
				ctx: this.chart.ctx,
				inRange: function(e) {
					return Math.pow(e - this.x, 2) < Math.pow(this.radius + this.hitDetectionRadius, 2)
				}
			}), this.datasets = [], this.options.showTooltips && f.bindEvents(this, this.options.tooltipEvents, function(h) {
				var j = "mouseout" !== h.type ? this.getPointsAtEvent(h) : [];
				this.eachPoints(function(e) {
					e.restore(["fillColor", "strokeColor"])
				}), f.each(j, function(e) {
					e.fillColor = e.highlightFill, e.strokeColor = e.highlightStroke
				}), this.showTooltip(j)
			}), f.each(g.datasets, function(h) {
				if (b.zui && b.zui.Color && b.zui.Color.get) {
					var k = b.zui.Color.get(h.color),
						j = k.toCssStr();
					h.fillColor || (h.fillColor = k.clone().fade(20).toCssStr()), h.strokeColor || (h.strokeColor = j), h.pointColor || (h.pointColor = j), h.pointStrokeColor || (h.pointStrokeColor = "#fff"), h.pointHighlightFill || (h.pointHighlightFill = "#fff"), h.pointHighlightStroke || (h.pointHighlightStroke = j)
				}
				var e = {
					label: h.label || null,
					fillColor: h.fillColor,
					strokeColor: h.strokeColor,
					pointColor: h.pointColor,
					pointStrokeColor: h.pointStrokeColor,
					showTooltips: h.showTooltips !== !1,
					points: []
				};
				this.datasets.push(e), f.each(h.data, function(l, m) {
					e.points.push(new this.PointClass({
						value: l,
						label: g.labels[m],
						datasetLabel: h.label,
						strokeColor: h.pointStrokeColor,
						fillColor: h.pointColor,
						highlightFill: h.pointHighlightFill || h.pointColor,
						highlightStroke: h.pointHighlightStroke || h.pointStrokeColor
					}))
				}, this), this.buildScale(g.labels), this.eachPoints(function(l, m) {
					f.extend(l, {
						x: this.scale.calculateX(m),
						y: this.scale.endPoint
					}), l.save()
				}, this)
			}, this), this.render()
		},
		update: function() {
			this.scale.update(), f.each(this.activeElements, function(e) {
				e.restore(["fillColor", "strokeColor"])
			}), this.eachPoints(function(e) {
				e.save()
			}), this.render()
		},
		eachPoints: function(e) {
			f.each(this.datasets, function(g) {
				f.each(g.points, e, this)
			}, this)
		},
		getPointsAtEvent: function(h) {
			var j = [],
				g = f.getRelativePosition(h);
			return f.each(this.datasets, function(e) {
				f.each(e.points, function(k) {
					k.inRange(g.x, g.y) && j.push(k)
				})
			}, this), j
		},
		buildScale: function(g) {
			var j = this,
				k = function() {
					var e = [];
					return j.eachPoints(function(l) {
						e.push(l.value)
					}), e
				},
				h = {
					templateString: this.options.scaleLabel,
					height: this.chart.height,
					width: this.chart.width,
					ctx: this.chart.ctx,
					textColor: this.options.scaleFontColor,
					fontSize: this.options.scaleFontSize,
					fontStyle: this.options.scaleFontStyle,
					fontFamily: this.options.scaleFontFamily,
					valuesCount: g.length,
					beginAtZero: this.options.scaleBeginAtZero,
					integersOnly: this.options.scaleIntegersOnly,
					calculateYRange: function(l) {
						var m = f.calculateScaleRange(k(), l, this.fontSize, this.beginAtZero, this.integersOnly);
						f.extend(this, m)
					},
					xLabels: g,
					font: f.fontString(this.options.scaleFontSize, this.options.scaleFontStyle, this.options.scaleFontFamily),
					lineWidth: this.options.scaleLineWidth,
					lineColor: this.options.scaleLineColor,
					showHorizontalLines: this.options.scaleShowHorizontalLines,
					showVerticalLines: this.options.scaleShowVerticalLines,
					showBeyondLine: this.options.scaleShowBeyondLine,
					gridLineWidth: this.options.scaleShowGridLines ? this.options.scaleGridLineWidth : 0,
					gridLineColor: this.options.scaleShowGridLines ? this.options.scaleGridLineColor : "rgba(0,0,0,0)",
					padding: this.options.showScale ? 0 : this.options.pointDotRadius + this.options.pointDotStrokeWidth,
					showLabels: this.options.scaleShowLabels,
					display: this.options.showScale
				};
			this.options.scaleOverride && f.extend(h, {
				calculateYRange: f.noop,
				steps: this.options.scaleSteps,
				stepValue: this.options.scaleStepWidth,
				min: this.options.scaleStartValue,
				max: this.options.scaleStartValue + this.options.scaleSteps * this.options.scaleStepWidth
			}), this.scale = new a.Scale(h)
		},
		addData: function(g, h) {
			f.each(g, function(j, e) {
				this.datasets[e].points.push(new this.PointClass({
					value: j,
					label: h,
					datasetLabel: this.datasets[e].label,
					x: this.scale.calculateX(this.scale.valuesCount + 1),
					y: this.scale.endPoint,
					strokeColor: this.datasets[e].pointStrokeColor,
					fillColor: this.datasets[e].pointColor
				}))
			}, this), this.scale.addXLabel(h), this.update()
		},
		removeData: function() {
			this.scale.removeXLabel(), f.each(this.datasets, function(e) {
				e.points.shift()
			}, this), this.update()
		},
		reflow: function() {
			var e = f.extend({
				height: this.chart.height,
				width: this.chart.width
			});
			this.scale.update(e)
		},
		draw: function(j) {
			var l = j || 1;
			this.clear();
			var h = this.chart.ctx,
				m = function(e) {
					return null !== e.value
				},
				k = function(o, p, n) {
					return f.findNextWhere(p, m, n) || o
				},
				g = function(o, p, n) {
					return f.findPreviousWhere(p, m, n) || o
				};
			this.scale.draw(l), f.each(this.datasets, function(e) {
				var n = f.where(e.points, m);
				f.each(e.points, function(p, o) {
					p.hasValue() && p.transition({
						y: this.scale.calculateY(p.value),
						x: this.scale.calculateX(o)
					}, l)
				}, this), this.options.bezierCurve && f.each(n, function(p, q) {
					var o = q > 0 && q < n.length - 1 ? this.options.bezierCurveTension : 0;
					p.controlPoints = f.splineCurve(g(p, n, q), p, k(p, n, q), o), p.controlPoints.outer.y > this.scale.endPoint ? p.controlPoints.outer.y = this.scale.endPoint : p.controlPoints.outer.y < this.scale.startPoint && (p.controlPoints.outer.y = this.scale.startPoint), p.controlPoints.inner.y > this.scale.endPoint ? p.controlPoints.inner.y = this.scale.endPoint : p.controlPoints.inner.y < this.scale.startPoint && (p.controlPoints.inner.y = this.scale.startPoint)
				}, this), h.lineWidth = this.options.datasetStrokeWidth, h.strokeStyle = e.strokeColor, h.beginPath(), f.each(n, function(o, p) {
					if (0 === p) {
						h.moveTo(o.x, o.y)
					} else {
						if (this.options.bezierCurve) {
							var q = g(o, n, p);
							h.bezierCurveTo(q.controlPoints.outer.x, q.controlPoints.outer.y, o.controlPoints.inner.x, o.controlPoints.inner.y, o.x, o.y)
						} else {
							h.lineTo(o.x, o.y)
						}
					}
				}, this), h.stroke(), this.options.datasetFill && n.length > 0 && (h.lineTo(n[n.length - 1].x, this.scale.endPoint), h.lineTo(n[0].x, this.scale.endPoint), h.fillStyle = e.fillColor, h.closePath(), h.fill()), f.each(n, function(o) {
					o.draw()
				})
			}, this)
		}
	}), b.fn.lineChart = function(g, j) {
		var h = [];
		return this.each(function() {
			var e = b(this);
			h.push(new a(this.getContext("2d")).Line(g, b.extend(e.data(), j)))
		}), 1 === h.length ? h[0] : h
	}
}.call(this, jQuery), function(b) {
	var c = b && b.zui ? b.zui : this,
		a = c.Chart,
		f = a.helpers,
		d = {
			segmentShowStroke: !0,
			segmentStrokeColor: "#fff",
			segmentStrokeWidth: 1,
			percentageInnerCutout: 50,
			scaleShowLabels: !1,
			scaleLabel: "<%=value%>",
			scaleLabelPlacement: "auto",
			animationSteps: 60,
			animationEasing: "easeOutBounce",
			animateRotate: !0,
			animateScale: !1,
			legendTemplate: '<ul class="<%=name.toLowerCase()%>-legend"><% for (var i=0; i<segments.length; i++){%><li><span style="background-color:<%=segments[i].fillColor%>"></span><%if(segments[i].label){%><%=segments[i].label%><%}%></li><%}%></ul>'
		};
	a.Type.extend({
		name: "Doughnut",
		defaults: d,
		initialize: function(e) {
			this.segments = [], this.outerRadius = (f.min([this.chart.width, this.chart.height]) - this.options.segmentStrokeWidth / 2) / 2, this.SegmentArc = a.Arc.extend({
				ctx: this.chart.ctx,
				x: this.chart.width / 2,
				y: this.chart.height / 2
			}), this.options.showTooltips && f.bindEvents(this, this.options.tooltipEvents, function(g) {
				var h = "mouseout" !== g.type ? this.getSegmentsAtEvent(g) : [];
				f.each(this.segments, function(j) {
					j.restore(["fillColor"])
				}), f.each(h, function(j) {
					j.fillColor = j.highlightColor
				}), this.showTooltip(h)
			}), this.calculateTotal(e), f.each(e, function(g, h) {
				this.addData(g, h, !0)
			}, this), this.render()
		},
		getSegmentsAtEvent: function(h) {
			var j = [],
				g = f.getRelativePosition(h);
			return f.each(this.segments, function(e) {
				e.inRange(g.x, g.y) && j.push(e)
			}, this), j
		},
		addData: function(j, g, l) {
			if (b.zui && b.zui.Color && b.zui.Color.get) {
				var k = new b.zui.Color.get(j.color);
				j.color = k.toCssStr(), j.highlight || (j.highlight = k.lighten(5).toCssStr())
			}
			var h = g || this.segments.length;
			this.segments.splice(h, 0, new this.SegmentArc({
				id: "undefined" == typeof j.id ? h : j.id,
				value: j.value,
				outerRadius: this.options.animateScale ? 0 : this.outerRadius,
				innerRadius: this.options.animateScale ? 0 : this.outerRadius / 100 * this.options.percentageInnerCutout,
				fillColor: j.color,
				highlightColor: j.highlight || j.color,
				showStroke: this.options.segmentShowStroke,
				strokeWidth: this.options.segmentStrokeWidth,
				strokeColor: this.options.segmentStrokeColor,
				startAngle: 1.5 * Math.PI,
				circumference: this.options.animateRotate ? 0 : this.calculateCircumference(j.value),
				showLabel: j.showLabel !== !1,
				circleBeginEnd: j.circleBeginEnd,
				label: j.label
			})), l || (this.reflow(), this.update())
		},
		calculateCircumference: function(e) {
			return 2 * Math.PI * (Math.abs(e) / this.total)
		},
		calculateTotal: function(e) {
			this.total = 0, f.each(e, function(g) {
				this.total += Math.abs(g.value)
			}, this)
		},
		update: function() {
			this.calculateTotal(this.segments), f.each(this.activeElements, function(e) {
				e.restore(["fillColor"])
			}), f.each(this.segments, function(e) {
				e.save()
			}), this.render()
		},
		removeData: function(g) {
			var h = f.isNumber(g) ? g : this.segments.length - 1;
			this.segments.splice(h, 1), this.reflow(), this.update()
		},
		reflow: function() {
			f.extend(this.SegmentArc.prototype, {
				x: this.chart.width / 2,
				y: this.chart.height / 2
			}), this.outerRadius = (f.min([this.chart.width, this.chart.height]) - this.options.segmentStrokeWidth / 2) / 2, f.each(this.segments, function(e) {
				e.update({
					outerRadius: this.outerRadius,
					innerRadius: this.outerRadius / 100 * this.options.percentageInnerCutout
				})
			}, this)
		},
		drawLabel: function(G, C, z) {
			var q = this.options,
				K = (G.endAngle + G.startAngle) / 2,
				w = q.scaleLabelPlacement;
			"inside" !== w && "outside" !== w && this.chart.width - this.chart.height > 50 && G.circumference < Math.PI / 18 && (w = "outside");
			var B = Math.cos(K) * G.outerRadius,
				D = Math.sin(K) * G.outerRadius,
				I = f.template(q.scaleLabel, {
					value: "undefined" == typeof C ? G.value : Math.round(C * G.value),
					label: G.label
				}),
				H = this.chart.ctx;
			H.font = f.fontString(q.scaleFontSize, q.scaleFontStyle, q.scaleFontFamily), H.textBaseline = "middle", H.textAlign = "center";
			var n = (H.measureText(I).width, this.chart.width / 2),
				x = this.chart.height / 2;
			if ("outside" === w) {
				var F = B >= 0,
					E = B + n,
					A = D + x;
				H.textAlign = F ? "left" : "right", H.measureText(I).width, B = F ? Math.max(n + G.outerRadius + 10, B + 30 + n) : Math.min(n - G.outerRadius - 10, B - 30 + n);
				var k = q.scaleFontSize * (q.scaleLineHeight || 1),
					j = Math.round((0.8 * D + x) / k) + 1,
					J = (Math.floor(this.chart.width / k) + 1, F ? 1 : -1);
				if (z[j * J] && (j > 1 ? j-- : j++), z[j * J]) {
					return
				}
				D = (j - 1) * k + q.scaleFontSize / 2, z[j * J] = !0, H.beginPath(), H.moveTo(E, A), H.lineTo(B, D), B = F ? B + 5 : B - 5, H.lineTo(B, D), H.strokeStyle = b.zui && b.zui.Color ? new b.zui.Color(G.fillColor).fade(40).toCssStr() : G.fillColor, H.strokeWidth = q.scaleLineWidth, H.stroke(), H.fillStyle = G.fillColor
			} else {
				B = 0.7 * B + n, D = 0.7 * D + x, H.fillStyle = b.zui && b.zui.Color ? new b.zui.Color(G.fillColor).contrast().toCssStr() : "#fff"
			}
			H.fillText(I, B, D)
		},
		draw: function(h) {
			var j = h ? h : 1;
			this.clear();
			var g;
			if (f.each(this.segments, function(l, e) {
				l.transition({
					circumference: this.calculateCircumference(l.value),
					outerRadius: this.outerRadius,
					innerRadius: this.outerRadius / 100 * this.options.percentageInnerCutout
				}, j), l.endAngle = l.startAngle + l.circumference, this.options.reverseDrawOrder || l.draw(), 0 === e && (l.startAngle = 1.5 * Math.PI), e < this.segments.length - 1 && (this.segments[e + 1].startAngle = l.endAngle)
			}, this), this.options.reverseDrawOrder && f.each(this.segments.slice().reverse(), function(l, m) {
				l.draw()
			}, this), this.options.scaleShowLabels) {
				var k = this.segments.slice().sort(function(l, m) {
					return m.value - l.value
				}),
					g = {};
				f.each(k, function(l, m) {
					l.showLabel && this.drawLabel(l, h, g)
				}, this)
			}
		}
	}), a.types.Doughnut.extend({
		name: "Pie",
		defaults: f.merge(d, {
			percentageInnerCutout: 0
		})
	}), b.fn.pieChart = function(g, j) {
		var h = [];
		return this.each(function() {
			var e = b(this);
			h.push(new a(this.getContext("2d")).Pie(g, b.extend(e.data(), j)))
		}), 1 === h.length ? h[0] : h
	}, b.fn.doughnutChart = function(g, j) {
		var h = [];
		return this.each(function() {
			var e = b(this);
			h.push(new a(this.getContext("2d")).Doughnut(g, b.extend(e.data(), j)))
		}), 1 === h.length ? h[0] : h
	}
}.call(this, jQuery), function(b) {
	var c = b && b.zui ? b.zui : this,
		a = c.Chart,
		f = a.helpers,
		d = {
			scaleBeginAtZero: !0,
			scaleShowGridLines: !0,
			scaleGridLineColor: "rgba(0,0,0,.05)",
			scaleGridLineWidth: 1,
			scaleShowHorizontalLines: !0,
			scaleShowVerticalLines: !0,
			scaleShowBeyondLine: !0,
			barShowStroke: !0,
			barStrokeWidth: 1,
			scaleValuePlacement: "auto",
			barValueSpacing: 5,
			barDatasetSpacing: 1,
			legendTemplate: '<ul class="<%=name.toLowerCase()%>-legend"><% for (var i=0; i<datasets.length; i++){%><li><span style="background-color:<%=datasets[i].fillColor%>"></span><%if(datasets[i].label){%><%=datasets[i].label%><%}%></li><%}%></ul>'
		};
	a.Type.extend({
		name: "Bar",
		defaults: d,
		initialize: function(g) {
			var h = this.options;
			this.ScaleClass = a.Scale.extend({
				offsetGridLines: !0,
				calculateBarX: function(l, o, k) {
					var p = this.calculateBaseWidth(),
						m = this.calculateX(k) - p / 2,
						j = this.calculateBarWidth(l);
					return m + j * o + o * h.barDatasetSpacing + j / 2
				},
				calculateBaseWidth: function() {
					return this.calculateX(1) - this.calculateX(0) - 2 * h.barValueSpacing
				},
				calculateBarWidth: function(j) {
					var k = this.calculateBaseWidth() - (j - 1) * h.barDatasetSpacing;
					return k / j
				}
			}), this.datasets = [], this.options.showTooltips && f.bindEvents(this, this.options.tooltipEvents, function(j) {
				var k = "mouseout" !== j.type ? this.getBarsAtEvent(j) : [];
				this.eachBars(function(e) {
					e.restore(["fillColor", "strokeColor"])
				}), f.each(k, function(e) {
					e.fillColor = e.highlightFill, e.strokeColor = e.highlightStroke
				}), this.showTooltip(k)
			}), this.BarClass = a.Rectangle.extend({
				strokeWidth: this.options.barStrokeWidth,
				showStroke: this.options.barShowStroke,
				ctx: this.chart.ctx
			}), f.each(g.datasets, function(j, m) {
				if (b.zui && b.zui.Color && b.zui.Color.get) {
					var k = b.zui.Color.get(j.color),
						e = k.toCssStr();
					j.fillColor || (j.fillColor = k.clone().fade(50).toCssStr()), j.strokeColor || (j.strokeColor = e)
				}
				var l = {
					label: j.label || null,
					fillColor: j.fillColor,
					strokeColor: j.strokeColor,
					bars: []
				};
				this.datasets.push(l), f.each(j.data, function(o, p) {
					l.bars.push(new this.BarClass({
						value: o,
						label: g.labels[p],
						datasetLabel: j.label,
						strokeColor: j.strokeColor,
						fillColor: j.fillColor,
						highlightFill: j.highlightFill || j.fillColor,
						highlightStroke: j.highlightStroke || j.strokeColor
					}))
				}, this)
			}, this), this.buildScale(g.labels), this.BarClass.prototype.base = this.scale.endPoint, this.eachBars(function(k, l, j) {
				f.extend(k, {
					width: this.scale.calculateBarWidth(this.datasets.length),
					x: this.scale.calculateBarX(this.datasets.length, j, l),
					y: this.scale.endPoint
				}), k.save()
			}, this), this.render()
		},
		update: function() {
			this.scale.update(), f.each(this.activeElements, function(e) {
				e.restore(["fillColor", "strokeColor"])
			}), this.eachBars(function(e) {
				e.save()
			}), this.render()
		},
		eachBars: function(e) {
			f.each(this.datasets, function(h, g) {
				f.each(h.bars, e, this, g)
			}, this)
		},
		getBarsAtEvent: function(j) {
			for (var l, h = [], m = f.getRelativePosition(j), k = function(e) {
					h.push(e.bars[l])
				}, g = 0; g < this.datasets.length; g++) {
				for (l = 0; l < this.datasets[g].bars.length; l++) {
					if (this.datasets[g].bars[l].inRange(m.x, m.y)) {
						return f.each(this.datasets, k), h
					}
				}
			}
			return h
		},
		buildScale: function(h) {
			var j = this,
				g = function() {
					var e = [];
					return j.eachBars(function(l) {
						e.push(l.value)
					}), e
				},
				k = {
					templateString: this.options.scaleLabel,
					height: this.chart.height,
					width: this.chart.width,
					ctx: this.chart.ctx,
					textColor: this.options.scaleFontColor,
					fontSize: this.options.scaleFontSize,
					fontStyle: this.options.scaleFontStyle,
					fontFamily: this.options.scaleFontFamily,
					valuesCount: h.length,
					beginAtZero: this.options.scaleBeginAtZero,
					integersOnly: this.options.scaleIntegersOnly,
					calculateYRange: function(l) {
						var m = f.calculateScaleRange(g(), l, this.fontSize, this.beginAtZero, this.integersOnly);
						f.extend(this, m)
					},
					xLabels: h,
					font: f.fontString(this.options.scaleFontSize, this.options.scaleFontStyle, this.options.scaleFontFamily),
					lineWidth: this.options.scaleLineWidth,
					lineColor: this.options.scaleLineColor,
					showHorizontalLines: this.options.scaleShowHorizontalLines,
					showVerticalLines: this.options.scaleShowVerticalLines,
					showBeyondLine: this.options.scaleShowBeyondLine,
					gridLineWidth: this.options.scaleShowGridLines ? this.options.scaleGridLineWidth : 0,
					gridLineColor: this.options.scaleShowGridLines ? this.options.scaleGridLineColor : "rgba(0,0,0,0)",
					padding: this.options.showScale ? 0 : this.options.barShowStroke ? this.options.barStrokeWidth : 0,
					showLabels: this.options.scaleShowLabels,
					display: this.options.showScale
				};
			this.options.scaleOverride && f.extend(k, {
				calculateYRange: f.noop,
				steps: this.options.scaleSteps,
				stepValue: this.options.scaleStepWidth,
				min: this.options.scaleStartValue,
				max: this.options.scaleStartValue + this.options.scaleSteps * this.options.scaleStepWidth
			}), this.scale = new this.ScaleClass(k)
		},
		addData: function(g, h) {
			f.each(g, function(j, e) {
				this.datasets[e].bars.push(new this.BarClass({
					value: j,
					label: h,
					x: this.scale.calculateBarX(this.datasets.length, e, this.scale.valuesCount + 1),
					y: this.scale.endPoint,
					width: this.scale.calculateBarWidth(this.datasets.length),
					base: this.scale.endPoint,
					strokeColor: this.datasets[e].strokeColor,
					fillColor: this.datasets[e].fillColor
				}))
			}, this), this.scale.addXLabel(h), this.update()
		},
		removeData: function() {
			this.scale.removeXLabel(), f.each(this.datasets, function(e) {
				e.bars.shift()
			}, this), this.update()
		},
		reflow: function() {
			f.extend(this.BarClass.prototype, {
				y: this.scale.endPoint,
				base: this.scale.endPoint
			});
			var e = f.extend({
				height: this.chart.height,
				width: this.chart.width
			});
			this.scale.update(e)
		},
		drawLabel: function(h, k) {
			var g = this.options;
			k = k || g.scaleValuePlacement, k = k ? k.toLowerCase() : "auto", "auto" === k && (k = h.y < 15 ? "insdie" : "outside");
			var l = "insdie" === k ? h.y + 10 : h.y - 10,
				j = this.chart.ctx;
			j.font = f.fontString(g.scaleFontSize, g.scaleFontStyle, g.scaleFontFamily), j.textBaseline = "middle", j.textAlign = "center", j.fillStyle = g.scaleFontColor, j.fillText(h.value, h.x, l)
		},
		draw: function(h) {
			var j = h || 1;
			this.clear();
			this.chart.ctx;
			this.scale.draw(j);
			var g = this.options.scaleShowLabels && this.options.scaleValuePlacement;
			f.each(this.datasets, function(e, k) {
				f.each(e.bars, function(l, m) {
					l.hasValue() && (l.base = this.scale.endPoint, l.transition({
						x: this.scale.calculateBarX(this.datasets.length, k, m),
						y: this.scale.calculateY(l.value),
						width: this.scale.calculateBarWidth(this.datasets.length)
					}, j).draw()), g && this.drawLabel(l)
				}, this)
			}, this)
		}
	}), b.fn.barChart = function(g, j) {
		var h = [];
		return this.each(function() {
			var e = b(this);
			h.push(new a(this.getContext("2d")).Bar(g, b.extend(e.data(), j)))
		}), 1 === h.length ? h[0] : h
	}
}.call(this, jQuery),
/*!
 * Datetimepicker for Bootstrap
 * Copyright 2012 Stefan Petre
 * Licensed under the Apache License v2.0
 */
!
function(b) {
	function c() {
		return new Date(Date.UTC.apply(Date, arguments))
	}
	var a = function(j, g) {
			var h = this;
			this.element = b(j), this.language = (g.language || this.element.data("date-language") || (b.zui && b.zui.clientLang ? b.zui.clientLang().replace("_", "-") : "zh-cn")).toLowerCase(), this.language = this.language in f ? this.language : "en", this.isRTL = f[this.language].rtl || !1, this.formatType = g.formatType || this.element.data("format-type") || "standard", this.format = d.parseFormat(g.format || this.element.data("date-format") || f[this.language].format || d.getDefaultFormat(this.formatType, "input"), this.formatType), this.isInline = !1, this.isVisible = !1, this.isInput = this.element.is("input"), this.component = !! this.element.is(".date") && this.element.find(".input-group-addon .icon-th, .input-group-addon .icon-time, .input-group-addon .icon-calendar").parent(), this.componentReset = !! this.element.is(".date") && this.element.find(".input-group-addon .icon-remove").parent(), this.hasInput = this.component && this.element.find("input").length, this.component && 0 === this.component.length && (this.component = !1), this.linkField = g.linkField || this.element.data("link-field") || !1, this.linkFormat = d.parseFormat(g.linkFormat || this.element.data("link-format") || d.getDefaultFormat(this.formatType, "link"), this.formatType), this.minuteStep = g.minuteStep || this.element.data("minute-step") || 5, this.pickerPosition = g.pickerPosition || this.element.data("picker-position") || "bottom-right", this.showMeridian = g.showMeridian || this.element.data("show-meridian") || !1, this.initialDate = g.initialDate || new Date, this.pickerClass = g.eleClass, this.pickerId = g.eleId, this._attachEvents(), this.formatViewType = "datetime", "formatViewType" in g ? this.formatViewType = g.formatViewType : "formatViewType" in this.element.data() && (this.formatViewType = this.element.data("formatViewType")), this.minView = 0, "minView" in g ? this.minView = g.minView : "minView" in this.element.data() && (this.minView = this.element.data("min-view")), this.minView = d.convertViewMode(this.minView), this.maxView = d.modes.length - 1, "maxView" in g ? this.maxView = g.maxView : "maxView" in this.element.data() && (this.maxView = this.element.data("max-view")), this.maxView = d.convertViewMode(this.maxView), this.wheelViewModeNavigation = !1, "wheelViewModeNavigation" in g ? this.wheelViewModeNavigation = g.wheelViewModeNavigation : "wheelViewModeNavigation" in this.element.data() && (this.wheelViewModeNavigation = this.element.data("view-mode-wheel-navigation")), this.wheelViewModeNavigationInverseDirection = !1, "wheelViewModeNavigationInverseDirection" in g ? this.wheelViewModeNavigationInverseDirection = g.wheelViewModeNavigationInverseDirection : "wheelViewModeNavigationInverseDirection" in this.element.data() && (this.wheelViewModeNavigationInverseDirection = this.element.data("view-mode-wheel-navigation-inverse-dir")), this.wheelViewModeNavigationDelay = 100, "wheelViewModeNavigationDelay" in g ? this.wheelViewModeNavigationDelay = g.wheelViewModeNavigationDelay : "wheelViewModeNavigationDelay" in this.element.data() && (this.wheelViewModeNavigationDelay = this.element.data("view-mode-wheel-navigation-delay")), this.startViewMode = 2, "startView" in g ? this.startViewMode = g.startView : "startView" in this.element.data() && (this.startViewMode = this.element.data("start-view")), this.startViewMode = d.convertViewMode(this.startViewMode), this.viewMode = this.startViewMode, this.viewSelect = this.minView, "viewSelect" in g ? this.viewSelect = g.viewSelect : "viewSelect" in this.element.data() && (this.viewSelect = this.element.data("view-select")), this.viewSelect = d.convertViewMode(this.viewSelect), this.forceParse = !0, "forceParse" in g ? this.forceParse = g.forceParse : "dateForceParse" in this.element.data() && (this.forceParse = this.element.data("date-force-parse")), this.picker = b(d.template).appendTo(this.isInline ? this.element : "body").on({
				click: b.proxy(this.click, this),
				mousedown: b.proxy(this.mousedown, this)
			}), this.wheelViewModeNavigation && (b.fn.mousewheel ? this.picker.on({
				mousewheel: b.proxy(this.mousewheel, this)
			}) : console.log("Mouse Wheel event is not supported. Please include the jQuery Mouse Wheel plugin before enabling this option")), this.isInline ? this.picker.addClass("datetimepicker-inline") : this.picker.addClass("datetimepicker-dropdown-" + this.pickerPosition + " dropdown-menu"), this.isRTL && (this.picker.addClass("datetimepicker-rtl"), this.picker.find(".prev span, .next span").toggleClass("icon-arrow-left icon-arrow-right")), b(document).on("mousedown", function(k) {
				0 === b(k.target).closest(".datetimepicker").length && h.hide()
			}), this.autoclose = !1, "autoclose" in g ? this.autoclose = g.autoclose : "dateAutoclose" in this.element.data() && (this.autoclose = this.element.data("date-autoclose")), this.keyboardNavigation = !0, "keyboardNavigation" in g ? this.keyboardNavigation = g.keyboardNavigation : "dateKeyboardNavigation" in this.element.data() && (this.keyboardNavigation = this.element.data("date-keyboard-navigation")), this.todayBtn = g.todayBtn || this.element.data("date-today-btn") || !1, this.todayHighlight = g.todayHighlight || this.element.data("date-today-highlight") || !1, this.weekStart = (g.weekStart || this.element.data("date-weekstart") || f[this.language].weekStart || 0) % 7, this.weekEnd = (this.weekStart + 6) % 7, this.startDate = -(1 / 0), this.endDate = 1 / 0, this.daysOfWeekDisabled = [], this.setStartDate(g.startDate || this.element.data("date-startdate")), this.setEndDate(g.endDate || this.element.data("date-enddate")), this.setDaysOfWeekDisabled(g.daysOfWeekDisabled || this.element.data("date-days-of-week-disabled")), this.fillDow(), this.fillMonths(), this.update(), this.showMode(), this.isInline && this.show()
		};
	a.prototype = {
		constructor: a,
		_events: [],
		_attachEvents: function() {
			this._detachEvents(), this.isInput ? this._events = [
				[this.element,
				{
					focus: b.proxy(this.show, this),
					keyup: b.proxy(this.update, this),
					keydown: b.proxy(this.keydown, this)
				}]
			] : this.component && this.hasInput ? (this._events = [
				[this.element.find("input"),
				{
					focus: b.proxy(this.show, this),
					keyup: b.proxy(this.update, this),
					keydown: b.proxy(this.keydown, this)
				}],
				[this.component,
				{
					click: b.proxy(this.show, this)
				}]
			], this.componentReset && this._events.push([this.componentReset,
			{
				click: b.proxy(this.reset, this)
			}])) : this.element.is("div") ? this.isInline = !0 : this._events = [
				[this.element,
				{
					click: b.proxy(this.show, this)
				}]
			];
			for (var h, g, j = 0; j < this._events.length; j++) {
				h = this._events[j][0], g = this._events[j][1], h.on(g)
			}
		},
		_detachEvents: function() {
			for (var h, j, g = 0; g < this._events.length; g++) {
				h = this._events[g][0], j = this._events[g][1], h.off(j)
			}
			this._events = []
		},
		show: function(g) {
			this.picker.show(), this.height = this.component ? this.component.outerHeight() : this.element.outerHeight(), this.forceParse && this.update(), this.place(), b(window).on("resize", b.proxy(this.place, this)), g && (g.stopPropagation(), g.preventDefault()), this.isVisible = !0, this.element.trigger({
				type: "show",
				date: this.date
			})
		},
		hide: function(g) {
			this.isVisible && (this.isInline || (this.picker.hide(), b(window).off("resize", this.place), this.viewMode = this.startViewMode, this.showMode(), this.isInput || b(document).off("mousedown", this.hide), this.forceParse && (this.isInput && this.element.val() || this.hasInput && this.element.find("input").val()) && this.setValue(), this.isVisible = !1, this.element.trigger({
				type: "hide",
				date: this.date
			})))
		},
		remove: function() {
			this._detachEvents(), this.picker.remove(), delete this.picker, delete this.element.data().datetimepicker
		},
		getDate: function() {
			var e = this.getUTCDate();
			return new Date(e.getTime() + 60000 * e.getTimezoneOffset())
		},
		getUTCDate: function() {
			return this.date
		},
		setDate: function(e) {
			this.setUTCDate(new Date(e.getTime() - 60000 * e.getTimezoneOffset()))
		},
		setUTCDate: function(e) {
			e >= this.startDate && e <= this.endDate ? (this.date = e, this.setValue(), this.viewDate = this.date, this.fill()) : this.element.trigger({
				type: "outOfRange",
				date: e,
				startDate: this.startDate,
				endDate: this.endDate
			})
		},
		setFormat: function(g) {
			this.format = d.parseFormat(g, this.formatType);
			var h;
			this.isInput ? h = this.element : this.component && (h = this.element.find("input")), h && h.val() && this.setValue()
		},
		setValue: function() {
			var g = this.getFormattedDate();
			this.isInput ? this.element.val(g) : (this.component && this.element.find("input").val(g), this.element.data("date", g)), this.linkField && b("#" + this.linkField).val(this.getFormattedDate(this.linkFormat))
		},
		getFormattedDate: function(e) {
			return void 0 == e && (e = this.format), d.formatDate(this.date, e, this.language, this.formatType)
		},
		setStartDate: function(e) {
			this.startDate = e || -(1 / 0), this.startDate !== -(1 / 0) && (this.startDate = d.parseDate(this.startDate, this.format, this.language, this.formatType)), this.update(), this.updateNavArrows()
		},
		setEndDate: function(e) {
			this.endDate = e || 1 / 0, this.endDate !== 1 / 0 && (this.endDate = d.parseDate(this.endDate, this.format, this.language, this.formatType)), this.update(), this.updateNavArrows()
		},
		setDaysOfWeekDisabled: function(g) {
			this.daysOfWeekDisabled = g || [], b.isArray(this.daysOfWeekDisabled) || (this.daysOfWeekDisabled = this.daysOfWeekDisabled.split(/,\s*/)), this.daysOfWeekDisabled = b.map(this.daysOfWeekDisabled, function(e) {
				return parseInt(e, 10)
			}), this.update(), this.updateNavArrows()
		},
		place: function() {
			if (!this.isInline) {
				var l = 0;
				b("div").each(function() {
					var e = parseInt(b(this).css("zIndex"), 10);
					e > l && (l = e)
				});
				var h, p, m, j = l + 10;
				this.component ? (h = this.component.offset(), m = h.left, "bottom-left" !== this.pickerPosition && "top-left" !== this.pickerPosition && "auto-left" !== this.pickerPosition || (m += this.component.outerWidth() - this.picker.outerWidth())) : (h = this.element.offset(), m = h.left);
				var g = 0 === this.pickerPosition.indexOf("auto-"),
					k = g ? (h.top + this.picker.outerHeight() > b(window).height() + b(window).scrollTop() ? "top" : "bottom") + (0 === this.pickerPosition.lastIndexOf("-left") ? "-left" : "-right") : this.pickerPosition;
				p = "top-left" === k || "top-right" === k ? h.top - this.picker.outerHeight() : h.top + this.height, this.picker.css({
					top: p,
					left: m,
					zIndex: j
				}).attr("class", "datetimepicker dropdown-menu datetimepicker-dropdown-" + k), this.pickerClass && this.picker.addClass(this.pickerClass), this.pickerId && this.picker.attr("id", this.pickerId)
			}
		},
		update: function() {
			var g, h = !1;
			arguments && arguments.length && ("string" == typeof arguments[0] || arguments[0] instanceof Date) ? (g = arguments[0], h = !0) : (g = this.element.data("date") || (this.isInput ? this.element.val() : this.element.find("input").val()) || this.initialDate, ("string" == typeof g || g instanceof String) && (g = g.replace(/^\s+|\s+$/g, ""))), g || (g = new Date, h = !1), this.date = d.parseDate(g, this.format, this.language, this.formatType), h && this.setValue(), this.date < this.startDate ? this.viewDate = new Date(this.startDate) : this.date > this.endDate ? this.viewDate = new Date(this.endDate) : this.viewDate = new Date(this.date), this.fill()
		},
		fillDow: function() {
			for (var g = this.weekStart, h = "<tr>"; g < this.weekStart + 7;) {
				h += '<th class="dow">' + f[this.language].daysMin[g++ % 7] + "</th>"
			}
			h += "</tr>", this.picker.find(".datetimepicker-days thead").append(h)
		},
		fillMonths: function() {
			for (var g = "", h = 0; h < 12;) {
				g += '<span class="month">' + f[this.language].monthsShort[h++] + "</span>"
			}
			this.picker.find(".datetimepicker-months td").html(g)
		},
		fill: function() {
			if (null != this.date && null != this.viewDate) {
				var X = new Date(this.viewDate),
					N = X.getUTCFullYear(),
					ae = X.getUTCMonth(),
					Q = X.getUTCDate(),
					V = X.getUTCHours(),
					Y = X.getUTCMinutes(),
					ac = this.startDate !== -(1 / 0) ? this.startDate.getUTCFullYear() : -(1 / 0),
					ab = this.startDate !== -(1 / 0) ? this.startDate.getUTCMonth() : -(1 / 0),
					J = this.endDate !== 1 / 0 ? this.endDate.getUTCFullYear() : 1 / 0,
					R = this.endDate !== 1 / 0 ? this.endDate.getUTCMonth() : 1 / 0,
					aa = new c(this.date.getUTCFullYear(), this.date.getUTCMonth(), this.date.getUTCDate()).valueOf(),
					Z = new Date;
				if (this.picker.find(".datetimepicker-days thead th:eq(1)").text(f[this.language].months[ae] + " " + N), "time" == this.formatViewType) {
					var U = V % 12 ? V % 12 : 12,
						H = (U < 10 ? "0" : "") + U,
						B = (Y < 10 ? "0" : "") + Y,
						ad = f[this.language].meridiem[V < 12 ? 0 : 1];
					this.picker.find(".datetimepicker-hours thead th:eq(1)").text(H + ":" + B + " " + ad.toUpperCase()), this.picker.find(".datetimepicker-minutes thead th:eq(1)").text(H + ":" + B + " " + ad.toUpperCase())
				} else {
					this.picker.find(".datetimepicker-hours thead th:eq(1)").text(Q + " " + f[this.language].months[ae] + " " + N), this.picker.find(".datetimepicker-minutes thead th:eq(1)").text(Q + " " + f[this.language].months[ae] + " " + N)
				}
				this.picker.find("tfoot th.today").text(f[this.language].today).toggle(this.todayBtn !== !1), this.updateNavArrows(), this.fillMonths();
				var G = c(N, ae - 1, 28, 0, 0, 0, 0),
					E = d.getDaysInMonth(G.getUTCFullYear(), G.getUTCMonth());
				G.setUTCDate(E), G.setUTCDate(E - (G.getUTCDay() - this.weekStart + 7) % 7);
				var O = new Date(G);
				O.setUTCDate(O.getUTCDate() + 42), O = O.valueOf();
				for (var af, W = []; G.valueOf() < O;) {
					G.getUTCDay() == this.weekStart && W.push("<tr>"), af = "", G.getUTCFullYear() < N || G.getUTCFullYear() == N && G.getUTCMonth() < ae ? af += " old" : (G.getUTCFullYear() > N || G.getUTCFullYear() == N && G.getUTCMonth() > ae) && (af += " new"), this.todayHighlight && G.getUTCFullYear() == Z.getFullYear() && G.getUTCMonth() == Z.getMonth() && G.getUTCDate() == Z.getDate() && (af += " today"), G.valueOf() == aa && (af += " active"), (G.valueOf() + 86400000 <= this.startDate || G.valueOf() > this.endDate || b.inArray(G.getUTCDay(), this.daysOfWeekDisabled) !== -1) && (af += " disabled"), W.push('<td class="day' + af + '">' + G.getUTCDate() + "</td>"), G.getUTCDay() == this.weekEnd && W.push("</tr>"), G.setUTCDate(G.getUTCDate() + 1)
				}
				this.picker.find(".datetimepicker-days tbody").empty().append(W.join("")), W = [];
				for (var e = "", j = "", K = "", o = 0; o < 24; o++) {
					var n = c(N, ae, Q, o);
					af = "", n.valueOf() + 3600000 <= this.startDate || n.valueOf() > this.endDate ? af += " disabled" : V == o && (af += " active"), this.showMeridian && 2 == f[this.language].meridiem.length ? (j = o < 12 ? f[this.language].meridiem[0] : f[this.language].meridiem[1], j != K && ("" != K && W.push("</fieldset>"), W.push('<fieldset class="hour"><legend>' + j.toUpperCase() + "</legend>")), K = j, e = o % 12 ? o % 12 : 12, W.push('<span class="hour' + af + " hour_" + (o < 12 ? "am" : "pm") + '">' + e + "</span>"), 23 == o && W.push("</fieldset>")) : (e = o + ":00", W.push('<span class="hour' + af + '">' + e + "</span>"))
				}
				this.picker.find(".datetimepicker-hours td").html(W.join("")), W = [], e = "", j = "", K = "";
				for (var o = 0; o < 60; o += this.minuteStep) {
					var n = c(N, ae, Q, V, o, 0);
					af = "", n.valueOf() < this.startDate || n.valueOf() > this.endDate ? af += " disabled" : Math.floor(Y / this.minuteStep) == Math.floor(o / this.minuteStep) && (af += " active"), this.showMeridian && 2 == f[this.language].meridiem.length ? (j = V < 12 ? f[this.language].meridiem[0] : f[this.language].meridiem[1], j != K && ("" != K && W.push("</fieldset>"), W.push('<fieldset class="minute"><legend>' + j.toUpperCase() + "</legend>")), K = j, e = V % 12 ? V % 12 : 12, W.push('<span class="minute' + af + '">' + e + ":" + (o < 10 ? "0" + o : o) + "</span>"), 59 == o && W.push("</fieldset>")) : (e = o + ":00", W.push('<span class="minute' + af + '">' + V + ":" + (o < 10 ? "0" + o : o) + "</span>"))
				}
				this.picker.find(".datetimepicker-minutes td").html(W.join(""));
				var I = this.date.getUTCFullYear(),
					q = this.picker.find(".datetimepicker-months").find("th:eq(1)").text(N).end().find("span").removeClass("active");
				I == N && q.eq(this.date.getUTCMonth()).addClass("active"), (N < ac || N > J) && q.addClass("disabled"), N == ac && q.slice(0, ab).addClass("disabled"), N == J && q.slice(R + 1).addClass("disabled"), W = "", N = 10 * parseInt(N / 10, 10);
				var A = this.picker.find(".datetimepicker-years").find("th:eq(1)").text(N + "-" + (N + 9)).end().find("td");
				N -= 1;
				for (var o = -1; o < 11; o++) {
					W += '<span class="year' + (o == -1 || 10 == o ? " old" : "") + (I == N ? " active" : "") + (N < ac || N > J ? " disabled" : "") + '">' + N + "</span>", N += 1
				}
				A.html(W), this.place()
			}
		},
		updateNavArrows: function() {
			var h = new Date(this.viewDate),
				j = h.getUTCFullYear(),
				g = h.getUTCMonth(),
				l = h.getUTCDate(),
				k = h.getUTCHours();
			switch (this.viewMode) {
			case 0:
				this.startDate !== -(1 / 0) && j <= this.startDate.getUTCFullYear() && g <= this.startDate.getUTCMonth() && l <= this.startDate.getUTCDate() && k <= this.startDate.getUTCHours() ? this.picker.find(".prev").css({
					visibility: "hidden"
				}) : this.picker.find(".prev").css({
					visibility: "visible"
				}), this.endDate !== 1 / 0 && j >= this.endDate.getUTCFullYear() && g >= this.endDate.getUTCMonth() && l >= this.endDate.getUTCDate() && k >= this.endDate.getUTCHours() ? this.picker.find(".next").css({
					visibility: "hidden"
				}) : this.picker.find(".next").css({
					visibility: "visible"
				});
				break;
			case 1:
				this.startDate !== -(1 / 0) && j <= this.startDate.getUTCFullYear() && g <= this.startDate.getUTCMonth() && l <= this.startDate.getUTCDate() ? this.picker.find(".prev").css({
					visibility: "hidden"
				}) : this.picker.find(".prev").css({
					visibility: "visible"
				}), this.endDate !== 1 / 0 && j >= this.endDate.getUTCFullYear() && g >= this.endDate.getUTCMonth() && l >= this.endDate.getUTCDate() ? this.picker.find(".next").css({
					visibility: "hidden"
				}) : this.picker.find(".next").css({
					visibility: "visible"
				});
				break;
			case 2:
				this.startDate !== -(1 / 0) && j <= this.startDate.getUTCFullYear() && g <= this.startDate.getUTCMonth() ? this.picker.find(".prev").css({
					visibility: "hidden"
				}) : this.picker.find(".prev").css({
					visibility: "visible"
				}), this.endDate !== 1 / 0 && j >= this.endDate.getUTCFullYear() && g >= this.endDate.getUTCMonth() ? this.picker.find(".next").css({
					visibility: "hidden"
				}) : this.picker.find(".next").css({
					visibility: "visible"
				});
				break;
			case 3:
			case 4:
				this.startDate !== -(1 / 0) && j <= this.startDate.getUTCFullYear() ? this.picker.find(".prev").css({
					visibility: "hidden"
				}) : this.picker.find(".prev").css({
					visibility: "visible"
				}), this.endDate !== 1 / 0 && j >= this.endDate.getUTCFullYear() ? this.picker.find(".next").css({
					visibility: "hidden"
				}) : this.picker.find(".next").css({
					visibility: "visible"
				})
			}
		},
		mousewheel: function(h) {
			if (h.preventDefault(), h.stopPropagation(), !this.wheelPause) {
				this.wheelPause = !0;
				var g = h.originalEvent,
					k = g.wheelDelta,
					j = k > 0 ? 1 : 0 === k ? 0 : -1;
				this.wheelViewModeNavigationInverseDirection && (j = -j), this.showMode(j), setTimeout(b.proxy(function() {
					this.wheelPause = !1
				}, this), this.wheelViewModeNavigationDelay)
			}
		},
		click: function(m) {
			m.stopPropagation(), m.preventDefault();
			var j = b(m.target).closest("span, td, th, legend");
			if (1 == j.length) {
				if (j.is(".disabled")) {
					return void this.element.trigger({
						type: "outOfRange",
						date: this.viewDate,
						startDate: this.startDate,
						endDate: this.endDate
					})
				}
				switch (j[0].nodeName.toLowerCase()) {
				case "th":
					switch (j[0].className) {
					case "switch":
						this.showMode(1);
						break;
					case "prev":
					case "next":
						var y = d.modes[this.viewMode].navStep * ("prev" == j[0].className ? -1 : 1);
						switch (this.viewMode) {
						case 0:
							this.viewDate = this.moveHour(this.viewDate, y);
							break;
						case 1:
							this.viewDate = this.moveDate(this.viewDate, y);
							break;
						case 2:
							this.viewDate = this.moveMonth(this.viewDate, y);
							break;
						case 3:
						case 4:
							this.viewDate = this.moveYear(this.viewDate, y)
						}
						this.fill();
						break;
					case "today":
						var w = new Date;
						w = c(w.getFullYear(), w.getMonth(), w.getDate(), w.getHours(), w.getMinutes(), w.getSeconds(), 0), w < this.startDate ? w = this.startDate : w > this.endDate && (w = this.endDate), this.viewMode = this.startViewMode, this.showMode(0), this._setDate(w), this.fill(), this.autoclose && this.hide()
					}
					break;
				case "span":
					if (!j.is(".disabled")) {
						var e = this.viewDate.getUTCFullYear(),
							k = this.viewDate.getUTCMonth(),
							o = this.viewDate.getUTCDate(),
							v = this.viewDate.getUTCHours(),
							q = this.viewDate.getUTCMinutes(),
							x = this.viewDate.getUTCSeconds();
						if (j.is(".month") ? (this.viewDate.setUTCDate(1), k = j.parent().find("span").index(j), o = this.viewDate.getUTCDate(), this.viewDate.setUTCMonth(k), this.element.trigger({
							type: "changeMonth",
							date: this.viewDate
						}), this.viewSelect >= 3 && this._setDate(c(e, k, o, v, q, x, 0))) : j.is(".year") ? (this.viewDate.setUTCDate(1), e = parseInt(j.text(), 10) || 0, this.viewDate.setUTCFullYear(e), this.element.trigger({
							type: "changeYear",
							date: this.viewDate
						}), this.viewSelect >= 4 && this._setDate(c(e, k, o, v, q, x, 0))) : j.is(".hour") ? (v = parseInt(j.text(), 10) || 0, (j.hasClass("hour_am") || j.hasClass("hour_pm")) && (12 == v && j.hasClass("hour_am") ? v = 0 : 12 != v && j.hasClass("hour_pm") && (v += 12)), this.viewDate.setUTCHours(v), this.element.trigger({
							type: "changeHour",
							date: this.viewDate
						}), this.viewSelect >= 1 && this._setDate(c(e, k, o, v, q, x, 0))) : j.is(".minute") && (q = parseInt(j.text().substr(j.text().indexOf(":") + 1), 10) || 0, this.viewDate.setUTCMinutes(q), this.element.trigger({
							type: "changeMinute",
							date: this.viewDate
						}), this.viewSelect >= 0 && this._setDate(c(e, k, o, v, q, x, 0))), 0 != this.viewMode) {
							var g = this.viewMode;
							this.showMode(-1), this.fill(), g == this.viewMode && this.autoclose && this.hide()
						} else {
							this.fill(), this.autoclose && this.hide()
						}
					}
					break;
				case "td":
					if (j.is(".day") && !j.is(".disabled")) {
						var o = parseInt(j.text(), 10) || 1,
							e = this.viewDate.getUTCFullYear(),
							k = this.viewDate.getUTCMonth(),
							v = this.viewDate.getUTCHours(),
							q = this.viewDate.getUTCMinutes(),
							x = this.viewDate.getUTCSeconds();
						j.is(".old") ? 0 === k ? (k = 11, e -= 1) : k -= 1 : j.is(".new") && (11 == k ? (k = 0, e += 1) : k += 1), this.viewDate.setUTCFullYear(e), this.viewDate.setUTCMonth(k, o), this.element.trigger({
							type: "changeDay",
							date: this.viewDate
						}), this.viewSelect >= 2 && this._setDate(c(e, k, o, v, q, x, 0))
					}
					var g = this.viewMode;
					this.showMode(-1), this.fill(), g == this.viewMode && this.autoclose && this.hide()
				}
			}
		},
		_setDate: function(h, j) {
			j && "date" != j || (this.date = h), j && "view" != j || (this.viewDate = h), this.fill(), this.setValue();
			var g;
			this.isInput ? g = this.element : this.component && (g = this.element.find("input")), g && (g.change(), this.autoclose && (!j || "date" == j)), this.element.trigger({
				type: "changeDate",
				date: this.date
			}), null === h && (this.date = this.viewDate)
		},
		moveMinute: function(h, j) {
			if (!j) {
				return h
			}
			var g = new Date(h.valueOf());
			return g.setUTCMinutes(g.getUTCMinutes() + j * this.minuteStep), g
		},
		moveHour: function(h, j) {
			if (!j) {
				return h
			}
			var g = new Date(h.valueOf());
			return g.setUTCHours(g.getUTCHours() + j), g
		},
		moveDate: function(h, j) {
			if (!j) {
				return h
			}
			var g = new Date(h.valueOf());
			return g.setUTCDate(g.getUTCDate() + j), g
		},
		moveMonth: function(u, p) {
			if (!p) {
				return u
			}
			var m, j, h = new Date(u.valueOf()),
				v = h.getUTCDate(),
				q = h.getUTCMonth(),
				g = Math.abs(p);
			if (p = p > 0 ? 1 : -1, 1 == g) {
				j = p == -1 ?
				function() {
					return h.getUTCMonth() == q
				} : function() {
					return h.getUTCMonth() != m
				}, m = q + p, h.setUTCMonth(m), (m < 0 || m > 11) && (m = (m + 12) % 12)
			} else {
				for (var k = 0; k < g; k++) {
					h = this.moveMonth(h, p)
				}
				m = h.getUTCMonth(), h.setUTCDate(v), j = function() {
					return m != h.getUTCMonth()
				}
			}
			for (; j();) {
				h.setUTCDate(--v), h.setUTCMonth(m)
			}
			return h
		},
		moveYear: function(g, h) {
			return this.moveMonth(g, 12 * h)
		},
		dateWithinRange: function(e) {
			return e >= this.startDate && e <= this.endDate
		},
		keydown: function(j) {
			if (this.picker.is(":not(:visible)")) {
				return void(27 == j.keyCode && this.show())
			}
			var l, h, p, m = !1;
			switch (j.keyCode) {
			case 27:
				this.hide(), j.preventDefault();
				break;
			case 37:
			case 39:
				if (!this.keyboardNavigation) {
					break
				}
				l = 37 == j.keyCode ? -1 : 1, viewMode = this.viewMode, j.ctrlKey ? viewMode += 2 : j.shiftKey && (viewMode += 1), 4 == viewMode ? (h = this.moveYear(this.date, l), p = this.moveYear(this.viewDate, l)) : 3 == viewMode ? (h = this.moveMonth(this.date, l), p = this.moveMonth(this.viewDate, l)) : 2 == viewMode ? (h = this.moveDate(this.date, l), p = this.moveDate(this.viewDate, l)) : 1 == viewMode ? (h = this.moveHour(this.date, l), p = this.moveHour(this.viewDate, l)) : 0 == viewMode && (h = this.moveMinute(this.date, l), p = this.moveMinute(this.viewDate, l)), this.dateWithinRange(h) && (this.date = h, this.viewDate = p, this.setValue(), this.update(), j.preventDefault(), m = !0);
				break;
			case 38:
			case 40:
				if (!this.keyboardNavigation) {
					break
				}
				l = 38 == j.keyCode ? -1 : 1, viewMode = this.viewMode, j.ctrlKey ? viewMode += 2 : j.shiftKey && (viewMode += 1), 4 == viewMode ? (h = this.moveYear(this.date, l), p = this.moveYear(this.viewDate, l)) : 3 == viewMode ? (h = this.moveMonth(this.date, l), p = this.moveMonth(this.viewDate, l)) : 2 == viewMode ? (h = this.moveDate(this.date, 7 * l), p = this.moveDate(this.viewDate, 7 * l)) : 1 == viewMode ? this.showMeridian ? (h = this.moveHour(this.date, 6 * l), p = this.moveHour(this.viewDate, 6 * l)) : (h = this.moveHour(this.date, 4 * l), p = this.moveHour(this.viewDate, 4 * l)) : 0 == viewMode && (h = this.moveMinute(this.date, 4 * l), p = this.moveMinute(this.viewDate, 4 * l)), this.dateWithinRange(h) && (this.date = h, this.viewDate = p, this.setValue(), this.update(), j.preventDefault(), m = !0);
				break;
			case 13:
				if (0 != this.viewMode) {
					var k = this.viewMode;
					this.showMode(-1), this.fill(), k == this.viewMode && this.autoclose && this.hide()
				} else {
					this.fill(), this.autoclose && this.hide()
				}
				j.preventDefault();
				break;
			case 9:
				this.hide()
			}
			if (m) {
				var g;
				this.isInput ? g = this.element : this.component && (g = this.element.find("input")), g && g.change(), this.element.trigger({
					type: "changeDate",
					date: this.date
				})
			}
		},
		showMode: function(g) {
			if (g) {
				var h = Math.max(0, Math.min(d.modes.length - 1, this.viewMode + g));
				h >= this.minView && h <= this.maxView && (this.element.trigger({
					type: "changeMode",
					date: this.viewDate,
					oldViewMode: this.viewMode,
					newViewMode: h
				}), this.viewMode = h)
			}
			this.picker.find(">div").hide().filter(".datetimepicker-" + d.modes[this.viewMode].clsName).css("display", "block"), this.updateNavArrows()
		},
		reset: function(e) {
			this._setDate(null, "date")
		}
	}, b.fn.datetimepicker = function(g) {
		var h = Array.apply(null, arguments);
		return h.shift(), this.each(function() {
			var k = b(this),
				j = k.data("datetimepicker"),
				e = "object" == typeof g && g;
			j || k.data("datetimepicker", j = new a(this, b.extend({}, b.fn.datetimepicker.defaults, k.data(), e))), "string" == typeof g && "function" == typeof j[g] && j[g].apply(j, h)
		})
	}, b.fn.datetimepicker.defaults = {
		pickerPosition: "auto-right"
	}, b.fn.datetimepicker.Constructor = a;
	var f = b.fn.datetimepicker.dates = {
		en: {
			days: ["Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"],
			daysShort: ["Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"],
			daysMin: ["Su", "Mo", "Tu", "We", "Th", "Fr", "Sa", "Su"],
			months: ["January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"],
			monthsShort: ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"],
			meridiem: ["am", "pm"],
			suffix: ["st", "nd", "rd", "th"],
			today: "Today"
		}
	};
	f["zh-cn"] = {
		days: ["星期日", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六", "星期日"],
		daysShort: ["周日", "周一", "周二", "周三", "周四", "周五", "周六", "周日"],
		daysMin: ["日", "一", "二", "三", "四", "五", "六", "日"],
		months: ["一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"],
		monthsShort: ["一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"],
		today: "今日",
		suffix: [],
		meridiem: []
	}, f["zh-tw"] = {
		days: ["星期日", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六", "星期日"],
		daysShort: ["周日", "周一", "周二", "周三", "周四", "周五", "周六", "周日"],
		daysMin: ["日", "一", "二", "三", "四", "五", "六", "日"],
		months: ["一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"],
		monthsShort: ["一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"],
		today: "今天",
		suffix: [],
		meridiem: ["上午", "下午"]
	};
	var d = {
		modes: [{
			clsName: "minutes",
			navFnc: "Hours",
			navStep: 1
		}, {
			clsName: "hours",
			navFnc: "Date",
			navStep: 1
		}, {
			clsName: "days",
			navFnc: "Month",
			navStep: 1
		}, {
			clsName: "months",
			navFnc: "FullYear",
			navStep: 1
		}, {
			clsName: "years",
			navFnc: "FullYear",
			navStep: 10
		}],
		isLeapYear: function(e) {
			return e % 4 === 0 && e % 100 !== 0 || e % 400 === 0
		},
		getDaysInMonth: function(g, h) {
			return [31, d.isLeapYear(g) ? 29 : 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31][h]
		},
		getDefaultFormat: function(g, h) {
			if ("standard" == g) {
				return "input" == h ? "yyyy-mm-dd hh:ii" : "yyyy-mm-dd hh:ii:ss"
			}
			if ("php" == g) {
				return "input" == h ? "Y-m-d H:i" : "Y-m-d H:i:s"
			}
			throw new Error("Invalid format type.")
		},
		validParts: function(e) {
			if ("standard" == e) {
				return /hh?|HH?|p|P|ii?|ss?|dd?|DD?|mm?|MM?|yy(?:yy)?/g
			}
			if ("php" == e) {
				return /[dDjlNwzFmMnStyYaABgGhHis]/g
			}
			throw new Error("Invalid format type.")
		},
		nonpunctuation: /[^ -\/:-@\[-`{-~\t\n\rTZ]+/g,
		parseFormat: function(h, j) {
			var g = h.replace(this.validParts(j), "\0").split("\0"),
				k = h.match(this.validParts(j));
			if (!g || !g.length || !k || 0 == k.length) {
				throw new Error("Invalid date format.")
			}
			return {
				separators: g,
				parts: k
			}
		},
		parseDate: function(k, J, F, e) {
			if (k instanceof Date) {
				var q = new Date(k.valueOf() - 60000 * k.getTimezoneOffset());
				return q.setMilliseconds(0), q
			}
			if (/^\d{4}\-\d{1,2}\-\d{1,2}$/.test(k) && (J = this.parseFormat("yyyy-mm-dd", e)), /^\d{4}\-\d{1,2}\-\d{1,2}[T ]\d{1,2}\:\d{1,2}$/.test(k) && (J = this.parseFormat("yyyy-mm-dd hh:ii", e)), /^\d{4}\-\d{1,2}\-\d{1,2}[T ]\d{1,2}\:\d{1,2}\:\d{1,2}[Z]{0,1}$/.test(k) && (J = this.parseFormat("yyyy-mm-dd hh:ii:ss", e)), /^[-+]\d+[dmwy]([\s,]+[-+]\d+[dmwy])*$/.test(k)) {
				var x, C, B = /([-+]\d+)([dmwy])/,
					I = k.match(/([-+]\d+)([dmwy])/g);
				k = new Date;
				for (var j = 0; j < I.length; j++) {
					switch (x = B.exec(I[j]), C = parseInt(x[1]), x[2]) {
					case "d":
						k.setUTCDate(k.getUTCDate() + C);
						break;
					case "m":
						k = a.prototype.moveMonth.call(a.prototype, k, C);
						break;
					case "w":
						k.setUTCDate(k.getUTCDate() + 7 * C);
						break;
					case "y":
						k = a.prototype.moveYear.call(a.prototype, k, C)
					}
				}
				return c(k.getUTCFullYear(), k.getUTCMonth(), k.getUTCDate(), k.getUTCHours(), k.getUTCMinutes(), k.getUTCSeconds(), 0)
			}
			var A, z, x, I = k && k.match(this.nonpunctuation) || [],
				k = new Date(0, 0, 0, 0, 0, 0, 0),
				n = {},
				H = ["hh", "h", "ii", "i", "ss", "s", "yyyy", "yy", "M", "MM", "m", "mm", "D", "DD", "d", "dd", "H", "HH", "p", "P"],
				E = {
					hh: function(g, h) {
						return g.setUTCHours(h)
					},
					h: function(g, h) {
						return g.setUTCHours(h)
					},
					HH: function(g, h) {
						return g.setUTCHours(12 == h ? 0 : h)
					},
					H: function(g, h) {
						return g.setUTCHours(12 == h ? 0 : h)
					},
					ii: function(g, h) {
						return g.setUTCMinutes(h)
					},
					i: function(g, h) {
						return g.setUTCMinutes(h)
					},
					ss: function(g, h) {
						return g.setUTCSeconds(h)
					},
					s: function(g, h) {
						return g.setUTCSeconds(h)
					},
					yyyy: function(g, h) {
						return g.setUTCFullYear(h)
					},
					yy: function(g, h) {
						return g.setUTCFullYear(2000 + h)
					},
					m: function(g, h) {
						for (h -= 1; h < 0;) {
							h += 12
						}
						for (h %= 12, g.setUTCMonth(h); g.getUTCMonth() != h;) {
							g.setUTCDate(g.getUTCDate() - 1)
						}
						return g
					},
					d: function(g, h) {
						return g.setUTCDate(h)
					},
					p: function(g, h) {
						return g.setUTCHours(1 == h ? g.getUTCHours() + 12 : g.getUTCHours())
					}
				};
			if (E.M = E.MM = E.mm = E.m, E.dd = E.d, E.P = E.p, k = c(k.getFullYear(), k.getMonth(), k.getDate(), k.getHours(), k.getMinutes(), k.getSeconds()), I.length == J.parts.length) {
				for (var j = 0, D = J.parts.length; j < D; j++) {
					if (A = parseInt(I[j], 10), x = J.parts[j], isNaN(A)) {
						switch (x) {
						case "MM":
							z = b(f[F].months).filter(function() {
								var g = this.slice(0, I[j].length),
									h = I[j].slice(0, g.length);
								return g == h
							}), A = b.inArray(z[0], f[F].months) + 1;
							break;
						case "M":
							z = b(f[F].monthsShort).filter(function() {
								var g = this.slice(0, I[j].length),
									h = I[j].slice(0, g.length);
								return g == h
							}), A = b.inArray(z[0], f[F].monthsShort) + 1;
							break;
						case "p":
						case "P":
							A = b.inArray(I[j].toLowerCase(), f[F].meridiem)
						}
					}
					n[x] = A
				}
				for (var G, j = 0; j < H.length; j++) {
					G = H[j], G in n && !isNaN(n[G]) && E[G](k, n[G])
				}
			}
			return k
		},
		formatDate: function(p, k, n, j) {
			if (null == p) {
				return ""
			}
			var o;
			if ("standard" == j) {
				o = {
					yy: p.getUTCFullYear().toString().substring(2),
					yyyy: p.getUTCFullYear(),
					m: p.getUTCMonth() + 1,
					M: f[n].monthsShort[p.getUTCMonth()],
					MM: f[n].months[p.getUTCMonth()],
					d: p.getUTCDate(),
					D: f[n].daysShort[p.getUTCDay()],
					DD: f[n].days[p.getUTCDay()],
					p: 2 == f[n].meridiem.length ? f[n].meridiem[p.getUTCHours() < 12 ? 0 : 1] : "",
					h: p.getUTCHours(),
					i: p.getUTCMinutes(),
					s: p.getUTCSeconds()
				}, 2 == f[n].meridiem.length ? o.H = o.h % 12 == 0 ? 12 : o.h % 12 : o.H = o.h, o.HH = (o.H < 10 ? "0" : "") + o.H, o.P = o.p.toUpperCase(), o.hh = (o.h < 10 ? "0" : "") + o.h, o.ii = (o.i < 10 ? "0" : "") + o.i, o.ss = (o.s < 10 ? "0" : "") + o.s, o.dd = (o.d < 10 ? "0" : "") + o.d, o.mm = (o.m < 10 ? "0" : "") + o.m
			} else {
				if ("php" != j) {
					throw new Error("Invalid format type.")
				}
				o = {
					y: p.getUTCFullYear().toString().substring(2),
					Y: p.getUTCFullYear(),
					F: f[n].months[p.getUTCMonth()],
					M: f[n].monthsShort[p.getUTCMonth()],
					n: p.getUTCMonth() + 1,
					t: d.getDaysInMonth(p.getUTCFullYear(), p.getUTCMonth()),
					j: p.getUTCDate(),
					l: f[n].days[p.getUTCDay()],
					D: f[n].daysShort[p.getUTCDay()],
					w: p.getUTCDay(),
					N: 0 == p.getUTCDay() ? 7 : p.getUTCDay(),
					S: p.getUTCDate() % 10 <= f[n].suffix.length ? f[n].suffix[p.getUTCDate() % 10 - 1] : "",
					a: 2 == f[n].meridiem.length ? f[n].meridiem[p.getUTCHours() < 12 ? 0 : 1] : "",
					g: p.getUTCHours() % 12 == 0 ? 12 : p.getUTCHours() % 12,
					G: p.getUTCHours(),
					i: p.getUTCMinutes(),
					s: p.getUTCSeconds()
				}, o.m = (o.n < 10 ? "0" : "") + o.n, o.d = (o.j < 10 ? "0" : "") + o.j, o.A = o.a.toString().toUpperCase(), o.h = (o.g < 10 ? "0" : "") + o.g, o.H = (o.G < 10 ? "0" : "") + o.G, o.i = (o.i < 10 ? "0" : "") + o.i, o.s = (o.s < 10 ? "0" : "") + o.s
			}
			for (var p = [], g = b.extend([], k.separators), m = 0, q = k.parts.length; m < q; m++) {
				g.length && p.push(g.shift()), p.push(o[k.parts[m]])
			}
			return g.length && p.push(g.shift()), p.join("")
		},
		convertViewMode: function(e) {
			switch (e) {
			case 4:
			case "decade":
				e = 4;
				break;
			case 3:
			case "year":
				e = 3;
				break;
			case 2:
			case "month":
				e = 2;
				break;
			case 1:
			case "day":
				e = 1;
				break;
			case 0:
			case "hour":
				e = 0
			}
			return e
		},
		headTemplate: '<thead><tr><th class="prev"><i class="icon-arrow-left"/></th><th colspan="5" class="switch"></th><th class="next"><i class="icon-arrow-right"/></th></tr></thead>',
		contTemplate: '<tbody><tr><td colspan="7"></td></tr></tbody>',
		footTemplate: '<tfoot><tr><th colspan="7" class="today"></th></tr></tfoot>'
	};
	d.template = '<div class="datetimepicker"><div class="datetimepicker-minutes"><table class=" table-condensed">' + d.headTemplate + d.contTemplate + d.footTemplate + '</table></div><div class="datetimepicker-hours"><table class=" table-condensed">' + d.headTemplate + d.contTemplate + d.footTemplate + '</table></div><div class="datetimepicker-days"><table class=" table-condensed">' + d.headTemplate + "<tbody></tbody>" + d.footTemplate + '</table></div><div class="datetimepicker-months"><table class="table-condensed">' + d.headTemplate + d.contTemplate + d.footTemplate + '</table></div><div class="datetimepicker-years"><table class="table-condensed">' + d.headTemplate + d.contTemplate + d.footTemplate + "</table></div></div>", b.fn.datetimepicker.DPGlobal = d, b.fn.datetimepicker.noConflict = function() {
		return b.fn.datetimepicker = old, this
	}, b(document).on("focus.datetimepicker.data-api click.datetimepicker.data-api", '[data-provide="datetimepicker"]', function(h) {
		var g = b(this);
		g.data("datetimepicker") || (h.preventDefault(), g.datetimepicker("show"))
	}), b(function() {
		b('[data-provide="datetimepicker-inline"]').datetimepicker()
	})
}(window.jQuery), /*! bootbox.js v4.4.0 http://bootboxjs.com/license.txt */

function(a, b) {
	"function" == typeof define && define.amd ? define(["jquery"], b) : "object" == typeof exports ? module.exports = b(require("jquery")) : a.bootbox = b(a.jQuery)
}(this, function t(B, x) {
	function q(a) {
		var c = v[A.locale];
		return c ? c[a] : v.en[a]
	}
	function k(c, a, e) {
		c.stopPropagation(), c.preventDefault();
		var d = B.isFunction(e) && e.call(a, c) === !1;
		d || a.modal("hide")
	}
	function G(c) {
		var d, a = 0;
		for (d in c) {
			a++
		}
		return a
	}
	function E(c, a) {
		var d = 0;
		B.each(c, function(f, g) {
			a(f, g, d++)
		})
	}
	function b(c) {
		var a, d;
		if ("object" != typeof c) {
			throw new Error("Please supply an object of options")
		}
		if (!c.message) {
			throw new Error("Please specify a message")
		}
		return c = B.extend({}, A, c), c.buttons || (c.buttons = {}), a = c.buttons, d = G(a), E(a, function(e, g, f) {
			if (B.isFunction(g) && (g = a[e] = {
				callback: g
			}), "object" !== B.type(g)) {
				throw new Error("button with key " + e + " must be an object")
			}
			g.label || (g.label = e), g.className || (2 === d && ("ok" === e || "confirm" === e) || 1 === d ? g.className = "btn-primary" : g.className = "btn-default")
		}), c
	}
	function w(c, d) {
		var a = c.length,
			f = {};
		if (a < 1 || a > 2) {
			throw new Error("Invalid argument length")
		}
		return 2 === a || "string" == typeof c[0] ? (f[d[0]] = c[0], f[d[1]] = c[1]) : f = c[0], f
	}
	function y(c, a, d) {
		return B.extend(!0, {}, c, w(a, d))
	}
	function D(c, d, a, g) {
		var f = {
			className: "bootbox-" + c,
			buttons: C.apply(null, d)
		};
		return F(y(f, g, a), d)
	}
	function C() {
		for (var f = {}, h = 0, d = arguments.length; h < d; h++) {
			var l = arguments[h],
				g = l.toLowerCase(),
				c = l.toUpperCase();
			f[g] = {
				label: q(c)
			}
		}
		return f
	}
	function F(a, c) {
		var d = {};
		return E(c, function(f, g) {
			d[g] = !0
		}), E(a.buttons, function(e) {
			if (d[e] === x) {
				throw new Error("button key " + e + " is not allowed (options are " + c.join("\n") + ")")
			}
		}), a
	}
	var j = {
		dialog: "<div class='bootbox modal' tabindex='-1' role='dialog'><div class='modal-dialog'><div class='modal-content'><div class='modal-body'><div class='bootbox-body'></div></div></div></div></div>",
		header: "<div class='modal-header'><h4 class='modal-title'></h4></div>",
		footer: "<div class='modal-footer'></div>",
		closeButton: "<button type='button' class='bootbox-close-button close' data-dismiss='modal' aria-hidden='true'>&times;</button>",
		form: "<form class='bootbox-form'></form>",
		inputs: {
			text: "<input class='bootbox-input bootbox-input-text form-control' autocomplete=off type=text />",
			textarea: "<textarea class='bootbox-input bootbox-input-textarea form-control'></textarea>",
			email: "<input class='bootbox-input bootbox-input-email form-control' autocomplete='off' type='email' />",
			select: "<select class='bootbox-input bootbox-input-select form-control'></select>",
			checkbox: "<div class='checkbox'><label><input class='bootbox-input bootbox-input-checkbox' type='checkbox' /></label></div>",
			date: "<input class='bootbox-input bootbox-input-date form-control' autocomplete=off type='date' />",
			time: "<input class='bootbox-input bootbox-input-time form-control' autocomplete=off type='time' />",
			number: "<input class='bootbox-input bootbox-input-number form-control' autocomplete=off type='number' />",
			password: "<input class='bootbox-input bootbox-input-password form-control' autocomplete='off' type='password' />"
		}
	},
		A = {
			locale: B.zui && B.zui.clientLang ? B.zui.clientLang() : "zh_cn",
			backdrop: "static",
			animate: !0,
			className: null,
			closeButton: !0,
			show: !0,
			container: "body"
		},
		z = {};
	z.alert = function() {
		var a;
		if (a = D("alert", ["ok"], ["message", "callback"], arguments), a.callback && !B.isFunction(a.callback)) {
			throw new Error("alert requires callback property to be a function when provided")
		}
		return a.buttons.ok.callback = a.onEscape = function() {
			return !B.isFunction(a.callback) || a.callback.call(this)
		}, z.dialog(a)
	}, z.confirm = function() {
		var a;
		if (a = D("confirm", ["confirm", "cancel"], ["message", "callback"], arguments), a.buttons.cancel.callback = a.onEscape = function() {
			return a.callback.call(this, !1)
		}, a.buttons.confirm.callback = function() {
			return a.callback.call(this, !0)
		}, !B.isFunction(a.callback)) {
			throw new Error("confirm requires a callback")
		}
		return z.dialog(a)
	}, z.prompt = function() {
		var H, e, d, I, a, h, u;
		if (I = B(j.form), e = {
			className: "bootbox-prompt",
			buttons: C("cancel", "confirm"),
			value: "",
			inputType: "text"
		}, H = F(y(e, arguments, ["title", "callback"]), ["confirm", "cancel"]), h = H.show === x || H.show, H.message = I, H.buttons.cancel.callback = H.onEscape = function() {
			return H.callback.call(this, null)
		}, H.buttons.confirm.callback = function() {
			var c;
			switch (H.inputType) {
			case "text":
			case "textarea":
			case "email":
			case "select":
			case "date":
			case "time":
			case "number":
			case "password":
				c = a.val();
				break;
			case "checkbox":
				var f = a.find("input:checked");
				c = [], E(f, function(l, m) {
					c.push(B(m).val())
				})
			}
			return H.callback.call(this, c)
		}, H.show = !1, !H.title) {
			throw new Error("prompt requires a title")
		}
		if (!B.isFunction(H.callback)) {
			throw new Error("prompt requires a callback")
		}
		if (!j.inputs[H.inputType]) {
			throw new Error("invalid prompt type")
		}
		switch (a = B(j.inputs[H.inputType]), H.inputType) {
		case "text":
		case "textarea":
		case "email":
		case "date":
		case "time":
		case "number":
		case "password":
			a.val(H.value);
			break;
		case "select":
			var p = {};
			if (u = H.inputOptions || [], !B.isArray(u)) {
				throw new Error("Please pass an array of input options")
			}
			if (!u.length) {
				throw new Error("prompt with select requires options")
			}
			E(u, function(c, l) {
				var f = a;
				if (l.value === x || l.text === x) {
					throw new Error("given options in wrong format")
				}
				l.group && (p[l.group] || (p[l.group] = B("<optgroup/>").attr("label", l.group)), f = p[l.group]), f.append("<option value='" + l.value + "'>" + l.text + "</option>")
			}), E(p, function(c, f) {
				a.append(f)
			}), a.val(H.value);
			break;
		case "checkbox":
			var g = B.isArray(H.value) ? H.value : [H.value];
			if (u = H.inputOptions || [], !u.length) {
				throw new Error("prompt with checkbox requires options")
			}
			if (!u[0].value || !u[0].text) {
				throw new Error("given options in wrong format")
			}
			a = B("<div/>"), E(u, function(c, l) {
				var f = B(j.inputs[H.inputType]);
				f.find("input").attr("value", l.value), f.find("label").append(l.text), E(g, function(m, n) {
					n === l.value && f.find("input").prop("checked", !0)
				}), a.append(f)
			})
		}
		return H.placeholder && a.attr("placeholder", H.placeholder), H.pattern && a.attr("pattern", H.pattern), H.maxlength && a.attr("maxlength", H.maxlength), I.append(a), I.on("submit", function(c) {
			c.preventDefault(), c.stopPropagation(), d.find(".btn-primary").click()
		}), d = z.dialog(H), d.off("shown.zui.modal"), d.on("shown.zui.modal", function() {
			a.focus()
		}), h === !0 && d.modal("show"), d
	}, z.dialog = function(f) {
		f = b(f);
		var r = B(j.dialog),
			m = r.find(".modal-dialog"),
			a = r.find(".modal-body"),
			g = f.buttons,
			p = "",
			o = {
				onEscape: f.onEscape
			};
		if (B.fn.modal === x) {
			throw new Error("$.fn.modal is not defined; please double check you have included the Bootstrap JavaScript library. See http://getbootstrap.com/javascript/ for more details.")
		}
		if (E(g, function(c, d) {
			p += "<button data-bb-handler='" + c + "' type='button' class='btn " + d.className + "'>" + d.label + "</button>", o[c] = d.callback
		}), a.find(".bootbox-body").html(f.message), f.animate === !0 && r.addClass("fade"), f.className && r.addClass(f.className), "large" === f.size ? m.addClass("modal-lg") : "small" === f.size && m.addClass("modal-sm"), f.title && a.before(j.header), f.closeButton) {
			var e = B(j.closeButton);
			f.title ? r.find(".modal-header").prepend(e) : e.css("margin-top", "-10px").prependTo(a)
		}
		return f.title && r.find(".modal-title").html(f.title), p.length && (a.after(j.footer), r.find(".modal-footer").html(p)), r.on("hidden.zui.modal", function(c) {
			c.target === this && r.remove()
		}), r.on("shown.zui.modal", function() {
			r.find(".btn-primary:first").focus()
		}), "static" !== f.backdrop && r.on("click.dismiss.zui.modal", function(c) {
			r.children(".modal-backdrop").length && (c.currentTarget = r.children(".modal-backdrop").get(0)), c.target === c.currentTarget && r.trigger("escape.close.bb")
		}), r.on("escape.close.bb", function(c) {
			o.onEscape && k(c, r, o.onEscape)
		}), r.on("click", ".modal-footer button", function(d) {
			var c = B(this).data("bb-handler");
			k(d, r, o[c])
		}), r.on("click", ".bootbox-close-button", function(c) {
			k(c, r, o.onEscape)
		}), r.on("keyup", function(c) {
			27 === c.which && r.trigger("escape.close.bb")
		}), B(f.container).append(r), r.modal({
			backdrop: !! f.backdrop && "static",
			keyboard: !1,
			show: !1
		}), f.show && r.modal("show"), r
	}, z.setDefaults = function() {
		var a = {};
		2 === arguments.length ? a[arguments[0]] = arguments[1] : a = arguments[0], B.extend(A, a)
	}, z.hideAll = function() {
		return B(".bootbox").modal("hide"), z
	};
	var v = {
		en: {
			OK: "OK",
			CANCEL: "Cancel",
			CONFIRM: "OK"
		},
		zh_cn: {
			OK: "确认",
			CANCEL: "取消",
			CONFIRM: "确认"
		},
		zh_tw: {
			OK: "確認",
			CANCEL: "取消",
			CONFIRM: "確認"
		}
	};
	return z.addLocale = function(c, a) {
		return B.each(["OK", "CANCEL", "CONFIRM"], function(d, f) {
			if (!a[f]) {
				throw new Error("Please supply a translation for '" + f + "'")
			}
		}), v[c] = {
			OK: a.OK,
			CANCEL: a.CANCEL,
			CONFIRM: a.CONFIRM
		}, z
	}, z.removeLocale = function(a) {
		return delete v[a], z
	}, z.setLocale = function(a) {
		return z.setDefaults("locale", a)
	}, z.init = function(a) {
		return t(a || B)
	}, z
}),
/*!
Chosen, a Select Box Enhancer for jQuery and Prototype
by Patrick Filler for Harvest, http://getharvest.com

Version 1.1.0
Full source at https://github.com/harvesthq/chosen
Copyright (c) 2011 Harvest http://getharvest.com

MIT License, https://github.com/harvesthq/chosen/blob/master/LICENSE.md
*/
function() {
	var d, h, c, k, j, f = {}.hasOwnProperty,
		b = function(l, m) {
			function a() {
				this.constructor = l
			}
			for (var o in m) {
				f.call(m, o) && (l[o] = m[o])
			}
			return a.prototype = m.prototype, l.prototype = new a, l.__super__ = m.prototype, l
		},
		g = {
			zh_cn: {
				no_results_text: "没有找到"
			},
			zh_tw: {
				no_results_text: "沒有找到"
			},
			en: {
				no_results_text: "No results match"
			}
		};
	k = function() {
		function a() {
			this.options_index = 0, this.parsed = []
		}
		return a.prototype.add_node = function(e) {
			return "OPTGROUP" === e.nodeName.toUpperCase() ? this.add_group(e) : this.add_option(e)
		}, a.prototype.add_group = function(u) {
			var m, w, v, p, l, q;
			for (m = this.parsed.length, this.parsed.push({
				array_index: m,
				group: !0,
				label: this.escapeExpression(u.label),
				children: 0,
				disabled: u.disabled,
				title: u.title,
				search_keys: d.trim(u.getAttribute("data-keys") || "").replace(/,/g, " ")
			}), l = u.childNodes, q = [], v = 0, p = l.length; v < p; v++) {
				w = l[v], q.push(this.add_option(w, m, u.disabled))
			}
			return q
		}, a.prototype.add_option = function(m, l, o) {
			if ("OPTION" === m.nodeName.toUpperCase()) {
				return "" !== m.text ? (null != l && (this.parsed[l].children += 1), this.parsed.push({
					array_index: this.parsed.length,
					options_index: this.options_index,
					value: m.value,
					text: m.text,
					title: m.title,
					html: m.innerHTML,
					selected: m.selected,
					disabled: o === !0 ? o : m.disabled,
					group_array_index: l,
					classes: m.className,
					style: m.style.cssText,
					data: m.getAttribute("data-data"),
					search_keys: (d.trim(m.getAttribute("data-keys") || "") + m.value).replace(/,/, " ")
				})) : this.parsed.push({
					array_index: this.parsed.length,
					options_index: this.options_index,
					empty: !0
				}), this.options_index += 1
			}
		}, a.prototype.escapeExpression = function(m) {
			var n, l;
			return null == m || m === !1 ? "" : /[\&\<\>\"\'\`]/.test(m) ? (n = {
				"<": "&lt;",
				">": "&gt;",
				'"': "&quot;",
				"'": "&#x27;",
				"`": "&#x60;"
			}, l = /&(?!\w+;)|[\<\>\"\'\`]/g, m.replace(l, function(e) {
				return n[e] || "&amp;"
			})) : m
		}, a
	}(), k.select_to_array = function(n) {
		var q, m, r, p, l;
		for (m = new k, l = n.childNodes, r = 0, p = l.length; r < p; r++) {
			q = l[r], m.add_node(q)
		}
		return m.parsed
	}, h = function() {
		function a(e, l) {
			this.form_field = e, this.options = null != l ? l : {}, a.browser_is_supported() && (this.lang = g[this.options.lang || (d.zui.clientLang ? d.zui.clientLang() : "zh_cn")], this.is_multiple = this.form_field.multiple, this.set_default_text(), this.set_default_values(), this.setup(), this.set_up_html(), this.register_observers())
		}
		return a.prototype.set_default_values = function() {
			var e = this;
			return this.click_test_action = function(l) {
				return e.test_active_click(l)
			}, this.activate_action = function(l) {
				return e.activate_field(l)
			}, this.active_field = !1, this.mouse_on_container = !1, this.results_showing = !1, this.result_highlighted = null, this.allow_single_deselect = null != this.options.allow_single_deselect && null != this.form_field.options[0] && "" === this.form_field.options[0].text && this.options.allow_single_deselect, this.disable_search_threshold = this.options.disable_search_threshold || 0, this.disable_search = this.options.disable_search || !1, this.enable_split_word_search = null == this.options.enable_split_word_search || this.options.enable_split_word_search, this.group_search = null == this.options.group_search || this.options.group_search, this.search_contains = this.options.search_contains || !1, this.single_backstroke_delete = null == this.options.single_backstroke_delete || this.options.single_backstroke_delete, this.max_selected_options = this.options.max_selected_options || 1 / 0, this.drop_direction = this.options.drop_direction || "auto", this.middle_highlight = this.options.middle_highlight, this.compact_search = this.options.compact_search || !0, this.inherit_select_classes = this.options.inherit_select_classes || !1, this.display_selected_options = null == this.options.display_selected_options || this.options.display_selected_options, this.display_disabled_options = null == this.options.display_disabled_options || this.options.display_disabled_options
		}, a.prototype.set_default_text = function() {
			return this.form_field.getAttribute("data-placeholder") ? this.default_text = this.form_field.getAttribute("data-placeholder") : this.is_multiple ? this.default_text = this.options.placeholder_text_multiple || this.options.placeholder_text || a.default_multiple_text : this.default_text = this.options.placeholder_text_single || this.options.placeholder_text || a.default_single_text, this.results_none_found = this.form_field.getAttribute("data-no_results_text") || this.options.no_results_text || this.lang.no_results_text || a.default_no_result_text
		}, a.prototype.mouse_enter = function() {
			return this.mouse_on_container = !0
		}, a.prototype.mouse_leave = function() {
			return this.mouse_on_container = !1
		}, a.prototype.input_focus = function(l) {
			var m = this;
			if (this.is_multiple) {
				if (!this.active_field) {
					return setTimeout(function() {
						return m.container_mousedown()
					}, 50)
				}
			} else {
				if (!this.active_field) {
					return this.activate_field()
				}
			}
		}, a.prototype.input_blur = function(l) {
			var m = this;
			if (!this.mouse_on_container) {
				return this.active_field = !1, setTimeout(function() {
					return m.blur_test()
				}, 100)
			}
		}, a.prototype.results_option_build = function(m) {
			var q, l, u, r, p;
			for (q = "", p = this.results_data, u = 0, r = p.length; u < r; u++) {
				l = p[u], q += l.group ? this.result_add_group(l) : this.result_add_option(l), (null != m ? m.first : void 0) && (l.selected && this.is_multiple ? this.choice_build(l) : l.selected && !this.is_multiple && this.single_set_selected_text(l.text))
			}
			return q
		}, a.prototype.result_add_option = function(m) {
			var n, l;
			return m.search_match && this.include_option_in_results(m) ? (n = [], m.disabled || m.selected && this.is_multiple || n.push("active-result"), !m.disabled || m.selected && this.is_multiple || n.push("disabled-result"), m.selected && n.push("result-selected"), null != m.group_array_index && n.push("group-option"), "" !== m.classes && n.push(m.classes), l = document.createElement("li"), l.className = n.join(" "), l.style.cssText = m.style, l.title = m.title, l.setAttribute("data-option-array-index", m.array_index), l.setAttribute("data-data", m.data), l.innerHTML = m.search_text, this.outerHTML(l)) : ""
		}, a.prototype.result_add_group = function(l) {
			var m;
			return (l.search_match || l.group_match) && l.active_options > 0 ? (m = document.createElement("li"), m.className = "group-result", m.title = l.title, m.innerHTML = l.search_text, this.outerHTML(m)) : ""
		}, a.prototype.results_update_field = function() {
			if (this.set_default_text(), this.is_multiple || this.results_reset_cleanup(), this.result_clear_highlight(), this.results_build(), this.results_showing) {
				return this.winnow_results()
			}
		}, a.prototype.reset_single_select_options = function() {
			var m, p, l, r, q;
			for (r = this.results_data, q = [], p = 0, l = r.length; p < l; p++) {
				m = r[p], m.selected ? q.push(m.selected = !1) : q.push(void 0)
			}
			return q
		}, a.prototype.results_toggle = function() {
			return this.results_showing ? this.results_hide() : this.results_show()
		}, a.prototype.results_search = function(e) {
			return this.results_showing ? this.winnow_results(1) : this.results_show()
		}, a.prototype.winnow_results = function(F) {
			var A, y, w, v, G, D, m, x, z, C, B, E, q;
			for (this.no_results_clear(), G = 0, m = this.get_search_text(), A = m.replace(/[-[\]{}()*+?.,\\^$|#\s]/g, "\\$&"), v = this.search_contains ? "" : "^", w = new RegExp(v + A, "i"), C = new RegExp(A, "i"), q = this.results_data, B = 0, E = q.length; B < E; B++) {
				y = q[B], y.search_match = !1, D = null, this.include_option_in_results(y) && (y.group && (y.group_match = !1, y.active_options = 0), null != y.group_array_index && this.results_data[y.group_array_index] && (D = this.results_data[y.group_array_index], 0 === D.active_options && D.search_match && (G += 1), D.active_options += 1), y.group && !this.group_search || (y.search_text = y.group ? y.label : y.html, y.search_keys_match = this.search_string_match(y.search_keys, w), y.search_text_match = this.search_string_match(y.search_text, w), y.search_match = y.search_text_match || y.search_keys_match, y.search_match && !y.group && (G += 1), y.search_match ? (y.search_text_match && y.search_text.length ? (x = y.search_text.search(C), z = y.search_text.substr(0, x + m.length) + "</em>" + y.search_text.substr(x + m.length), y.search_text = z.substr(0, x) + "<em>" + z.substr(x)) : y.search_keys_match && y.search_keys.length && (x = y.search_keys.search(C), z = y.search_keys.substr(0, x + m.length) + "</em>" + y.search_keys.substr(x + m.length), y.search_text += '&nbsp; <small style="opacity: 0.7">' + z.substr(0, x) + "<em>" + z.substr(x) + "</small>"), null != D && (D.group_match = !0)) : null != y.group_array_index && this.results_data[y.group_array_index].search_match && (y.search_match = !0)))
			}
			return this.result_clear_highlight(), G < 1 && m.length ? (this.update_results_content(""), this.no_results(m)) : (this.update_results_content(this.results_option_build()), this.winnow_results_set_highlight(F))
		}, a.prototype.search_string_match = function(m, q) {
			var l, u, r, p;
			if (q.test(m)) {
				return !0
			}
			if (this.enable_split_word_search && (m.indexOf(" ") >= 0 || 0 === m.indexOf("[")) && (u = m.replace(/\[|\]/g, "").split(" "), u.length)) {
				for (r = 0, p = u.length; r < p; r++) {
					if (l = u[r], q.test(l)) {
						return !0
					}
				}
			}
		}, a.prototype.choices_count = function() {
			var m, o, l, p;
			if (null != this.selected_option_count) {
				return this.selected_option_count
			}
			for (this.selected_option_count = 0, p = this.form_field.options, o = 0, l = p.length; o < l; o++) {
				m = p[o], m.selected && "" != m.value && (this.selected_option_count += 1)
			}
			return this.selected_option_count
		}, a.prototype.choices_click = function(e) {
			if (e.preventDefault(), !this.results_showing && !this.is_disabled) {
				return this.results_show()
			}
		}, a.prototype.keyup_checker = function(m) {
			var n, l;
			switch (n = null != (l = m.which) ? l : m.keyCode, this.search_field_scale(), n) {
			case 8:
				if (this.is_multiple && this.backstroke_length < 1 && this.choices_count() > 0) {
					return this.keydown_backstroke()
				}
				if (!this.pending_backstroke) {
					return this.result_clear_highlight(), this.results_search()
				}
				break;
			case 13:
				if (m.preventDefault(), this.results_showing) {
					return this.result_select(m)
				}
				break;
			case 27:
				return this.results_showing && this.results_hide(), !0;
			case 9:
			case 38:
			case 40:
			case 16:
			case 91:
			case 17:
				break;
			default:
				return this.results_search()
			}
		}, a.prototype.clipboard_event_checker = function(l) {
			var m = this;
			return setTimeout(function() {
				return m.results_search()
			}, 50)
		}, a.prototype.container_width = function() {
			return null != this.options.width ? this.options.width : this.form_field && this.form_field.classList && this.form_field.classList.contains("form-control") ? "100%" : "" + this.form_field.offsetWidth + "px"
		}, a.prototype.include_option_in_results = function(e) {
			return !(this.is_multiple && !this.display_selected_options && e.selected) && (!(!this.display_disabled_options && e.disabled) && !e.empty)
		}, a.prototype.search_results_touchstart = function(e) {
			return this.touch_started = !0, this.search_results_mouseover(e)
		}, a.prototype.search_results_touchmove = function(e) {
			return this.touch_started = !1, this.search_results_mouseout(e)
		}, a.prototype.search_results_touchend = function(e) {
			if (this.touch_started) {
				return this.search_results_mouseup(e)
			}
		}, a.prototype.outerHTML = function(l) {
			var m;
			return l.outerHTML ? l.outerHTML : (m = document.createElement("div"), m.appendChild(l), m.innerHTML)
		}, a.browser_is_supported = function() {
			return "Microsoft Internet Explorer" === window.navigator.appName ? document.documentMode >= 8 : !/iP(od|hone)/i.test(window.navigator.userAgent) && (!/Android/i.test(window.navigator.userAgent) || !/Mobile/i.test(window.navigator.userAgent))
		}, a.default_multiple_text = "", a.default_single_text = "", a.default_no_result_text = "No results match", a
	}(), d = jQuery, d.fn.extend({
		chosen: function(a) {
			return h.browser_is_supported() ? this.each(function(m) {
				var n, l;
				n = d(this), l = n.data("chosen"), "destroy" === a && l ? l.destroy() : l || n.data("chosen", new c(this, a))
			}) : this
		}
	}), c = function(l) {
		function a() {
			return j = a.__super__.constructor.apply(this, arguments)
		}
		return b(a, l), a.prototype.setup = function() {
			return this.form_field_jq = d(this.form_field), this.current_selectedIndex = this.form_field.selectedIndex, this.is_rtl = this.form_field_jq.hasClass("chosen-rtl")
		}, a.prototype.set_up_html = function() {
			var o, m;
			o = ["chosen-container"], o.push("chosen-container-" + (this.is_multiple ? "multi" : "single")), this.inherit_select_classes && this.form_field.className && o.push(this.form_field.className), this.is_rtl && o.push("chosen-rtl");
			var p = this.form_field.getAttribute("data-css-class");
			return p && o.push(p), m = {
				"class": o.join(" "),
				style: "width: " + this.container_width() + ";",
				title: this.form_field.title
			}, this.form_field.id.length && (m.id = this.form_field.id.replace(/[^\w]/g, "_") + "_chosen"), this.container = d("<div />", m), this.is_multiple ? this.container.html('<ul class="chosen-choices"><li class="search-field"><input type="text" value="' + this.default_text + '" class="default" autocomplete="off" style="width:25px;" /></li></ul><div class="chosen-drop"><ul class="chosen-results"></ul></div>') : (this.container.html('<a class="chosen-single chosen-default" tabindex="-1"><span>' + this.default_text + '</span><div><b></b></div><div class="chosen-search"><input type="text" autocomplete="off" /></div></a><div class="chosen-drop"><ul class="chosen-results"></ul></div>'), this.compact_search && this.container.find(".chosen-search").appendTo(this.container.find(".chosen-single"))), this.form_field_jq.hide().after(this.container), this.dropdown = this.container.find("div.chosen-drop").first(), this.search_field = this.container.find("input").first(), this.search_results = this.container.find("ul.chosen-results").first(), this.search_field_scale(), this.search_no_results = this.container.find("li.no-results").first(), this.is_multiple ? (this.search_choices = this.container.find("ul.chosen-choices").first(), this.search_container = this.container.find("li.search-field").first()) : (this.search_container = this.container.find("div.chosen-search").first(), this.selected_item = this.container.find(".chosen-single").first()), this.options.drop_width && this.dropdown.css("width", this.options.drop_width).addClass("chosen-drop-size-limited"), this.results_build(), this.set_tab_index(), this.set_label_behavior(), this.form_field_jq.trigger("chosen:ready", {
				chosen: this
			})
		}, a.prototype.register_observers = function() {
			var e = this;
			return this.container.bind("mousedown.chosen", function(m) {
				e.container_mousedown(m)
			}), this.container.bind("mouseup.chosen", function(m) {
				e.container_mouseup(m)
			}), this.container.bind("mouseenter.chosen", function(m) {
				e.mouse_enter(m)
			}), this.container.bind("mouseleave.chosen", function(m) {
				e.mouse_leave(m)
			}), this.search_results.bind("mouseup.chosen", function(m) {
				e.search_results_mouseup(m)
			}), this.search_results.bind("mouseover.chosen", function(m) {
				e.search_results_mouseover(m)
			}), this.search_results.bind("mouseout.chosen", function(m) {
				e.search_results_mouseout(m)
			}), this.search_results.bind("mousewheel.chosen DOMMouseScroll.chosen", function(m) {
				e.search_results_mousewheel(m)
			}), this.search_results.bind("touchstart.chosen", function(m) {
				e.search_results_touchstart(m)
			}), this.search_results.bind("touchmove.chosen", function(m) {
				e.search_results_touchmove(m)
			}), this.search_results.bind("touchend.chosen", function(m) {
				e.search_results_touchend(m)
			}), this.form_field_jq.bind("chosen:updated.chosen", function(m) {
				e.results_update_field(m)
			}), this.form_field_jq.bind("chosen:activate.chosen", function(m) {
				e.activate_field(m)
			}), this.form_field_jq.bind("chosen:open.chosen", function(m) {
				e.container_mousedown(m)
			}), this.form_field_jq.bind("chosen:close.chosen", function(m) {
				e.input_blur(m)
			}), this.search_field.bind("blur.chosen", function(m) {
				e.input_blur(m)
			}), this.search_field.bind("keyup.chosen", function(m) {
				e.keyup_checker(m)
			}), this.search_field.bind("keydown.chosen", function(m) {
				e.keydown_checker(m)
			}), this.search_field.bind("focus.chosen", function(m) {
				e.input_focus(m)
			}), this.search_field.bind("cut.chosen", function(m) {
				e.clipboard_event_checker(m)
			}), this.search_field.bind("paste.chosen", function(m) {
				e.clipboard_event_checker(m)
			}), this.is_multiple ? this.search_choices.bind("click.chosen", function(m) {
				e.choices_click(m)
			}) : this.container.bind("click.chosen", function(m) {
				m.preventDefault()
			})
		}, a.prototype.destroy = function() {
			return d(this.container[0].ownerDocument).unbind("click.chosen", this.click_test_action), this.search_field[0].tabIndex && (this.form_field_jq[0].tabIndex = this.search_field[0].tabIndex), this.container.remove(), this.form_field_jq.removeData("chosen"), this.form_field_jq.show()
		}, a.prototype.search_field_disabled = function() {
			return this.is_disabled = this.form_field_jq[0].disabled, this.is_disabled ? (this.container.addClass("chosen-disabled"), this.search_field[0].disabled = !0, this.is_multiple || this.selected_item.unbind("focus.chosen", this.activate_action), this.close_field()) : (this.container.removeClass("chosen-disabled"), this.search_field[0].disabled = !1, this.is_multiple ? void 0 : this.selected_item.bind("focus.chosen", this.activate_action))
		}, a.prototype.container_mousedown = function(m) {
			if (!this.is_disabled && (m && "mousedown" === m.type && !this.results_showing && m.preventDefault(), null == m || !d(m.target).hasClass("search-choice-close"))) {
				return this.active_field ? this.is_multiple || !m || d(m.target)[0] !== this.selected_item[0] && !d(m.target).parents("a.chosen-single").length || (m.preventDefault(), this.results_toggle()) : (this.is_multiple && this.search_field.val(""), d(this.container[0].ownerDocument).bind("click.chosen", this.click_test_action), this.results_show()), this.activate_field()
			}
		}, a.prototype.container_mouseup = function(e) {
			if ("ABBR" === e.target.nodeName && !this.is_disabled) {
				return this.results_reset(e)
			}
		}, a.prototype.search_results_mousewheel = function(m) {
			var n;
			if (m.originalEvent && (n = -m.originalEvent.wheelDelta || m.originalEvent.detail), null != n) {
				return m.preventDefault(), "DOMMouseScroll" === m.type && (n = 40 * n), this.search_results.scrollTop(n + this.search_results.scrollTop())
			}
		}, a.prototype.blur_test = function(e) {
			if (!this.active_field && this.container.hasClass("chosen-container-active")) {
				return this.close_field()
			}
		}, a.prototype.close_field = function() {
			return d(this.container[0].ownerDocument).unbind("click.chosen", this.click_test_action), this.active_field = !1, this.results_hide(), this.container.removeClass("chosen-container-active"), this.clear_backstroke(), this.show_search_field_default(), this.search_field_scale()
		}, a.prototype.activate_field = function() {
			return this.container.addClass("chosen-container-active"), this.active_field = !0, this.search_field.val(this.search_field.val()), this.search_field.focus()
		}, a.prototype.test_active_click = function(n) {
			var m;
			return m = d(n.target).closest(".chosen-container"), m.length && this.container[0] === m[0] ? this.active_field = !0 : this.close_field()
		}, a.prototype.results_build = function() {
			return this.parsing = !0, this.selected_option_count = null, this.results_data = k.select_to_array(this.form_field), this.is_multiple ? this.search_choices.find("li.search-choice").remove() : this.is_multiple || (this.single_set_selected_text(), this.disable_search || this.form_field.options.length <= this.disable_search_threshold ? (this.search_field[0].readOnly = !0, this.container.addClass("chosen-container-single-nosearch"), this.container.removeClass("chosen-with-search")) : (this.search_field[0].readOnly = !1, this.container.removeClass("chosen-container-single-nosearch"), this.container.addClass("chosen-with-search"))), this.update_results_content(this.results_option_build({
				first: !0
			})), this.search_field_disabled(), this.show_search_field_default(), this.search_field_scale(), this.parsing = !1
		}, a.prototype.result_do_highlight = function(y, w) {
			var v, q, p, z, x, m, u = -1;
			y.length && (this.result_clear_highlight(), this.result_highlight = y, this.result_highlight.addClass("highlighted"), p = parseInt(this.search_results.css("maxHeight"), 10), m = this.result_highlight.outerHeight(), x = this.search_results.scrollTop(), z = p + x, q = this.result_highlight.position().top + this.search_results.scrollTop(), v = q + m, this.middle_highlight && (w || "always" === this.middle_highlight || v >= z || q < x) ? u = Math.min(q - m, Math.max(0, q - (p - m) / 2)) : v >= z ? u = v - p > 0 ? v - p : 0 : q < x && (u = q), u > -1 && this.search_results.scrollTop(u))
		}, a.prototype.result_clear_highlight = function() {
			return this.result_highlight && this.result_highlight.removeClass("highlighted"), this.result_highlight = null
		}, a.prototype.results_show = function() {
			if (this.is_multiple && this.max_selected_options <= this.choices_count()) {
				return this.form_field_jq.trigger("chosen:maxselected", {
					chosen: this
				}), !1
			}
			this.results_showing = !0, this.search_field.focus(), this.search_field.val(this.search_field.val()), this.winnow_results(1);
			var o = this.drop_direction;
			if (d.isFunction(o) && (o = o.call(this)), "auto" === o) {
				if (this.drop_directionFixed) {
					o = this.drop_directionFixed
				} else {
					var m = this.container.find(".chosen-drop"),
						p = this.container.offset();
					p.top + m.outerHeight() + 30 > d(window).height() + d(window).scrollTop() && (o = "up"), this.drop_directionFixed = o
				}
			}
			return this.container.toggleClass("chosen-up", "up" === o).addClass("chosen-with-drop"), this.form_field_jq.trigger("chosen:showing_dropdown", {
				chosen: this
			})
		}, a.prototype.update_results_content = function(e) {
			return this.search_results.html(e)
		}, a.prototype.results_hide = function() {
			return this.results_showing && (this.result_clear_highlight(), this.container.removeClass("chosen-with-drop"), this.form_field_jq.trigger("chosen:hiding_dropdown", {
				chosen: this
			}), this.drop_directionFixed = 0), this.results_showing = !1
		}, a.prototype.set_tab_index = function(m) {
			var n;
			if (this.form_field.tabIndex) {
				return n = this.form_field.tabIndex, this.form_field.tabIndex = -1, this.search_field[0].tabIndex = n
			}
		}, a.prototype.set_label_behavior = function() {
			var m = this;
			if (this.form_field_label = this.form_field_jq.parents("label"), !this.form_field_label.length && this.form_field.id.length && (this.form_field_label = d("label[for='" + this.form_field.id + "']")), this.form_field_label.length > 0) {
				return this.form_field_label.bind("click.chosen", function(e) {
					return m.is_multiple ? m.container_mousedown(e) : m.activate_field()
				})
			}
		}, a.prototype.show_search_field_default = function() {
			return this.is_multiple && this.choices_count() < 1 && !this.active_field ? (this.search_field.val(this.default_text), this.search_field.addClass("default")) : (this.search_field.val(""), this.search_field.removeClass("default"))
		}, a.prototype.search_results_mouseup = function(n) {
			var m;
			if (m = d(n.target).hasClass("active-result") ? d(n.target) : d(n.target).parents(".active-result").first(), m.length) {
				return this.result_highlight = m, this.result_select(n), this.search_field.focus()
			}
		}, a.prototype.search_results_mouseover = function(n) {
			var m;
			if (m = d(n.target).hasClass("active-result") ? d(n.target) : d(n.target).parents(".active-result").first()) {
				return this.result_do_highlight(m)
			}
		}, a.prototype.search_results_mouseout = function(m) {
			if (d(m.target).hasClass("active-result")) {
				return this.result_clear_highlight()
			}
		}, a.prototype.choice_build = function(p) {
			var m, r, q = this;
			return m = d("<li />", {
				"class": "search-choice"
			}).html("<span title='" + p.html + "'>" + p.html + "</span>"), p.disabled ? m.addClass("search-choice-disabled") : (r = d("<a />", {
				"class": "search-choice-close",
				"data-option-array-index": p.array_index
			}), r.bind("click.chosen", function(e) {
				return q.choice_destroy_link_click(e)
			}), m.append(r)), this.search_container.before(m)
		}, a.prototype.choice_destroy_link_click = function(m) {
			if (m.preventDefault(), m.stopPropagation(), !this.is_disabled) {
				return this.choice_destroy(d(m.target))
			}
		}, a.prototype.choice_destroy = function(e) {
			if (this.result_deselect(e[0].getAttribute("data-option-array-index"))) {
				return this.show_search_field_default(), this.is_multiple && this.choices_count() > 0 && this.search_field.val().length < 1 && this.results_hide(), e.parents("li").first().remove(), this.search_field_scale()
			}
		}, a.prototype.results_reset = function() {
			if (this.reset_single_select_options(), this.form_field.options[0].selected = !0, this.single_set_selected_text(), this.show_search_field_default(), this.results_reset_cleanup(), this.form_field_jq.trigger("change"), this.active_field) {
				return this.results_hide()
			}
		}, a.prototype.results_reset_cleanup = function() {
			return this.current_selectedIndex = this.form_field.selectedIndex, this.selected_item.find("abbr").remove()
		}, a.prototype.result_select = function(n) {
			var o, m;
			if (this.result_highlight) {
				return o = this.result_highlight, this.result_clear_highlight(), this.is_multiple && this.max_selected_options <= this.choices_count() ? (this.form_field_jq.trigger("chosen:maxselected", {
					chosen: this
				}), !1) : (this.is_multiple ? o.removeClass("active-result") : this.reset_single_select_options(), m = this.results_data[o[0].getAttribute("data-option-array-index")], m.selected = !0, this.form_field.options[m.options_index].selected = !0, this.selected_option_count = null, this.is_multiple ? this.choice_build(m) : this.single_set_selected_text(m.text), (n.metaKey || n.ctrlKey) && this.is_multiple || this.results_hide(), this.search_field.val(""), (this.is_multiple || this.form_field.selectedIndex !== this.current_selectedIndex) && this.form_field_jq.trigger("change", {
					selected: this.form_field.options[m.options_index].value
				}), this.current_selectedIndex = this.form_field.selectedIndex, this.search_field_scale())
			}
		}, a.prototype.single_set_selected_text = function(e) {
			return null == e && (e = this.default_text), e === this.default_text ? this.selected_item.addClass("chosen-default") : (this.single_deselect_control_build(), this.selected_item.removeClass("chosen-default")), this.compact_search && this.search_field.attr("placeholder", e), this.selected_item.find("span").attr("title", e).text(e)
		}, a.prototype.result_deselect = function(m) {
			var n;
			return n = this.results_data[m], !this.form_field.options[n.options_index].disabled && (n.selected = !1, this.form_field.options[n.options_index].selected = !1, this.selected_option_count = null, this.result_clear_highlight(), this.results_showing && this.winnow_results(), this.form_field_jq.trigger("change", {
				deselected: this.form_field.options[n.options_index].value
			}), this.search_field_scale(), !0)
		}, a.prototype.single_deselect_control_build = function() {
			if (this.allow_single_deselect) {
				return this.selected_item.find("abbr").length || this.selected_item.find("span").first().after('<abbr class="search-choice-close"></abbr>'), this.selected_item.addClass("chosen-single-with-deselect")
			}
		}, a.prototype.get_search_text = function() {
			return this.search_field.val() === this.default_text ? "" : d("<div/>").text(d.trim(this.search_field.val())).html()
		}, a.prototype.winnow_results_set_highlight = function(n) {
			var o, m;
			if (m = this.is_multiple ? [] : this.search_results.find(".result-selected.active-result"), o = m.length ? m.first() : this.search_results.find(".active-result").first(), null != o) {
				return this.result_do_highlight(o, n)
			}
		}, a.prototype.no_results = function(n) {
			var m;
			return m = d('<li class="no-results">' + this.results_none_found + ' "<span></span>"</li>'), m.find("span").first().html(n), this.search_results.append(m), this.form_field_jq.trigger("chosen:no_results", {
				chosen: this
			})
		}, a.prototype.no_results_clear = function() {
			return this.search_results.find(".no-results").remove()
		}, a.prototype.keydown_arrow = function() {
			var e;
			return this.results_showing && this.result_highlight ? (e = this.result_highlight.nextAll("li.active-result").first()) ? this.result_do_highlight(e) : void 0 : this.results_show()
		}, a.prototype.keyup_arrow = function() {
			var e;
			return this.results_showing || this.is_multiple ? this.result_highlight ? (e = this.result_highlight.prevAll("li.active-result"), e.length ? this.result_do_highlight(e.first()) : (this.choices_count() > 0 && this.results_hide(), this.result_clear_highlight())) : void 0 : this.results_show()
		}, a.prototype.keydown_backstroke = function() {
			var e;
			return this.pending_backstroke ? (this.choice_destroy(this.pending_backstroke.find("a").first()), this.clear_backstroke()) : (e = this.search_container.siblings("li.search-choice").last(), e.length && !e.hasClass("search-choice-disabled") ? (this.pending_backstroke = e, this.single_backstroke_delete ? this.keydown_backstroke() : this.pending_backstroke.addClass("search-choice-focus")) : void 0)
		}, a.prototype.clear_backstroke = function() {
			return this.pending_backstroke && this.pending_backstroke.removeClass("search-choice-focus"), this.pending_backstroke = null
		}, a.prototype.keydown_checker = function(n) {
			var o, m;
			switch (o = null != (m = n.which) ? m : n.keyCode, this.search_field_scale(), 8 !== o && this.pending_backstroke && this.clear_backstroke(), o) {
			case 8:
				this.backstroke_length = this.search_field.val().length;
				break;
			case 9:
				this.results_showing && !this.is_multiple && this.result_select(n), this.mouse_on_container = !1;
				break;
			case 13:
				n.preventDefault();
				break;
			case 38:
				n.preventDefault(), this.keyup_arrow();
				break;
			case 40:
				n.preventDefault(), this.keydown_arrow()
			}
		}, a.prototype.search_field_scale = function() {
			var x, v, q, p, z, y, m, u, w;
			if (this.is_multiple) {
				for (q = 0, m = 0, z = "position:absolute; left: -1000px; top: -1000px; display:none;", y = ["font-size", "font-style", "font-weight", "font-family", "line-height", "text-transform", "letter-spacing"], u = 0, w = y.length; u < w; u++) {
					p = y[u], z += p + ":" + this.search_field.css(p) + ";"
				}
				return x = d("<div />", {
					style: z
				}), x.text(this.search_field.val()), d("body").append(x), m = x.width() + 25, x.remove(), v = this.container.outerWidth(), m > v - 10 && (m = v - 10), this.search_field.css({
					width: m + "px"
				})
			}
		}, a
	}(h)
}.call(this), function(b) {
	var c = "zui.selectable",
		a = function(e, g) {
			this.name = c, this.$ = b(e), this.id = b.zui.uuid(), this.selectOrder = 1, this.selections = {}, this.getOptions(g), this._init()
		},
		f = function(h, j, g) {
			return h >= g.left && h <= g.left + g.width && j >= g.top && j <= g.top + g.height
		},
		d = function(j, l) {
			var h = Math.max(j.left, l.left),
				m = Math.max(j.top, l.top),
				k = Math.min(j.left + j.width, l.left + l.width),
				g = Math.min(j.top + j.height, l.top + l.height);
			return f(h, m, j) && f(k, g, j) && f(h, m, l) && f(k, g, l)
		};
	a.DEFAULTS = {
		selector: "li,tr,div",
		trigger: "",
		selectClass: "active",
		rangeStyle: {
			border: "1px solid " + (b.zui.colorset ? b.zui.colorset.primary : "#3280fc"),
			backgroundColor: b.zui.colorset ? new b.zui.Color(b.zui.colorset.primary).fade(20).toCssStr() : "rgba(50, 128, 252, 0.2)"
		},
		clickBehavior: "toggle",
		ignoreVal: 3,
		listenClick: !0
	}, a.prototype.getOptions = function(g) {
		this.options = b.extend({}, a.DEFAULTS, this.$.data(), g)
	}, a.prototype.select = function(e) {
		this.toggle(e, !0)
	}, a.prototype.unselect = function(e) {
		this.toggle(e, !1)
	}, a.prototype.toggle = function(p, j, u) {
		var q, k, h = this.options.selector,
			m = this;
		if (void 0 === p) {
			return void this.$.find(h).each(function() {
				m.toggle(this, j)
			})
		}
		if ("object" == typeof p ? (q = b(p).closest(h), k = q.data("id")) : (k = p, q = m.$.find('.slectable-item[data-id="' + k + '"]')), q && q.length) {
			if (k || (k = b.zui.uuid(), q.attr("data-id", k)), void 0 !== j && null !== j || (j = !m.selections[k]), !! j != !! m.selections[k]) {
				var g;
				b.isFunction(u) && (g = u(j)), g !== !0 && (m.selections[k] = !! j && m.selectOrder++, m.callEvent(j ? "select" : "unselect", {
					id: k,
					selections: m.selections,
					target: q,
					selected: m.getSelectedArray()
				}, m))
			}
			m.options.selectClass && q.toggleClass(m.options.selectClass, j)
		}
	}, a.prototype.getSelectedArray = function() {
		var g = [];
		return b.each(this.selections, function(h, e) {
			e && g.push(h)
		}), g
	}, a.prototype.syncSelectionsFromClass = function() {
		var g = this;
		g.$children = g.$.find(g.options.selector);
		g.selections = {}, g.$children.each(function() {
			var e = b(this);
			g.selections[e.data("id")] = e.hasClass(g.options.selectClass)
		})
	}, a.prototype._init = function() {
		var M, I, F, A, Q, D, H, J = this.options,
			O = this,
			N = J.ignoreVal,
			z = !0,
			E = "." + this.name + "." + this.id,
			L = b.isFunction(J.checkFunc) ? J.checkFunc : null,
			K = b.isFunction(J.rangeFunc) ? J.rangeFunc : null,
			G = !1,
			q = null,
			j = "mousedown" + E,
			P = function() {
				A && O.$children.each(function() {
					var l = b(this),
						h = l.offset();
					h.width = l.outerWidth(), h.height = l.outerHeight();
					var m = K ? K.call(this, A, h) : d(A, h);
					if (L) {
						var g = L.call(O, {
							intersect: m,
							target: l,
							range: A,
							targetRange: h
						});
						g === !0 ? O.select(l) : g === !1 && O.unselect(l)
					} else {
						m ? O.select(l) : O.multiKey || O.unselect(l)
					}
				})
			},
			o = function(e) {
				G && (Q = e.pageX, D = e.pageY, A = {
					width: Math.abs(Q - M),
					height: Math.abs(D - I),
					left: Q > M ? M : Q,
					top: D > I ? I : D
				}, z && A.width < N && A.height < N || (F || (F = b('.selectable-range[data-id="' + O.id + '"]'), F.length || (F = b('<div class="selectable-range" data-id="' + O.id + '"></div>').css(b.extend({
					zIndex: 1060,
					position: "absolute",
					top: M,
					left: I,
					pointerEvents: "none"
				}, O.options.rangeStyle)).appendTo(b("body")))), F.css(A), clearTimeout(H), H = setTimeout(P, 10), z = !1))
			},
			k = function(g) {
				b(document).off(E), clearTimeout(q), G && (G = !1, F && F.remove(), z || A && (clearTimeout(H), P(), A = null), O.callEvent("finish", {
					selections: O.selections,
					selected: O.getSelectedArray()
				}), g.preventDefault())
			},
			B = function(l) {
				if (G) {
					return k(l)
				}
				var g = b.zui.getMouseButtonCode(J.mouseButton);
				if (!(g > -1 && l.button !== g || O.altKey || 3 === l.which || O.callEvent("start", l) === !1)) {
					var e = O.$children = O.$.find(J.selector);
					e.addClass("slectable-item");
					var h = O.multiKey ? "multi" : J.clickBehavior;
					if ("single" === h && O.unselect(), J.listenClick && ("multi" === h ? O.toggle(l.target) : "single" === h ? O.select(l.target) : "toggle" === h && O.toggle(l.target, null, function(m) {
						O.unselect()
					})), O.callEvent("startDrag", l) === !1) {
						return void O.callEvent("finish", {
							selections: O.selections,
							selected: O.getSelectedArray()
						})
					}
					M = l.pageX, I = l.pageY, F = null, z = !0, G = !0, b(document).on("mousemove" + E, o).on("mouseup" + E, k), q = setTimeout(function() {
						b(document).on(j, k)
					}, 10), l.preventDefault()
				}
			},
			R = J.container && "default" !== J.container ? b(J.container) : this.$;
		J.trigger ? R.on(j, J.trigger, B) : R.on(j, B), b(document).on("keydown", function(g) {
			var h = g.keyCode;
			17 === h || 91 == h ? O.multiKey = h : 18 === h && (O.altKey = !0)
		}).on("keyup", function(e) {
			O.multiKey = !1, O.altKey = !1
		})
	}, a.prototype.callEvent = function(j, g) {
		var l = b.Event(j + "." + this.name);
		this.$.trigger(l, g);
		var k = l.result,
			h = this.options[j];
		return b.isFunction(h) && (k = h.apply(this, b.isArray(g) ? g : [g])), k
	}, b.fn.selectable = function(e) {
		return this.each(function() {
			var j = b(this),
				h = j.data(c),
				g = "object" == typeof e && e;
			h || j.data(c, h = new a(this, g)), "string" == typeof e && h[e]()
		})
	}, b.fn.selectable.Constructor = a, b(function() {
		b('[data-ride="selectable"]').selectable()
	})
}(jQuery), +
function(d, g, c) {
	if (!d.fn.droppable) {
		return void console.error("Sortable requires droppable.js")
	}
	var j = "zui.sortable",
		h = {
			selector: "li,div",
			dragCssClass: "invisible",
			sortingClass: "sortable-sorting"
		},
		f = "order",
		b = function(k, a) {
			var l = this;
			l.$ = d(k), l.options = d.extend({}, h, l.$.data(), a), l.init()
		};
	b.DEFAULTS = h, b.NAME = j, b.prototype.init = function() {
		var w, s = this,
			p = s.$,
			m = s.options,
			z = m.selector,
			k = m.containerSelector,
			q = m.sortingClass,
			v = m.dragCssClass,
			y = m.targetSelector,
			x = m.reverse,
			A = function(a) {
				a = a || s.getItems(1);
				var l = a.length;
				l && a.each(function(o) {
					var n = x ? l - o : o;
					d(this).attr("data-" + f, n).data(f, n)
				})
			};
		A(), p.droppable({
			handle: m.trigger,
			target: y ? y : k ? z + "," + k : z,
			selector: z,
			container: p,
			always: m.always,
			flex: !0,
			lazy: m.lazy,
			canMoveHere: m.canMoveHere,
			dropToClass: m.dropToClass,
			before: m.before,
			nested: !! k,
			mouseButton: m.mouseButton,
			stopPropagation: m.stopPropagation,
			start: function(a) {
				v && a.element.addClass(v), w = !1, s.trigger("start", a)
			},
			drag: function(a) {
				if (p.addClass(q), a.isIn) {
					var u = a.element,
						e = a.target,
						B = k && e.is(k);
					if (B) {
						if (!e.children(z).filter(".dragging").length) {
							e.append(u);
							var r = s.getItems(1);
							A(r), s.trigger(f, {
								list: r,
								element: u
							})
						}
						return
					}
					var n = u.data(f),
						l = e.data(f);
					if (n === l) {
						return A(r)
					}
					n > l ? e[x ? "after" : "before"](u) : e[x ? "before" : "after"](u), w = !0;
					var r = s.getItems(1);
					A(r), s.trigger(f, {
						list: r,
						element: u
					})
				}
			},
			finish: function(a) {
				v && a.element && a.element.removeClass(v), p.removeClass(q), s.trigger("finish", {
					list: s.getItems(),
					element: a.element,
					changed: w
				})
			}
		})
	}, b.prototype.destroy = function() {
		this.$.droppable("destroy"), this.$.data(j, null)
	}, b.prototype.reset = function() {
		this.destroy(), this.init()
	}, b.prototype.getItems = function(k) {
		var a = this.$.find(this.options.selector).not(".drag-shadow");
		return k ? a : a.map(function() {
			var l = d(this);
			return {
				item: l,
				order: l.data("order")
			}
		})
	}, b.prototype.trigger = function(k, a) {
		return d.zui.callEvent(this.options[k], a, this)
	}, d.fn.sortable = function(a) {
		return this.each(function() {
			var e = d(this),
				l = e.data(j),
				k = "object" == typeof a && a;
			l ? "object" == typeof a && l.reset() : e.data(j, l = new b(this, k)), "string" == typeof a && l[a]()
		})
	}, d.fn.sortable.Constructor = b
}(jQuery, window, document),
/*!
 * jQuery Form Plugin
 * version: 4.2.2
 * Requires jQuery v1.7.2 or later
 * Project repository: https://github.com/jquery-form/form

 * Copyright 2017 Kevin Morris
 * Copyright 2006 M. Alsup

 * Dual licensed under the LGPL-2.1+ or MIT licenses
 * https://github.com/jquery-form/form#license

 * This library is free software; you can redistribute it and/or
 * modify it under the terms of the GNU Lesser General Public
 * License as published by the Free Software Foundation; either
 * version 2.1 of the License, or (at your option) any later version.
 * This library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
 * Lesser General Public License for more details.
 */
function(r){"function"==typeof define&&define.amd?define(["jquery"],r):"object"==typeof module&&module.exports?module.exports=function(e,t){return void 0===t&&(t="undefined"!=typeof window?require("jquery"):require("jquery")(e)),r(t),t}:r(jQuery)}(function(q){"use strict";var m=/\r?\n/g,S={};S.fileapi=void 0!==q('<input type="file">').get(0).files,S.formdata=void 0!==window.FormData;var _=!!q.fn.prop;function o(e){var t=e.data;e.isDefaultPrevented()||(e.preventDefault(),q(e.target).closest("form").ajaxSubmit(t))}function i(e){var t=e.target,r=q(t);if(!r.is("[type=submit],[type=image]")){var a=r.closest("[type=submit]");if(0===a.length)return;t=a[0]}var n,o=t.form;"image"===(o.clk=t).type&&(void 0!==e.offsetX?(o.clk_x=e.offsetX,o.clk_y=e.offsetY):"function"==typeof q.fn.offset?(n=r.offset(),o.clk_x=e.pageX-n.left,o.clk_y=e.pageY-n.top):(o.clk_x=e.pageX-t.offsetLeft,o.clk_y=e.pageY-t.offsetTop)),setTimeout(function(){o.clk=o.clk_x=o.clk_y=null},100)}function N(){var e;q.fn.ajaxSubmit.debug&&(e="[jquery.form] "+Array.prototype.join.call(arguments,""),window.console&&window.console.log?window.console.log(e):window.opera&&window.opera.postError&&window.opera.postError(e))}q.fn.attr2=function(){if(!_)return this.attr.apply(this,arguments);var e=this.prop.apply(this,arguments);return e&&e.jquery||"string"==typeof e?e:this.attr.apply(this,arguments)},q.fn.ajaxSubmit=function(M,e,t,r){if(!this.length)return N("ajaxSubmit: skipping submit process - no element selected"),this;var O,a,n,o,X=this;"function"==typeof M?M={success:M}:"string"==typeof M||!1===M&&0<arguments.length?(M={url:M,data:e,dataType:t},"function"==typeof r&&(M.success=r)):void 0===M&&(M={}),O=M.method||M.type||this.attr2("method"),n=(n=(n="string"==typeof(a=M.url||this.attr2("action"))?q.trim(a):"")||window.location.href||"")&&(n.match(/^([^#]+)/)||[])[1],o=/(MSIE|Trident)/.test(navigator.userAgent||"")&&/^https/i.test(window.location.href||"")?"javascript:false":"about:blank",M=q.extend(!0,{url:n,success:q.ajaxSettings.success,type:O||q.ajaxSettings.type,iframeSrc:o},M);var i={};if(this.trigger("form-pre-serialize",[this,M,i]),i.veto)return N("ajaxSubmit: submit vetoed via form-pre-serialize trigger"),this;if(M.beforeSerialize&&!1===M.beforeSerialize(this,M))return N("ajaxSubmit: submit aborted via beforeSerialize callback"),this;var s=M.traditional;void 0===s&&(s=q.ajaxSettings.traditional);var u,c,C=[],l=this.formToArray(M.semantic,C,M.filtering);if(M.data&&(c=q.isFunction(M.data)?M.data(l):M.data,M.extraData=c,u=q.param(c,s)),M.beforeSubmit&&!1===M.beforeSubmit(l,this,M))return N("ajaxSubmit: submit aborted via beforeSubmit callback"),this;if(this.trigger("form-submit-validate",[l,this,M,i]),i.veto)return N("ajaxSubmit: submit vetoed via form-submit-validate trigger"),this;var f=q.param(l,s);u&&(f=f?f+"&"+u:u),"GET"===M.type.toUpperCase()?(M.url+=(0<=M.url.indexOf("?")?"&":"?")+f,M.data=null):M.data=f;var d,m,p,h=[];M.resetForm&&h.push(function(){X.resetForm()}),M.clearForm&&h.push(function(){X.clearForm(M.includeHidden)}),!M.dataType&&M.target?(d=M.success||function(){},h.push(function(e,t,r){var a=arguments,n=M.replaceTarget?"replaceWith":"html";q(M.target)[n](e).each(function(){d.apply(this,a)})})):M.success&&(q.isArray(M.success)?q.merge(h,M.success):h.push(M.success)),M.success=function(e,t,r){for(var a=M.context||this,n=0,o=h.length;n<o;n++)h[n].apply(a,[e,t,r||X,X])},M.error&&(m=M.error,M.error=function(e,t,r){var a=M.context||this;m.apply(a,[e,t,r,X])}),M.complete&&(p=M.complete,M.complete=function(e,t){var r=M.context||this;p.apply(r,[e,t,X])});var v=0<q("input[type=file]:enabled",this).filter(function(){return""!==q(this).val()}).length,g="multipart/form-data",x=X.attr("enctype")===g||X.attr("encoding")===g,y=S.fileapi&&S.formdata;N("fileAPI :"+y);var b,T=(v||x)&&!y;!1!==M.iframe&&(M.iframe||T)?M.closeKeepAlive?q.get(M.closeKeepAlive,function(){b=w(l)}):b=w(l):b=(v||x)&&y?function(e){for(var r=new FormData,t=0;t<e.length;t++)r.append(e[t].name,e[t].value);if(M.extraData){var a=function(e){var t,r,a=q.param(e,M.traditional).split("&"),n=a.length,o=[];for(t=0;t<n;t++)a[t]=a[t].replace(/\+/g," "),r=a[t].split("="),o.push([decodeURIComponent(r[0]),decodeURIComponent(r[1])]);return o}(M.extraData);for(t=0;t<a.length;t++)a[t]&&r.append(a[t][0],a[t][1])}M.data=null;var n=q.extend(!0,{},q.ajaxSettings,M,{contentType:!1,processData:!1,cache:!1,type:O||"POST"});M.uploadProgress&&(n.xhr=function(){var e=q.ajaxSettings.xhr();return e.upload&&e.upload.addEventListener("progress",function(e){var t=0,r=e.loaded||e.position,a=e.total;e.lengthComputable&&(t=Math.ceil(r/a*100)),M.uploadProgress(e,r,a,t)},!1),e});n.data=null;var o=n.beforeSend;return n.beforeSend=function(e,t){M.formData?t.data=M.formData:t.data=r,o&&o.call(this,e,t)},q.ajax(n)}(l):q.ajax(M),X.removeData("jqxhr").data("jqxhr",b);for(var j=0;j<C.length;j++)C[j]=null;return this.trigger("form-submit-notify",[this,M]),this;function w(e){var t,r,l,f,o,d,m,p,a,n,h,v,i=X[0],g=q.Deferred();if(g.abort=function(e){p.abort(e)},e)for(r=0;r<C.length;r++)t=q(C[r]),_?t.prop("disabled",!1):t.removeAttr("disabled");(l=q.extend(!0,{},q.ajaxSettings,M)).context=l.context||l,o="jqFormIO"+(new Date).getTime();var s=i.ownerDocument,u=X.closest("body");if(l.iframeTarget?(n=(d=q(l.iframeTarget,s)).attr2("name"))?o=n:d.attr2("name",o):(d=q('<iframe name="'+o+'" src="'+l.iframeSrc+'" />',s)).css({position:"absolute",top:"-1000px",left:"-1000px"}),m=d[0],p={aborted:0,responseText:null,responseXML:null,status:0,statusText:"n/a",getAllResponseHeaders:function(){},getResponseHeader:function(){},setRequestHeader:function(){},abort:function(e){var t="timeout"===e?"timeout":"aborted";N("aborting upload... "+t),this.aborted=1;try{m.contentWindow.document.execCommand&&m.contentWindow.document.execCommand("Stop")}catch(e){}d.attr("src",l.iframeSrc),p.error=t,l.error&&l.error.call(l.context,p,t,e),f&&q.event.trigger("ajaxError",[p,l,t]),l.complete&&l.complete.call(l.context,p,t)}},(f=l.global)&&0==q.active++&&q.event.trigger("ajaxStart"),f&&q.event.trigger("ajaxSend",[p,l]),l.beforeSend&&!1===l.beforeSend.call(l.context,p,l))return l.global&&q.active--,g.reject(),g;if(p.aborted)return g.reject(),g;(a=i.clk)&&(n=a.name)&&!a.disabled&&(l.extraData=l.extraData||{},l.extraData[n]=a.value,"image"===a.type&&(l.extraData[n+".x"]=i.clk_x,l.extraData[n+".y"]=i.clk_y));var x=1,y=2;function b(t){var r=null;try{t.contentWindow&&(r=t.contentWindow.document)}catch(e){N("cannot get iframe.contentWindow document: "+e)}if(r)return r;try{r=t.contentDocument?t.contentDocument:t.document}catch(e){N("cannot get iframe.contentDocument: "+e),r=t.document}return r}var c=q("meta[name=csrf-token]").attr("content"),T=q("meta[name=csrf-param]").attr("content");function j(){var e=X.attr2("target"),t=X.attr2("action"),r=X.attr("enctype")||X.attr("encoding")||"multipart/form-data";i.setAttribute("target",o),O&&!/post/i.test(O)||i.setAttribute("method","POST"),t!==l.url&&i.setAttribute("action",l.url),l.skipEncodingOverride||O&&!/post/i.test(O)||X.attr({encoding:"multipart/form-data",enctype:"multipart/form-data"}),l.timeout&&(v=setTimeout(function(){h=!0,A(x)},l.timeout));var a=[];try{if(l.extraData)for(var n in l.extraData)l.extraData.hasOwnProperty(n)&&(q.isPlainObject(l.extraData[n])&&l.extraData[n].hasOwnProperty("name")&&l.extraData[n].hasOwnProperty("value")?a.push(q('<input type="hidden" name="'+l.extraData[n].name+'">',s).val(l.extraData[n].value).appendTo(i)[0]):a.push(q('<input type="hidden" name="'+n+'">',s).val(l.extraData[n]).appendTo(i)[0]));l.iframeTarget||d.appendTo(u),m.attachEvent?m.attachEvent("onload",A):m.addEventListener("load",A,!1),setTimeout(function e(){try{var t=b(m).readyState;N("state = "+t),t&&"uninitialized"===t.toLowerCase()&&setTimeout(e,50)}catch(e){N("Server abort: ",e," (",e.name,")"),A(y),v&&clearTimeout(v),v=void 0}},15);try{i.submit()}catch(e){document.createElement("form").submit.apply(i)}}finally{i.setAttribute("action",t),i.setAttribute("enctype",r),e?i.setAttribute("target",e):X.removeAttr("target"),q(a).remove()}}T&&c&&(l.extraData=l.extraData||{},l.extraData[T]=c),l.forceSync?j():setTimeout(j,10);var w,S,k,D=50;function A(e){if(!p.aborted&&!k){if((S=b(m))||(N("cannot access response document"),e=y),e===x&&p)return p.abort("timeout"),void g.reject(p,"timeout");if(e===y&&p)return p.abort("server abort"),void g.reject(p,"error","server abort");if(S&&S.location.href!==l.iframeSrc||h){m.detachEvent?m.detachEvent("onload",A):m.removeEventListener("load",A,!1);var t,r="success";try{if(h)throw"timeout";var a="xml"===l.dataType||S.XMLDocument||q.isXMLDoc(S);if(N("isXml="+a),!a&&window.opera&&(null===S.body||!S.body.innerHTML)&&--D)return N("requeing onLoad callback, DOM not available"),void setTimeout(A,250);var n=S.body?S.body:S.documentElement;p.responseText=n?n.innerHTML:null,p.responseXML=S.XMLDocument?S.XMLDocument:S,a&&(l.dataType="xml"),p.getResponseHeader=function(e){return{"content-type":l.dataType}[e.toLowerCase()]},n&&(p.status=Number(n.getAttribute("status"))||p.status,p.statusText=n.getAttribute("statusText")||p.statusText);var o,i,s,u=(l.dataType||"").toLowerCase(),c=/(json|script|text)/.test(u);c||l.textarea?(o=S.getElementsByTagName("textarea")[0])?(p.responseText=o.value,p.status=Number(o.getAttribute("status"))||p.status,p.statusText=o.getAttribute("statusText")||p.statusText):c&&(i=S.getElementsByTagName("pre")[0],s=S.getElementsByTagName("body")[0],i?p.responseText=i.textContent?i.textContent:i.innerText:s&&(p.responseText=s.textContent?s.textContent:s.innerText)):"xml"===u&&!p.responseXML&&p.responseText&&(p.responseXML=F(p.responseText));try{w=E(p,u,l)}catch(e){r="parsererror",p.error=t=e||r}}catch(e){N("error caught: ",e),r="error",p.error=t=e||r}p.aborted&&(N("upload aborted"),r=null),p.status&&(r=200<=p.status&&p.status<300||304===p.status?"success":"error"),"success"===r?(l.success&&l.success.call(l.context,w,"success",p),g.resolve(p.responseText,"success",p),f&&q.event.trigger("ajaxSuccess",[p,l])):r&&(void 0===t&&(t=p.statusText),l.error&&l.error.call(l.context,p,r,t),g.reject(p,"error",t),f&&q.event.trigger("ajaxError",[p,l,t])),f&&q.event.trigger("ajaxComplete",[p,l]),f&&!--q.active&&q.event.trigger("ajaxStop"),l.complete&&l.complete.call(l.context,p,r),k=!0,l.timeout&&clearTimeout(v),setTimeout(function(){l.iframeTarget?d.attr("src",l.iframeSrc):d.remove(),p.responseXML=null},100)}}}var F=q.parseXML||function(e,t){return window.ActiveXObject?((t=new ActiveXObject("Microsoft.XMLDOM")).async="false",t.loadXML(e)):t=(new DOMParser).parseFromString(e,"text/xml"),t&&t.documentElement&&"parsererror"!==t.documentElement.nodeName?t:null},L=q.parseJSON||function(e){return window.eval("("+e+")")},E=function(e,t,r){var a=e.getResponseHeader("content-type")||"",n=("xml"===t||!t)&&0<=a.indexOf("xml"),o=n?e.responseXML:e.responseText;return n&&"parsererror"===o.documentElement.nodeName&&q.error&&q.error("parsererror"),r&&r.dataFilter&&(o=r.dataFilter(o,t)),"string"==typeof o&&(("json"===t||!t)&&0<=a.indexOf("json")?o=L(o):("script"===t||!t)&&0<=a.indexOf("javascript")&&q.globalEval(o)),o};return g}},q.fn.ajaxForm=function(e,t,r,a){if(("string"==typeof e||!1===e&&0<arguments.length)&&(e={url:e,data:t,dataType:r},"function"==typeof a&&(e.success=a)),(e=e||{}).delegation=e.delegation&&q.isFunction(q.fn.on),e.delegation||0!==this.length)return e.delegation?(q(document).off("submit.form-plugin",this.selector,o).off("click.form-plugin",this.selector,i).on("submit.form-plugin",this.selector,e,o).on("click.form-plugin",this.selector,e,i),this):(e.beforeFormUnbind&&e.beforeFormUnbind(this,e),this.ajaxFormUnbind().on("submit.form-plugin",e,o).on("click.form-plugin",e,i));var n={s:this.selector,c:this.context};return!q.isReady&&n.s?(N("DOM not ready, queuing ajaxForm"),q(function(){q(n.s,n.c).ajaxForm(e)})):N("terminating; zero elements found by selector"+(q.isReady?"":" (DOM not ready)")),this},q.fn.ajaxFormUnbind=function(){return this.off("submit.form-plugin click.form-plugin")},q.fn.formToArray=function(e,t,r){var a=[];if(0===this.length)return a;var n,o,i,s,u,c,l,f,d,m,p=this[0],h=this.attr("id"),v=(v=e||void 0===p.elements?p.getElementsByTagName("*"):p.elements)&&q.makeArray(v);if(h&&(e||/(Edge|Trident)\//.test(navigator.userAgent))&&(n=q(':input[form="'+h+'"]').get()).length&&(v=(v||[]).concat(n)),!v||!v.length)return a;for(q.isFunction(r)&&(v=q.map(v,r)),o=0,c=v.length;o<c;o++)if((m=(u=v[o]).name)&&!u.disabled)if(e&&p.clk&&"image"===u.type)p.clk===u&&(a.push({name:m,value:q(u).val(),type:u.type}),a.push({name:m+".x",value:p.clk_x},{name:m+".y",value:p.clk_y}));else if((s=q.fieldValue(u,!0))&&s.constructor===Array)for(t&&t.push(u),i=0,l=s.length;i<l;i++)a.push({name:m,value:s[i]});else if(S.fileapi&&"file"===u.type){t&&t.push(u);var g=u.files;if(g.length)for(i=0;i<g.length;i++)a.push({name:m,value:g[i],type:u.type});else a.push({name:m,value:"",type:u.type})}else null!=s&&(t&&t.push(u),a.push({name:m,value:s,type:u.type,required:u.required}));return e||!p.clk||(m=(d=(f=q(p.clk))[0]).name)&&!d.disabled&&"image"===d.type&&(a.push({name:m,value:f.val()}),a.push({name:m+".x",value:p.clk_x},{name:m+".y",value:p.clk_y})),a},q.fn.formSerialize=function(e){return q.param(this.formToArray(e))},q.fn.fieldSerialize=function(n){var o=[];return this.each(function(){var e=this.name;if(e){var t=q.fieldValue(this,n);if(t&&t.constructor===Array)for(var r=0,a=t.length;r<a;r++)o.push({name:e,value:t[r]});else null!=t&&o.push({name:this.name,value:t})}}),q.param(o)},q.fn.fieldValue=function(e){for(var t=[],r=0,a=this.length;r<a;r++){var n=this[r],o=q.fieldValue(n,e);null==o||o.constructor===Array&&!o.length||(o.constructor===Array?q.merge(t,o):t.push(o))}return t},q.fieldValue=function(e,t){var r=e.name,a=e.type,n=e.tagName.toLowerCase();if(void 0===t&&(t=!0),t&&(!r||e.disabled||"reset"===a||"button"===a||("checkbox"===a||"radio"===a)&&!e.checked||("submit"===a||"image"===a)&&e.form&&e.form.clk!==e||"select"===n&&-1===e.selectedIndex))return null;if("select"!==n)return q(e).val().replace(m,"\r\n");var o=e.selectedIndex;if(o<0)return null;for(var i=[],s=e.options,u="select-one"===a,c=u?o+1:s.length,l=u?o:0;l<c;l++){var f=s[l];if(f.selected&&!f.disabled){var d=(d=f.value)||(f.attributes&&f.attributes.value&&!f.attributes.value.specified?f.text:f.value);if(u)return d;i.push(d)}}return i},q.fn.clearForm=function(e){return this.each(function(){q("input,select,textarea",this).clearFields(e)})},q.fn.clearFields=q.fn.clearInputs=function(r){var a=/^(?:color|date|datetime|email|month|number|password|range|search|tel|text|time|url|week)$/i;return this.each(function(){var e=this.type,t=this.tagName.toLowerCase();a.test(e)||"textarea"===t?this.value="":"checkbox"===e||"radio"===e?this.checked=!1:"select"===t?this.selectedIndex=-1:"file"===e?/MSIE/.test(navigator.userAgent)?q(this).replaceWith(q(this).clone(!0)):q(this).val(""):r&&(!0===r&&/hidden/.test(e)||"string"==typeof r&&q(this).is(r))&&(this.value="")})},q.fn.resetForm=function(){return this.each(function(){var t=q(this),e=this.tagName.toLowerCase();switch(e){case"input":this.checked=this.defaultChecked;case"textarea":return this.value=this.defaultValue,!0;case"option":case"optgroup":var r=t.parents("select");return r.length&&r[0].multiple?"option"===e?this.selected=this.defaultSelected:t.find("option").resetForm():r.resetForm(),!0;case"select":return t.find("option").each(function(e){if(this.selected=this.defaultSelected,this.defaultSelected&&!t[0].multiple)return t[0].selectedIndex=e,!1}),!0;case"label":var a=q(t.attr("for")),n=t.find("input,select,textarea");return a[0]&&n.unshift(a[0]),n.resetForm(),!0;case"form":return"function"!=typeof this.reset&&("object"!=typeof this.reset||this.reset.nodeType)||this.reset(),!0;default:return t.find("form,input,label,select,textarea").resetForm(),!0}})},q.fn.enable=function(e){return void 0===e&&(e=!0),this.each(function(){this.disabled=!e})},q.fn.selected=function(r){return void 0===r&&(r=!0),this.each(function(){var e,t=this.type;"checkbox"===t||"radio"===t?this.checked=r:"option"===this.tagName.toLowerCase()&&(e=q(this).parent("select"),r&&e[0]&&"select-one"===e[0].type&&e.find("option").selected(!1),this.selected=r)})},q.fn.ajaxSubmit.debug=!1}),
/*!
 * jQuery Hotkeys Plugin
 * Copyright 2010, John Resig
 * Dual licensed under the MIT or GPL Version 2 licenses.
 *
 * Based upon the plugin by Tzury Bar Yochay:
 * http://github.com/tzuryby/hotkeys
 *
 * Original idea by:
 * Binny V A, http://www.openjs.com/scripts/events/keyboard_shortcuts/
 */

function(a) {
	function b(d) {
		if ("string" == typeof d.data) {
			var c = d.handler,
				f = d.data.toLowerCase().split(" ");
			d.handler = function(p) {
				if (this === p.target || !/textarea|select/i.test(p.target.nodeName) && "text" !== p.target.type) {
					var q = "keypress" !== p.type && a.hotkeys.specialKeys[p.which],
						m = String.fromCharCode(p.which).toLowerCase(),
						j = "",
						n = {};
					p.altKey && "alt" !== q && (j += "alt+"), p.ctrlKey && "ctrl" !== q && (j += "ctrl+"), p.metaKey && !p.ctrlKey && "meta" !== q && (j += "meta+"), p.shiftKey && "shift" !== q && (j += "shift+"), q ? n[j + q] = !0 : (n[j + m] = !0, n[j + a.hotkeys.shiftNums[m]] = !0, "shift+" === j && (n[a.hotkeys.shiftNums[m]] = !0));
					for (var g = 0, k = f.length; g < k; g++) {
						if (n[f[g]]) {
							return c.apply(this, arguments)
						}
					}
				}
			}
		}
	}
	a.hotkeys = {
		version: "0.8",
		specialKeys: {
			8: "backspace",
			9: "tab",
			13: "return",
			16: "shift",
			17: "ctrl",
			18: "alt",
			19: "pause",
			20: "capslock",
			27: "esc",
			32: "space",
			33: "pageup",
			34: "pagedown",
			35: "end",
			36: "home",
			37: "left",
			38: "up",
			39: "right",
			40: "down",
			45: "insert",
			46: "del",
			96: "0",
			97: "1",
			98: "2",
			99: "3",
			100: "4",
			101: "5",
			102: "6",
			103: "7",
			104: "8",
			105: "9",
			106: "*",
			107: "+",
			109: "-",
			110: ".",
			111: "/",
			112: "f1",
			113: "f2",
			114: "f3",
			115: "f4",
			116: "f5",
			117: "f6",
			118: "f7",
			119: "f8",
			120: "f9",
			121: "f10",
			122: "f11",
			123: "f12",
			144: "numlock",
			145: "scroll",
			191: "/",
			224: "meta"
		},
		shiftNums: {
			"`": "~",
			1: "!",
			2: "@",
			3: "#",
			4: "$",
			5: "%",
			6: "^",
			7: "&",
			8: "*",
			9: "(",
			0: ")",
			"-": "_",
			"=": "+",
			";": ": ",
			"'": '"',
			",": "<",
			".": ">",
			"/": "?",
			"\\": "|"
		}
	}, a.each(["keydown", "keyup", "keypress"], function() {
		a.event.special[this] = {
			add: b
		}
	})
}(jQuery), function(d) {
	function g(k, a) {
		if (k === !1) {
			return k
		}
		if (!k) {
			return a
		}
		k === !0 ? k = {
			add: !0,
			"delete": !0,
			edit: !0,
			sort: !0
		} : "string" == typeof k && (k = k.split(","));
		var l;
		return d.isArray(k) && (l = {}, d.each(k, function(n, m) {
			d.isPlainObject(m) ? l[m.action] = m : l[m] = !0
		}), k = l), d.isPlainObject(k) && (l = {}, d.each(k, function(n, m) {
			m ? l[n] = d.extend({
				type: n
			}, b[n], d.isPlainObject(m) ? m : null) : l[n] = !1
		}), k = l), a ? d.extend(!0, {}, a, k) : k
	}
	function c(k, a, l) {
		return a = a || k.type, d(l || k.template).addClass("tree-action").attr(d.extend({
			"data-type": a,
			title: k.title || ""
		}, k.attr)).data("action", k)
	}
	var j = "zui.tree",
		h = 0,
		f = function(k, a) {
			this.name = j, this.$ = d(k), this.getOptions(a), this._init()
		},
		b = {
			sort: {
				template: '<a class="sort-handler" href="javascript:;"><i class="icon icon-move"></i></a>'
			},
			add: {
				template: '<a href="javascript:;"><i class="icon icon-plus"></i></a>'
			},
			edit: {
				template: '<a href="javascript:;"><i class="icon icon-pencil"></i></a>'
			},
			"delete": {
				template: '<a href="javascript:;"><i class="icon icon-trash"></i></a>'
			}
		};
	f.DEFAULTS = {
		animate: null,
		initialState: "normal",
		toggleTemplate: '<i class="list-toggle icon"></i>'
	}, f.prototype.add = function(w, u, p, m, y) {
		var x, k = d(w),
			q = this.options;
		if (k.is("li") ? (x = k.children("ul"), x.length || (x = d("<ul/>"), k.append(x), this._initList(x, k))) : x = k, x) {
			var v = this;
			d.isArray(u) || (u = [u]), d.each(u, function(r, a) {
				var A = d("<li/>").data(a).appendTo(x);
				void 0 !== a.id && A.attr("data-id", a.id);
				var z = q.itemWrapper ? d(q.itemWrapper === !0 ? '<div class="tree-item-wrapper"/>' : q.itemWrapper).appendTo(A) : A;
				if (a.html) {
					z.html(a.html)
				} else {
					if (d.isFunction(v.options.itemCreator)) {
						var l = v.options.itemCreator(A, a);
						l !== !0 && l !== !1 && z.html(l)
					} else {
						a.url ? z.append(d("<a/>", {
							href: a.url
						}).text(a.title || a.name)) : z.append(d("<span/>").text(a.title || a.name))
					}
				}
				v._initItem(A, a.idx || r, x, a), a.children && a.children.length && v.add(A, a.children)
			}), this._initList(x), p && !x.hasClass("tree") && v.expand(x.parent("li"), m, y)
		}
	}, f.prototype.reload = function(k) {
		var a = this;
		k && (a.$.empty(), a.add(a.$, k)), a.isPreserve && a.store.time && a.$.find("li:not(.tree-action-item)").each(function() {
			var l = d(this);
			a[a.store[l.data("id")] ? "expand" : "collapse"](l, !0, !0)
		})
	}, f.prototype._initList = function(w, u, p, k) {
		var q = this;
		w.hasClass("tree") ? (p = 0, u = null) : (u = (u || w.closest("li")).addClass("has-list"), u.find(".list-toggle").length || u.prepend(this.options.toggleTemplate), p = p || u.data("idx")), w.removeClass("has-active-item");
		var e = w.attr("data-idx", p || 0).children("li:not(.tree-action-item)").each(function(a) {
			q._initItem(d(this), a + 1, w)
		});
		1 !== e.length || e.find("ul").length || e.addClass("tree-single-item"), k = k || (u ? u.data() : null);
		var m = g(k ? k.actions : null, this.actions);
		if (m) {
			if (m.add && m.add.templateInList !== !1) {
				var v = w.children("li.tree-action-item");
				v.length ? v.detach().appendTo(w) : d('<li class="tree-action-item"/>').append(c(m.add, "add", m.add.templateInList)).appendTo(w)
			}
			m.sort && w.sortable(d.extend({
				dragCssClass: "tree-drag-holder",
				trigger: ".sort-handler",
				selector: "li:not(.tree-action-item)",
				finish: function(a) {
					q.callEvent("action", {
						action: m.sort,
						$list: w,
						target: a.target,
						item: k
					})
				}
			}, m.sort.options, d.isPlainObject(this.options.sortable) ? this.options.sortable : null))
		}
		u && (u.hasClass("open") || k && k.open) && u.addClass("open in")
	}, f.prototype._initItem = function(m, k, x, w) {
		if (void 0 === k) {
			var e = m.prev("li");
			k = e.length ? e.data("idx") + 1 : 1
		}
		if (x = x || m.closest("ul"), m.attr("data-idx", k).removeClass("tree-single-item"), !m.data("id")) {
			var p = k;
			x.hasClass("tree") || (p = x.parent("li").data("id") + "-" + p), m.attr("data-id", p)
		}
		m.hasClass("active") && x.parent("li").addClass("has-active-item"), w = w || m.data();
		var q = g(w.actions, this.actions);
		if (q) {
			var v = m.find(".tree-actions");
			v.length || (v = d('<div class="tree-actions"/>').appendTo(this.options.itemWrapper ? m.find(".tree-item-wrapper") : m), d.each(q, function(a, l) {
				l && v.append(c(l, a))
			}))
		}
		var u = m.children("ul");
		u.length && this._initList(u, m, k, w)
	}, f.prototype._init = function() {
		var k = this.options,
			l = this;
		this.actions = g(k.actions), this.$.addClass("tree"), k.animate && this.$.addClass("tree-animate"), this._initList(this.$);
		var e = k.initialState,
			m = d.zui && d.zui.store && d.zui.store.enable;
		m && (this.selector = j + "::" + (k.name || "") + "#" + (this.$.attr("id") || h++), this.store = d.zui.store[k.name ? "get" : "pageGet"](this.selector, {})), "preserve" === e && (m ? this.isPreserve = !0 : this.options.initialState = e = "normal"), this.reload(k.data), m && (this.isPreserve = !0), "expand" === e ? this.expand() : "collapse" === e && this.collapse(), this.$.on("click", '.list-toggle,a[href="#"],.tree-toggle', function(o) {
			var a = d(this),
				p = a.parent("li");
			l.callEvent("hit", {
				target: p,
				item: p.data()
			}), l.toggle(p), a.is("a") && o.preventDefault()
		}).on("click", ".tree-action", function() {
			var o = d(this),
				a = o.data();
			if (a.action && (a = a.action), "sort" !== a.type) {
				var p = o.closest("li:not(.tree-action-item)");
				l.callEvent("action", {
					action: a,
					target: this,
					$item: p,
					item: p.data()
				})
			}
		})
	}, f.prototype.preserve = function(k, a, m) {
		if (this.isPreserve) {
			if (k) {
				a = a || k.data("id"), m = void 0 === m && k.hasClass("open"), m ? this.store[a] = m : delete this.store[a], this.store.time = (new Date).getTime(), d.zui.store[this.options.name ? "set" : "pageSet"](this.selector, this.store)
			} else {
				var l = this;
				this.store = {}, this.$.find("li").each(function() {
					l.preserve(d(this))
				})
			}
		}
	}, f.prototype.expand = function(k, l, a) {
		k ? (k.addClass("open"), !l && this.options.animate ? setTimeout(function() {
			k.addClass("in")
		}, 10) : k.addClass("in")) : k = this.$.find("li.has-list").addClass("open in"), a || this.preserve(k), this.callEvent("expand", k, this)
	}, f.prototype.show = function(k, a, m) {
		var l = this;
		k.each(function() {
			var p = d(this);
			if (l.expand(p, a, m), p) {
				for (var o = p.parent("ul"); o && o.length && !o.hasClass("tree");) {
					var n = o.parent("li");
					n.length ? (l.expand(n, a, m), o = n.parent("ul")) : o = !1
				}
			}
		})
	}, f.prototype.collapse = function(k, l, a) {
		k ? !l && this.options.animate ? (k.removeClass("in"), setTimeout(function() {
			k.removeClass("open")
		}, 300)) : k.removeClass("open in") : k = this.$.find("li.has-list").removeClass("open in"), a || this.preserve(k), this.callEvent("collapse", k, this)
	}, f.prototype.toggle = function(a) {
		var k = a && a.hasClass("open") || a === !1 || void 0 === a && this.$.find("li.has-list.open").length;
		this[k ? "collapse" : "expand"](a)
	}, f.prototype.getOptions = function(a) {
		this.options = d.extend({}, f.DEFAULTS, this.$.data(), a), null === this.options.animate && this.$.hasClass("tree-animate") && (this.options.animate = !0)
	}, f.prototype.toData = function(k, a) {
		d.isFunction(k) && (a = k, k = null), k = k || this.$;
		var l = this;
		return k.children("li:not(.tree-action-item)").map(function() {
			var n = d(this),
				p = n.data();
			delete p["zui.droppable"];
			var m = n.children("ul");
			return m.length && (p.children = l.toData(m)), d.isFunction(a) ? a(p, n) : p
		}).get()
	}, f.prototype.callEvent = function(k, a) {
		var l;
		return d.isFunction(this.options[k]) && (l = this.options[k](a, this)), this.$.trigger(d.Event(k + "." + this.name, a)), l
	}, d.fn.tree = function(k, a) {
		return this.each(function() {
			var m = d(this),
				e = m.data(j),
				l = "object" == typeof k && k;
			e || m.data(j, e = new f(this, l)), "string" == typeof k && e[k](a)
		})
	}, d.fn.tree.Constructor = f, d(function() {
		d('[data-ride="tree"]').tree()
	})
}(jQuery), function(b) {
	var c = "zui.colorPicker",
		a = '<div class="colorpicker"><button type="button" class="btn dropdown-toggle" data-toggle="dropdown"><span class="cp-title"></span><i class="ic"></i></button><ul class="dropdown-menu clearfix"></ul></div>',
		f = {
			zh_cn: {
				errorTip: "不是有效的颜色值"
			},
			zh_tw: {
				errorTip: "不是有效的顏色值"
			},
			en: {
				errorTip: "Not a valid color value"
			}
		},
		d = function(e, g) {
			this.name = c, this.$ = b(e), this.getOptions(g), this.init()
		};
	d.DEFAULTS = {
		colors: ["#00BCD4", "#388E3C", "#3280fc", "#3F51B5", "#9C27B0", "#795548", "#F57C00", "#F44336", "#E91E63"],
		pullMenuRight: !0,
		wrapper: "btn-wrapper",
		tileSize: 30,
		lineCount: 5,
		optional: !0,
		tooltip: "top",
		icon: "caret-down",
		updateBtn: "auto"
	}, d.prototype.init = function() {
		var k = this,
			m = k.options,
			l = k.$,
			h = l.parent(),
			g = !1;
		h.hasClass("colorpicker") ? k.$picker = h : (k.$picker = b(m.template || a), g = !0), k.$picker.addClass(m.wrapper).find(".cp-title").toggle(void 0 !== m.title).text(m.title), k.$menu = k.$picker.find(".dropdown-menu").toggleClass("pull-right", m.pullMenuRight), k.$btn = k.$picker.find(".btn.dropdown-toggle"), k.$btn.find(".ic").addClass("icon-" + m.icon), m.btnTip && k.$picker.attr("data-toggle", "tooltip").tooltip({
			title: m.btnTip,
			placement: m.tooltip,
			container: "body"
		}), l.attr("data-provide", null), g && l.after(k.$picker), k.colors = {}, b.each(m.colors, function(e, q) {
			if (b.zui.Color.isColor(q)) {
				var p = new b.zui.Color(q);
				k.colors[p.toCssStr()] = p
			}
		}), k.updateColors(), k.$picker.on("click", ".cp-tile", function() {
			k.setValue(b(this).data("color"))
		});
		var j = function() {
				var e = l.val(),
					n = b.zui.Color.isColor(e);
				l.parent().toggleClass("has-error", !(n || m.optional && "" === e)), n ? k.setValue(e, !0) : m.optional && "" === e ? l.tooltip("hide") : l.is(":focus") || l.tooltip("show", m.errorTip)
			};
		l.is("input:not([type=hidden])") ? (m.tooltip && l.attr("data-toggle", "tooltip").tooltip({
			trigger: "manual",
			placement: m.tooltip,
			tipClass: "tooltip-danger",
			container: "body"
		}), l.on("keyup paste input change", j)) : l.appendTo(k.$picker), j()
	}, d.prototype.addColor = function(h) {
		h instanceof b.zui.Color || (h = new b.zui.Color(h));
		var g = h.toCssStr(),
			k = this.options;
		this.colors[g] || (this.colors[g] = h);
		var j = b('<a href="###" class="cp-tile"></a>', {
			titile: h
		}).data("color", h).css({
			color: h.contrast().toCssStr(),
			background: g,
			"border-color": h.luma() > 0.43 ? "#ccc" : "transparent"
		}).attr("data-color", g);
		this.$menu.append(b("<li/>").css({
			width: k.tileSize,
			height: k.tileSize
		}).append(j)), k.optional && this.$menu.find(".cp-tile.empty").parent().detach().appendTo(this.$menu)
	}, d.prototype.updateColors = function(k) {
		var h = (this.$picker, this.$menu.children("li:not(.heading)").remove()),
			m = this.options,
			k = k || this.colors,
			l = this,
			j = 0;
		if (b.each(k, function(n, o) {
			l.addColor(o), j++
		}), m.optional) {
			var g = b('<li><a class="cp-tile empty" href="###"></a></li>').css({
				width: m.tileSize,
				height: m.tileSize
			});
			this.$menu.append(g), j++
		}
		h.css("width", Math.min(j, m.lineCount) * m.tileSize + 6)
	}, d.prototype.setValue = function(u, p) {
		var k = this,
			j = k.options,
			w = k.$btn;
		k.$menu.find(".cp-tile.active").removeClass("active");
		var v = "",
			g = j.updateBtn;
		if ("auto" === g) {
			var m = w.find(".color-bar");
			g = !m.length ||
			function(e) {
				m.css("background", e || "")
			}
		}
		if (u) {
			var q = new b.zui.Color(u);
			v = q.toCssStr().toLowerCase(), g && (b.isFunction(g) ? g(v, w, k) : w.css({
				background: v,
				color: q.contrast().toCssStr(),
				borderColor: q.luma() > 0.43 ? "#ccc" : v
			})), k.colors[v] || k.addColor(q), p || k.$.val().toLowerCase() === v || k.$.val(v).trigger("change"), k.$menu.find('.cp-tile[data-color="' + v + '"]').addClass("active"), k.$.tooltip("hide"), k.$.trigger("colorchange", q)
		} else {
			g && (b.isFunction(g) ? g(null, w, k) : w.attr("style", null)), p || "" === k.$.val() || k.$.val(v).trigger("change"), j.optional && k.$.tooltip("hide"), k.$menu.find(".cp-tile.empty").addClass("active"), k.$.trigger("colorchange", null)
		}
		j.updateBorder && b(j.updateBorder).css("border-color", v), j.updateBackground && b(j.updateBackground).css("background-color", v), j.updateColor && b(j.updateColor).css("color", v), j.updateText && b(j.updateText).text(v)
	}, d.prototype.getOptions = function(j) {
		var g = b.extend({}, d.DEFAULTS, this.$.data(), j);
		"string" == typeof g.colors && (g.colors = g.colors.split(","));
		var h = (g.lang || b.zui.clientLang()).toLowerCase();
		g.errorTip || (g.errorTip = f[h].errorTip), b.fn.tooltip || (g.btnTip = !1), this.options = g
	}, b.fn.colorPicker = function(e) {
		return this.each(function() {
			var j = b(this),
				h = j.data(c),
				g = "object" == typeof e && e;
			h || j.data(c, h = new d(this, g)), "string" == typeof e && h[e]()
		})
	}, b.fn.colorPicker.Constructor = d, b(function() {
		b('[data-provide="colorpicker"]').colorPicker()
	})
}(jQuery), function(b, d) {
	function a(e) {
		return e === d && (e = g += 1), f[e % f.length]
	}
	var g = 0,
		f = ["#00a9fc", "#ff5d5d", "#fdc137", "#00da88", "#7ec5ff", "#8666b8", "#bd7b46", "#ff9100", "#ff3d00", "#f57f17", "#00e5ff", "#00b0ff", "#2979ff", "#3d5afe", "#651fff", "#d500f9", "#f50057", "#ff1744"];
	jQuery.fn.tableChart = function() {
		b(this).each(function() {
			var u = b(this),
				m = u.data(),
				k = m.chart || "pie",
				y = b(m.target);
			if (y.length) {
				var x = null;
				if ("pie" === k) {
					m = b.extend({
						scaleShowLabels: !0,
						scaleLabel: "<%=label%>: <%=value%>"
					}, m);
					var j = [],
						p = u.find("tbody > tr").each(function(h) {
							var r = b(this),
								l = a();
							r.attr("data-id", h).find(".chart-color-dot").css("background", l), j.push({
								label: r.find(".chart-label").text(),
								value: parseInt(r.find(".chart-value").text()),
								color: l,
								id: h
							})
						});
					j.length > 1 ? m.scaleLabelPlacement = "outside" : 1 === j.length && (m.scaleLabelPlacement = "inside", j.push({
						label: "",
						value: j[0].value / 2000,
						color: "#fff",
						showLabel: !1
					})), x = y.pieChart(j, m), y.on("mousemove", function(h) {
						var l = x.getSegmentsAtEvent(h);
						p.removeClass("active"), l.length && p.filter('[data-id="' + l[0].id + '"]').addClass("active")
					})
				} else {
					if ("bar" === k) {
						var q = a(),
							w = [],
							v = {
								label: u.find("thead .chart-label").text(),
								color: q,
								data: []
							},
							p = u.find("tbody > tr").each(function(l) {
								var h = b(this);
								w.push(h.find(".chart-label").text()), v.data.push(parseInt(h.find(".chart-value").text())), h.find(".chart-color-dot").css("background", q)
							}),
							j = {
								labels: w,
								datasets: [v]
							};
						w.length && (m.barValueSpacing = 5), x = y.barChart(j, m)
					} else {
						if ("line" === k) {
							var q = a(),
								w = [],
								v = {
									label: u.find("thead .chart-label").text(),
									color: q,
									data: []
								},
								p = u.find("tbody > tr").each(function(l) {
									var h = b(this);
									w.push(h.find(".chart-label").text()), v.data.push(parseInt(h.find(".chart-value").text())), h.find(".chart-color-dot").css("background", q)
								}),
								j = {
									labels: w,
									datasets: [v]
								};
							w.length && (m.barValueSpacing = 5), x = y.lineChart(j, m)
						}
					}
				}
				null !== x && u.data("zui.chart", x)
			}
		})
	}, b(".table-chart").tableChart();
	var c = function(j, q) {
			var p = b(j);

			if (!p.data("initProgressPie")) {
				p.data("initProgressPie", 1);
				var k = p.is("canvas") ? p : p.find("canvas"),
					h = b.extend({
						value: 0,
						color: p.attr('data-color') || b.getThemeColor("primary") || "#006af1",
						backColor: b.getThemeColor("pale") || "#E9F2FB",
						doughnut: !0,
						doughnutSize: 85,
						width: 20,
						height: 20,
						showTip: !1,
						name: "",
						tipTemplate: "<%=value%>%",
						animation: "auto",
						realValue: parseFloat(p.find(".progress-value").text())
					}, q, p.data()),
					m = k.length;
				m || (k = b("<canvas>").appendTo(p)), k.attr("width") !== d ? h.width = k.attr("width") : k.attr("width", h.width), k.attr("height") !== d ? h.height = k.attr("height") : k.attr("height", h.height), m || 8 != b.zui.browser.ie || G_vmlCanvasManager.initElement(k[0]), "auto" === h.animation && (h.animation = h.width > 30), h.value = Math.max(0, Math.min(100, h.value)), p.addClass("progress-pie-" + h.width);
				var e = [{
					value: h.value,
					label: h.name,
					color: h.color,
					circleBeginEnd: !0
				}, {
					value: 100 - h.value,
					label: "",
					color: h.backColor
				}];
				k[h.doughnut ? "doughnutChart" : "pieChart"](e, b.extend({
					segmentShowStroke: !1,
					animation: h.animation,
					showTooltips: h.showTip,
					tooltipTemplate: h.tipTemplate,
					percentageInnerCutout: h.doughnutSize,
					reverseDrawOrder: !0,
					animationEasing: "easeInOutQuart",
					onAnimationProgress: h.realValue ?
					function(l) {
						p.find(".progress-value").text(Math.floor(h.realValue * l))
					} : d,
					onAnimationComplete: h.realValue ?
					function(l) {
						p.find(".progress-value").text(h.realValue)
					} : d
				}, h.chartOptions))
			}
		};
	jQuery.fn.progressPie = function(h) {
		b(this).each(function() {
			var e = b(this);
			if (!e.closest(".hidden").length) {
				var j = e.closest(".tab-pane");
				j.length && !j.hasClass("active") ? b('[data-toggle="tab"][data-target="#' + j.attr("id") + '"]').one("shown.zui.tab", function() {
					c(e, h)
				}) : c(this, h)
			}
		})
	}, b(function() {
		b(".table-chart").tableChart();
		var h = b(".progress-pie");
		h.length > 100 ? setTimeout(function() {
			h.progressPie()
		}, 1000) : h.progressPie()
	})
}(jQuery, void 0), function(a) {
	jQuery.fn.sparkline = function(b) {
		a(this).each(function() {
			var v = a(this),
				m = a.extend({
					values: v.attr("values"),
					width: v.width() - 4,
					height: v.height() - 4
				}, v.data(), b),
				k = m.height,
				D = [],
				B = m.width,
				e = m.values.split(","),
				q = 0;
			for (var w in e) {
				var A = parseFloat(e[w]);
				NaN != A && (D.push(A), q = Math.max(A, q))
			}
			var z = (Math.min(q, 30), Math.min(B, Math.max(10, D.length * B / 30))),
				C = v.children("canvas");
			C.length || (v.append('<canvas class="projectline-canvas"></canvas>'), C = v.children("canvas")), C.attr("width", z).attr("height", k);
			var j = {
				labels: D,
				datasets: [{
					fillColor: a.getThemeColor("pale") || "rgba(0,0,255,0.05)",
					strokeColor: a.getThemeColor("primary") || "#0054EC",
					pointColor: a.getThemeColor("secondary") || "rgba(255,136,0,1)",
					pointStrokeColor: "#fff",
					data: D
				}]
			},
				y = {
					animation: !0,
					scaleOverride: !0,
					scaleStepWidth: Math.ceil(q / 10),
					scaleSteps: 10,
					scaleStartValue: 0,
					showScale: !1,
					showTooltips: !1,
					pointDot: !1,
					scaleShowGridLines: !1,
					datasetStrokeWidth: 1
				},
				x = a(C).lineChart(j, y);
			v.data("sparklineChart", x)
		})
	}, a(function() {
		a(".sparkline").sparkline()
	})
}(jQuery), function(a) {
	a(function() {
		a.fn.fixedDate = function() {
			return a(this).each(function() {
				var f = a(this).attr("autocomplete", "off");
				"0000-00-00" == f.val() && f.focus(function() {
					"0000-00-00" == f.val() && f.val("").datetimepicker("update")
				}).blur(function() {
					"" == f.val() && f.val("0000-00-00")
				})
			})
		};
		var c = {
			language: a("html").attr("lang"),
			weekStart: 1,
			todayBtn: 1,
			autoclose: 1,
			todayHighlight: 1,
			startView: 2,
			forceParse: 0,
			showMeridian: 1,
			format: "yyyy-mm-dd hh:ii",
			startDate: "1970-1-1"
		},
			b = a.extend({}, c, {
				minView: 2,
				format: "yyyy-mm-dd"
			}),
			d = a.extend({}, c, {
				startView: 1,
				minView: 0,
				maxView: 1,
				format: "hh:ii"
			});
		a(".datepicker-wrapper").click(function() {
			a(this).find(".form-date, .form-datetime, .form-time").datetimepicker("show").focus()
		}), window.datepickerOptions = c, a.fn.datepicker = function(f) {
			return this.datetimepicker(a.extend({}, b, f))
		}, a.fn.timepicker = function(f) {
			return this.datetimepicker(a.extend({}, d, f))
		}, a.fn.datepickerAll = function() {
			return this.find(".form-datetime").fixedDate().datetimepicker(c), this.find(".form-date").fixedDate().datepicker(), this.find(".form-time").fixedDate().timepicker(), this
		}, a("body").datepickerAll()
	})
}(jQuery), function(a) {
	var b = function(w, q) {
			q = a.extend({
				idStart: 0,
				idEnd: 9,
				chosen: !0,
				datetimepicker: !0,
				colorPicker: !0,
				hotkeys: !0
			}, q, w.data());
			var k = w.find(".template");
			!k.length && q.template && (k = a(q.template));
			var j = 0,
				B = 0,
				z = function(c) {
					c.is("select.chosen") ? c.next(".chosen-container").find("input").focus() : c.focus()
				},
				f = function(d) {
					var c = w.find("[data-ctrl-index]:focus,.chosen-container-active").first();
					if (c.length) {
						if (c.is(".chosen-container-active")) {
							if (c.hasClass("chosen-with-drop") && ("down" === d || "up" === d)) {
								return
							}
							c = c.prev("select.chosen")
						}
						var h = c.data("ctrlIndex"),
							e = c.closest("tr").data("row");
						"down" === d ? e < B - 1 ? e += 1 : e = 0 : "up" === d ? e > 0 ? e -= 1 : e = B - 1 : "left" === d ? h > 0 ? h -= 1 : h = j - 1 : "right" === d && (h < j - 1 ? h += 1 : h = 0), z(w.find('tr[data-row="' + e + '"]').find('[data-ctrl-index="' + h + '"]'))
					}
				},
				m = {
					options: q,
					focusNext: f,
					focusControl: z
				},
				v = w.find("tbody,.batch-rows"),
				y = function(c) {
					a.fn.chosen && q.chosen && c.find(".chosen").chosen(a.isPlainObject(q.chosen) ? q.chosen : null), a.fn.datetimepicker && q.datetimepicker && c.datepickerAll(a.isPlainObject(q.datetimepicker) ? q.datetimepicker : null), a.fn.colorPicker && q.colorPicker && c.find("input.colorpicker").colorPicker(a.isPlainObject(q.colorPicker) ? q.colorPicker : null);
					var d = 0;
					c.find('input[type!="hidden"],textarea,select').each(function() {
						var h = a(this);
						h.parent().hasClass("chosen-search") || h.attr("data-ctrl-index", d++)
					}), j = Math.max(j, d)
				};
			if (k.length) {
				var x = k.remove().html(),
					A = function(h, l) {
						var c = x;
						"number" != typeof h && (h = B), B = Math.max(h + 1, B), c = c.replace(/\$idPlus/g, h + 1).replace(/\$id/g, h);
						var d = a("<" + k[0].tagName.toLowerCase() + " />").html(c);
						return d.attr("data-row", h).addClass(k.attr("class")).removeClass("template"), q.rowCreator && q.rowCreator(d, h, q), l ? l.after(d) : v.append(d), y(d), d
					};
				a.extend(m, {
					createRow: A,
					template: x
				});
				for (var g = q.idStart; g <= q.idEnd; ++g) {
					A(g)
				}
			} else {
				y(w)
			}
			w.on("click", ".btn-copy", function() {
				var d = a(this),
					c = a(d.data("copyFrom")).val(),
					h = a(d.data("copyTo")).val(c).addClass("highlight");
				setTimeout(function() {
					h.removeClass("highlight")
				}, 2000)
			}), q.hotkeys && a(document).on("keydown", function(d) {
				var h = {
					"Ctrl+#37": "left",
					"Ctrl+#39": "right",
					"#38": "up",
					"#40": "down",
					"Ctrl+#38": "up",
					"Ctrl+#40": "down"
				},
					c = [];
				d.ctrlKey && c.push("Ctrl"), c.push("#" + d.keyCode);
				var l = h[c.join("+")];
				l && (f(l), d.ctrlKey && (d.stopPropagation(), d.preventDefault()))
			}), w.data("zui.batchActionForm", m)
		};
	a.fn.batchActionForm = function(c) {
		return this.each(function() {
			b(a(this), c)
		})
	}
}(jQuery), function(b, d) {
	var a = "zui.table",
		g = {
			zh_cn: {
				selectedItems: "已选择 <strong>{0}</strong> 项",
				attrTotal: "{0}总计 <strong>{1}</strong>"
			},
			zh_tw: {
				selectedItems: "已选择 <strong>{0}</strong> 项",
				attrTotal: "{0}总计 <strong>{1}</strong>"
			},
			en: {
				selectedItems: "Seleted <strong>{0}</strong> items",
				attrTotal: "{0} total <strong>{1}</strong>"
			}
		},
		f = /^((?!chrome|android).)*safari/i.test(navigator.userAgent),
		c = function(p, q) {
			var k = this;
			k.name = a;
			var n = k.$ = b(p);
			q = k.options = b.extend({}, c.DEFAULTS, this.$.data(), q);
			var j = q.lang || "zh_cn";
			k.lang = b.isPlainObject(j) ? b.extend(!0, {}, g[j.lang || b.zui.clientLang()], j) : g[j], n.attr("id") || (n.attr("id", "table-" + b.zui.uuid()), q.hot && console.warn("ZUI: table hot replace id not defined, the element id attribute should be set.")), n.attr("data-ride") || n.attr("data-ride", "table"), k.getTable().find("thead>tr>th").each(function() {
				var l = b(this);
				if (!l.attr("title")) {
					var h = b.trim(l.find("a").text() || l.text() || "");
					h.length && l.attr("title", h)
				}
			}), q.checkable && (n.on("click", ".check-all", function() {
				k.checkAll(!b(this).hasClass("checked"))
			}).on("click", "tbody>tr", function(h) {
				b(h.target).closest('.btn,a,.not-check,.form-control,input[type="text"],.chosen-container').length || k.checkRow(b(this))
			}).on("click", 'tbody input[type="checkbox"],tbody label[for]', function(l) {
				l.stopPropagation();
				var h = b(this);
				h.is("label") && (h = h.closest(".checkbox-primary").find('input[type="checkbox"]')), k.checkRow(h.closest("tr"), h.is(":checked"))
			}), q.selectable && n.selectable(b.extend({}, {
				selector: k.isDataTable ? ".fixed-left tbody>tr" : "tbody>tr",
				selectClass: "",
				trigger: ".c-id",
				clickBehavior: "multi",
				listenClick: !1,
				select: function(h) {
					k.checkRow(h.target, !0), b.cookie("ajax_dragSelected") || (b.cookie("ajax_dragSelected", "on", {
						expires: config.cookieLife,
						path: config.webRoot
					}), b.ajaxSendScore("dragSelected"))
				},
				unselect: function(e) {
					k.checkRow(e.target, !1)
				},
				rangeStyle: {
					border: "1px solid #006af1",
					backgroundColor: "rgba(50,128,252,0.2)",
					borderRadius: "2px"
				}
			}, b.isPlainObject(q.selectable) ? q.selectable : null)));
			var m = k.$form = n.is("form") ? n : n.find("form");
			m.length && (q.ajaxForm ? m.ajaxForm(b.isPlainObject(q.ajaxForm) ? q.ajaxForm : null) : m.on("click", "[data-form-action]", function() {
				m.attr("action", b(this).data("formAction")).submit()
			})), (q.fixFooter || q.fixHeader) && (k.pageFooterHeight = b("#footer").outerHeight(), k.updateFixUI(), b(window).on("scroll resize", function() {
				k.updateFixUI()
			}).on("sidebar.toggle", function() {
				setTimeout(function() {
					k.updateFixUI()
				}, 200)
			})), q.group && (n.on("click", ".group-toggle", function() {
				k.toggleRowGroup(b(this).closest("tr").data("id"))
			}), b(document).on("click", ".group-collapse-all", function() {
				k.toggleGroups(!1)
			}).on("click", ".group-expand-all", function() {
				k.toggleGroups(!0)
			})), k.defaultStatistic = n.find(".table-statistic").html(), k.updateStatistic(), k.initModals(), k.checkItems = {}, k.updateCheckUI()
		};
	c.prototype.reload = function(k) {
		var h = this,
			m = h.options,
			l = m.replaceId;
		if (!l) {
			return k && k()
		}
		"self" === l && (l = h.$.attr("id"));
		var j = b("<div></div>");
		h.$.addClass("load-indicator loading"), j.load(window.location.href + " #" + l, function() {
			h.$.empty().html(j.children().html()).removeClass("load-indicator loading"), h.$.trigger("beforeTableReload"), h.updateStatistic(), h.initModals(), h.$.datepickerAll();
			var p = h.$.find("tbody>tr"),
				e = !1;
			b.each(h.checkItems, function(o, q) {
				q && (h.checkRow(p.filter('[data-id="' + o + '"]'), !0, !0), e = !0)
			}), e && h.updateCheckUI(), h.$.trigger("tableReload");
			var n = b("#mainMenu>.btn-toolbar>.btn-active-text>.label");
			n.length && n.text(h.getTable().find("tbody:first>tr:not(.table-children)").length), h.$.find('[data-ride="pager"]').pager(), k && k(), m.afterReload && m.afterReload()
		})
	}, c.prototype.initModals = function() {
		var j = this,
			h = j.options,
			l = j.$.find(h.iframeModalTrigger);
		if (l.length) {
			var k = {
				type: "iframe",
				onHide: h.replaceId ?
				function() {
					var e = b.cookie("selfClose");
					(1 == e || h.hot) && (b("#triggerModal").data("cancel-reload", 1), j.reload(function() {
						b.cookie("selfClose", 0)
					}))
				} : null
			};
			l.modalTrigger(k)
		}
	}, c.prototype.getTable = function() {
		var h = this.$;
		if (this.isDataTable) {
			return h.find("div.datatable")
		}
		var j = h.is("table") ? h : h.find("table:not(.fixed-header-copy)").first();
		return j.is(".datatable") && (this.isDataTable = !0, j = h.find("div.datatable")), j
	}, c.prototype.toggleGroups = function(j) {
		var h = this,
			k = {};
		h.$.find("tbody>tr").each(function() {
			var e = b(this).closest("tr").data("id");
			k[e] || h.toggleRowGroup(e, j)
		})
	}, c.prototype.toggleRowGroup = function(h, l) {
		var k = this.$.find('tbody>tr[data-id="' + h + '"]'),
			j = k.filter(".group-summary"),
			e = l === d ? !j.hasClass("hidden") : !! l;
		k.not(".group-summary").toggleClass("hidden", !e), j.toggleClass("hidden", e), b("body").toggleClass("table-group-collapsed", !this.$.find("tbody>tr.group-summary.hidden").length)
	}, c.prototype.updateStatistic = function() {
		var h = this,
			m = h.$.find(".table-statistic");
		if (m.length) {
			if (h.defaultStatistic === d && (h.defaultStatistic = m.html()), h.options.statisticCreator) {
				return void m.html(h.options.statisticCreator(h) || h.defaultStatistic)
			}
			var l = h.statisticCols;
			if (!l && l !== !1) {
				l = {};
				var j = !1;
				h.getTable().find("thead th").each(function(p) {
					var o = b(this),
						q = o.data("statistic");
					q && (j = !0, l[p] = {
						format: q,
						name: o.text()
					})
				}), h.statisticCols = !! j && l
			}
			var e = 0;
			l && b.each(l, function(n) {
				l[n].total = 0, l[n].checkedTotal = 0
			}), h.$.find(h.isDataTable ? ".fixed-left tbody>tr" : "tbody>tr").each(function() {
				var p = b(this),
					o = p.hasClass("checked"),
					q = p.children("td");
				o && e++, l && b.each(l, function(n) {
					var r = parseFloat(q.eq(n).text());
					isNaN(r) && (r = 0), l[n].total += r, o && (l[n].checkedTotal += r)
				})
			});
			var k = [];
			if (e) {
				k.push(h.lang.selectedItems.format(e))
			} else {
				if (h.defaultStatistic) {
					return void m.html(h.defaultStatistic)
				}
			}
			l && b.each(l, function(o) {
				var p = l[o],
					q = p[e ? "checkedTotal" : "total"];
				p.format && (q = p.format.format(q)), k.push(h.lang.attrTotal.format(p.name, q))
			}), m.html(k.join(", "))
		}
	}, c.prototype.updateFixUI = function(m) {
		var j = this,
			o = (new Date).getTime();
		if (!m && (j.lastUpdateCall && clearTimeout(j.lastUpdateCall), !j.lastUpdateTime || o - j.lastUpdateTime < 100)) {
			return void(j.lastUpdateCall = setTimeout(function() {
				j.updateFixUI(!0)
			}, 30))
		}
		if (j.lastUpdateTime = o, j.lastUpdateCall && (clearTimeout(j.lastUpdateCall), j.lastUpdateCall = null), f) {
			var k = j.getTable();
			if (k.parent().is(".table-responsive")) {
				var h = k.find("thead"),
					l = 0;
				h.find("th").each(function() {
					l += b(this).outerWidth()
				}), k.css("min-width", l)
			}
		}
		j.options.fixHeader && !j.isDataTable && j.fixHeader(), j.options.fixFooter && j.fixFooter()
	}, c.prototype.fixHeader = function() {
		var v = this,
			q = v.getTable(),
			m = q.find("thead"),
			k = m[0].getBoundingClientRect(),
			z = v.options.fixFooter,
			y = b.isFunction(z) ? z(k, m) : k.top < ("number" == typeof z ? z : -5),
			j = v.$.find(".fix-table-copy-wrapper"),
			p = q.parent(),
			u = p.is(".table-responsive");
		if (y) {
			if (j.length || (j = b('<div class="fix-table-copy-wrapper" style="overflow: hidden; position:fixed; z-index: 3; top: 0;"></div>').append(b('<table class="fixed-header-copy"></table>').addClass(q.attr("class")).append(m.clone())).insertAfter(q)), u) {
				var x = p[0].getBoundingClientRect();
				j.css({
					left: x.left,
					width: p.width()
				}), j.find(".fixed-header-copy").css({
					left: k.left - x.left,
					position: "relative",
					minWidth: q.width()
				}), p.data("fixHeaderScroll") || p.data("fixHeaderScroll", 1).on("scroll", function() {
					v.fixHeader()
				})
			} else {
				j.css({
					left: k.left,
					width: k.width
				})
			}
			var w = j.find("th");
			m.find("th").each(function(h) {
				w.eq(h).css("width", b(this).outerWidth())
			})
		} else {
			j.remove()
		}
	}, c.prototype.fixFooter = function() {
		var v, q = this,
			m = q.getTable(),
			k = q.$.find(".table-footer");
		if (q.isDataTable) {
			v = m[0].getBoundingClientRect()
		} else {
			var z = m.find("tbody");
			if (!z.length) {
				return
			}
			v = z[0].getBoundingClientRect(), v = z[0].getBoundingClientRect()
		}
		var y = q.options.fixFooter;
		k.toggleClass("fixed-footer", !! j);
		var j = b.isFunction(y) ? y(v, k) : v.bottom > window.innerHeight - 50 - ("number" == typeof y ? y : q.pageFooterHeight || 5);
		k.toggleClass("fixed-footer", !! j), m.toggleClass("with-footer-fixed", !! j), m.trigger("fixFooter", j);
		var p = b("body"),
			u = p.hasClass("body-modal");
		if (j) {
			var x = m.parent(),
				w = x.is(".table-responsive");
			k.css({
				bottom: q.pageFooterHeight || 0,
				left: w ? x[0].getBoundingClientRect().left : v.left,
				width: w ? x.width() : v.width
			}), u && p.css("padding-bottom", 40)
		} else {
			k.css({
				width: "",
				left: 0,
				bottom: 0
			}), u && p.css("padding-bottom", 0)
		}
	}, c.prototype.checkAll = function(j) {
		var h = this,
			k = h.$.find(h.isDataTable ? ".fixed-left tbody>tr" : "tbody>tr");
		k.each(function() {
			h.checkRow(b(this), j, !0)
		}), h.updateCheckUI()
	}, c.prototype.checkRow = function(h, e, l) {
		var k = this;
		k.isDataTable && !h.is(".datatable-row-left") && (h = k.getTable().find('.datatable-row-left[data-index="' + h.data("index") + '"]'));
		var j = h.find('input[type="checkbox"]');
		j.length && (e === d && (e = !j.is(":checked")), k.isDataTable ? k.getTable().find('.datatable-row[data-index="' + h.data("index") + '"]').toggleClass("checked", e) : h.toggleClass("checked", e), this.checkItems[h.data("id")] = e, j.prop("checked", e).trigger("change"), l || k.updateCheckUI())
	}, c.prototype.updateCheckUI = function() {
		var q = this,
			k = q.getTable(),
			v = k.find(q.isDataTable ? ".fixed-left tbody>tr" : "tbody>tr").not(".group-summary"),
			u = !1,
			m = null,
			j = 0,
			p = !1,
			h = v.length;
		v.each(function(r) {
			var e = b(this),
				o = e.find('input[type="checkbox"]');
			p = o.is(":checked");
			var l = q.isDataTable ? k.find('.datatable-row[data-index="' + e.data("index") + '"]') : e;
			l.toggleClass("checked", p), l.toggleClass("row-check-begin", p && !u), m && m.toggleClass("row-check-end", !p && u), p && (j += 1), m = l, u = p, h === r + 1 && l.toggleClass("row-check-end", p)
		}), q.$.toggleClass("has-row-checked", j > 0).find(".check-all").toggleClass("checked", !(!h || j !== h)), q.updateStatistic(), q.options.onCheckChange && q.options.onCheckChange()
	}, c.DEFAULTS = {
		checkable: !0,
		ajaxForm: !1,
		selectable: !0,
		fixHeader: !0,
		fixFooter: !0,
		iframeWidth: 900,
		replaceId: "self",
		hot: !1,
		iframeModalTrigger: ".iframe"
	}, b.fn.table = function(h) {
		return this.each(function() {
			var k = b(this),
				j = k.data(a),
				e = "object" == typeof h && h;
			j || k.data(a, j = new c(this, e)), "string" == typeof h && j[h]()
		})
	}, c.NAME = a, b.fn.table.Constructor = c, b(function() {
		b('[data-ride="table"]').table()
	})
}(jQuery, void 0), function(b, d, a) {
	b.fn._ajaxForm = b.fn.ajaxForm;
	var g = {
		timeout: d.config ? d.config.timeout : 0,
		dataType: "json",
		method: "post"
	},
		f = "";
	b.fn.enableForm = function(h, k, j) {
		return h === a && (h = !0), this.each(function() {
			var e = b(this);
			k || e.find('[type="submit"]').attr("disabled", h ? null : "disabled"), !j && e.hasClass("load-indicator") && e.toggleClass("loading", !h), e.toggleClass("form-disabled", !h)
		})
	}, b.enableForm = function(j, h, k) {
		j === !1 ? b("form").enableForm(j, h, k) : b("form.form-disabled").enableForm(!0, h, k)
	};
	var c = function(j, h, k) {
			"string" == typeof h && (k = h, h = null), k = k || "show", b.zui.messager ? b.zui.messager[k](j, h) : alert(j)
		};
	b.ajaxForm = function(j, n) {
		var e = b(j);
		if (e.length > 1) {
			return e.each(function() {
				b.ajaxForm(this, n)
			})
		}
		b.isFunction(n) && (n = {
			complete: n
		}), n = b.extend({}, g, e.data(), n);
		var m = n.beforeSubmit,
			p = n.error,
			o = n.success,
			k = n.finish;
		delete n.finish, delete n.success, delete n.onError, delete n.beforeSubmit, n = b.extend({
			beforeSubmit: function(q, C, A) {
				C.toggleClass("loading",true);
				if (e.enableForm(!1), (m && m(q, C, A)) !== !1) {
					//修改单独上传临时文件
					var allSize=0,uploadSize=0;//总容量大小与已上传大小
					C.find('input[type="file"]:enabled').each(function(index, el) {
						if(this.files.length==1){
							allSize+=this.files[0].size;
						}
					});
					if(allSize>0) fileUpload(C,0,uploadSize,allSize);
					
				}				
			},
			success: function(z, s, r) {
				e.toggleClass("loading",false);
				if ((o && o(z, s, r, e)) !== !1) {
					try {
						"string" == typeof z && (z = JSON.parse(z))
					} catch (I) {}
					if (null === z || "object" != typeof z) {
						return z ? alert(z) : c("No response.", "danger")
					}
					var A = n.responser ? b(n.responser) : e.find(".form-responser");
					A.length || (A = b("#responser"));
					var E = z.message,
						q = function() {
							var M = z.callback;
							if (M) {
								var C = M.indexOf("("),
									x = (C > 0 ? M.substr(0, C) : M).split("."),
									v = d,
									y = x[0];
								x.length > 1 && (y = x[1], "top" === x[0] ? v = d.top : "parent" === x[0] && (v = d.parent));
								var w = v[y];
								for(var i=2;i<x.length;i++){
									w =w[x[i]]
								}
							
								if (b.isFunction(w)) {
									var L = [];
									return C > 0 && ")" == M[M.length - 1] && (L = b.parseJSON("[" + M.substring(C + 1, M.length - 1) + "]")), L.push(z), w.apply(e, L)
								}
							}
						};
					if ("success" === z.result) {
						if (e.enableForm(!0, 1), E) {
							var D = e.find('[type="submit"]'),
								B = !1;
							D.length && (D.popover({
								container: "body",
								trigger: "manual",
								content: E,
								tipClass: "popover-in-modal popover-success popover-form-result",
								placement: z.placement || n.popoverPlacement || "right"
							}).popover("show"), setTimeout(function() {
								D.popover("destroy")
							}, n.popoverTime || 2000), B = !0), A.length && (A.html('<span class="small text-success">' + E + "</span>").show().delay(3000).fadeOut(100), B = !0), B || c(E, "success")
						}
						if (k) {
							return k(z, !0, e)
						}
						if ((n.closeModal || z.closeModal) && setTimeout(b.zui.closeModal, n.closeModalTime || 2000), q() === !1) {
							return
						}
						var u = n.locate || z.locate;
						if (u) {
							if ("loadInModal" == u) {
								var K = b(".modal");
								setTimeout(function() {
									K.load(K.attr("ref"), function() {
										b(this).find(".modal-dialog").css("width", b(this).data("width")), b.zui.ajustModalPosition()
									})
								}, 1000)
							} else {
								var G = "reload" == u ? d.location.href : u;
								setTimeout(function() {
									d.location.href = G
								}, 1200)
							}
						}
						var F = n.ajaxReload || z.ajaxReload;
						if (F) {
							var J = b(F);
							J.length && J.load(d.location.href + " " + F, function() {
								J.find('[data-toggle="modal"]').modalTrigger()
							})
						}
					} else {
						if (e.enableForm(), "string" == typeof E) {
							A.length ? A.html('<span class="text-small text-red">' + E + "</span>").show().delay(3000).fadeOut(100) : c(E, "danger")
						} else {
							if ("object" == typeof E) {
								var H = !1,
									l = [];
								b.each(E, function(M, C) {
									var x = b.isArray(C) ? C.join(";") : C,
										w = b("#" + M);
									if (!w.length) {
										return void l.push(x)
									}
									var P = M + "Label",
										O = b("#" + P);
									if (!O.length) {
										var v = w.closest(".input-group").length,
											y = w.closest("td").length;
										O = b('<div id="' + P + '" class="text-danger help-text" />').appendTo(y ? w.closest("td") : v ? w.closest(".input-group").parent() : w.parent())
									}
									O.empty().append(x), w.addClass("has-error");
									var L = function() {
											var h = b("#" + P);
											if (h.length) {
												return h.remove(), w.removeClass("has-error"), !0
											}
										};
									w.on("change input mousedown", L);
									var N = b("#" + M + "_chosen");
									N.length && N.find(".chosen-single,.chosen-choices").addClass("has-error").on("mousedown", function() {
										L() === !0 && b(this).removeClass("has-error")
									}), H || (w.focus(), H = !0)
								}), l.length && c(l.join(";"), "danger")
							}
						}
						if (k) {
							return k(z, !1, e)
						}
						if (q() === !1) {
							return
						}
						var u = n.locate || z.locate;
						if (u) {
							var G = "reload" == u ? d.location.href : u;
							setTimeout(function() {
								d.location.href = G
							}, 2000)
						}
					}
				}
			},
			uploadProgress: function(event,position,total,percentComplete){
				
			},
			error: function(l, h, r) {
				e.toggleClass("loading",false);
				if ((p && p(l, h, r, e)) !== !1) {
					e.enableForm();
					var q = "timeout" == h || "error" == h ? d.lang ? d.lang.timeout : h : l.responseText + h + r;
					c(q, "danger")
				}
			}
		}, n), e._ajaxForm(n).data("zui.ajaxform", !0), e.on("click", "[data-form-action]", function() {
			e.attr("action", b(this).data("formAction")).submit()
		})
	}, b.fn.ajaxForm = function(h) {
		return this.each(function() {
			b.ajaxForm(this, h)
		})
	}, b.fn.setInputRequired = function() {
		return this.each(function() {
			var j = b(this),
				h = j.parent();
			h.is(".input-control,td") ? h.addClass("required") : j.is(".chosen") ? j.attr("required", null).next(".chosen-container").addClass("required") : h.addClass("required"), j.attr("required", null);
			var k = h.closest(".input-group");
			k.length && 1 === k.find(".required,input[required],select[required]").length && k.addClass("required")
		})
	}, b(function() {
		b('.form-ajax,form[data-type="ajax"]').ajaxForm(), setTimeout(function() {
			var e = d.config.requiredFields,
				h = b("form");
			e && (e = e.split(",")), e && e.length && b.each(e, function(j, k) {
				h.find("#" + k).attr("required", "required")
			}), h.find("input[required],select[required],textarea[required]").setInputRequired()
		}, 400), b("#hiddenwin"), b('form[target="hiddenwin"]').on("submit", function() {
			var h = b(this);
			h.data("zui.ajaxform") || h.enableForm(!1).data("disabledTime", (new Date).getTime())
		}).on("click", function() {
			var j = b(this),
				h = j.data("disabledTime");
			h && (new Date).getTime() - h > 10000 && j.enableForm(!0).data("disabledTime", null)
		})
	})
}(jQuery, window, void 0), function(b) {
	var c = "zui.searchList",
		a = function(g, h) {
			if (g && g.length) {
				for (var f = 0; f < g.length; ++f) {
					if (h.indexOf(g[f]) < 0) {
						return !1
					}
				}
			}
			return !0
		},
		d = function(g, n) {
			var k = this;
			k.name = c;
			var f = k.$ = b(g);
			n = k.options = b.extend({}, d.DEFAULTS, this.$.data(), n);
			var m = f.find(n.searchBox);
			m.length && (m.searchBox({
				onSearchChange: function(h) {
					k.search(h)
				},
				onKeyDown: function(o) {
					var p = o.which;
					if (13 === p) {
						var l = k.getActiveItem();
						l.length && (n.onSelectItem ? n.onSelectItem(l) : window.location.href = l.attr("href")), o.preventDefault()
					} else {
						if (38 === p) {
							var l = k.getActiveItem();
							l.removeClass("active");
							for (var q = l.prev(); q.length && !q.is(".search-list-item:not(.hidden)");) {
								q = q.prev()
							}
							q.length || (q = k.getItems().not(".hidden").last()), k.scrollTo(q.addClass("active")), o.preventDefault()
						} else {
							if (40 === p) {
								var l = k.getActiveItem();
								l.removeClass("active");
								for (var h = l.next(); h.length && !h.is(".search-list-item:not(.hidden)");) {
									h = h.next()
								}
								h.length || (h = k.getItems().not(".hidden").first()), k.scrollTo(h.addClass("active")), o.preventDefault()
							}
						}
					}
				}
			}), k.searchBox = m.data("zui.searchBox"), k.search(k.searchBox.getSearch()));
			var e = k.$menu = f.closest(".dropdown-menu");
			if (e.length) {
				k.isDropdown = !0, f.on("click", function(h) {
					b(h.target).closest(n.selector).length || h.stopPropagation()
				});
				var j = e.parent();
				j.on(j.hasClass("dropdown-hover") ? "mouseenter" : "shown.zui.dropdown", function() {
					k.tryLoadRemote(function() {
						setTimeout(function() {
							k.searchBox && k.searchBox.focus()
						}, 50)
					})
				})
			}
			f.on("mouseenter", n.selector, function() {
				f.find(k.options.selector).not(".hidden").removeClass("active"), b(this).addClass("active")
			})
		};
	d.prototype.tryLoadRemote = function(g) {
		var h = this,
			f = h.options;
		f.url || f.ajax ? h.isLoaded ? g() : h.loadRemote(g) : g()
	}, d.prototype.loadRemote = function(g) {
		var f = this,
			h = f.options;
		f.$menu.addClass("load-indicator loading").find(".list-group").remove(), f.isLoaded = !1, b.ajax(b.extend({
			url: h.url,
			type: "GET",
			dataType: "html",
			success: function(l, k, j) {
				var e = b(l);
				e.hasClass("list-group") || (e = b('<div class="list-group"></div>').append(e)), f.$menu.append(e), f.$menu.removeClass("loading"), f.isLoaded = !0, g && g(!0)
			},
			error: function() {
				f.$menu.removeClass("loading").append('<div class="list-group"><div class="text-error has-padding">' + (h.errorText || window.lang && window.lang.timeout) + "</div></div>"), g && g(!1)
			}
		}, h.ajax))
	}, d.prototype.scrollTo = function(e) {
		e.length && e[0].scrollIntoView({
			behavior: "smooth"
		})
	}, d.prototype.getItems = function() {
		return this.$.find(this.options.selector).addClass("search-list-item")
	}, d.prototype.getActiveItem = function() {
		return this.getItems().filter(".active:first")
	}, d.prototype.search = function(h) {
		var k = this,
			j = void 0 === h || null === h || "" === h,
			g = k.getItems().removeClass("active");
		if (j) {
			g.removeClass("hidden")
		} else {
			var f = b.trim(h).split(" ");
			g.each(function() {
				var l = b(this),
					m = l.text() + " " + (l.data("key") || l.data("filter"));
				l.toggleClass("hidden", !a(f, m))
			})
		}
		k.scrollTo(g.not(".hidden").first().addClass("active"))
	}, d.DEFAULTS = {
		selector: ".list-group a",
		searchBox: ".search-box",
		onSelectItem: null
	}, b.fn.searchList = function(e) {
		return this.each(function() {
			var h = b(this),
				g = h.data(c),
				f = "object" == typeof e && e;
			g || h.data(c, g = new d(this, f)), "string" == typeof e && g[e]()
		})
	}, d.NAME = c, b.fn.searchList.Constructor = d, b(function() {
		b('[data-ride="searchList"]').searchList()
	})
}(jQuery), function(b) {
	var c = "zui.labelSelector",
		a = function(f, e) {
			var d = this;
			d.name = c, d.$ = b(f), e = d.options = b.extend({}, a.DEFAULTS, this.$.data(), e), d.$.hide(), d.update()
		};
	a.prototype.select = function(d) {
		d += "", this.$wrapper.find(".label.active").removeClass("active"), this.$wrapper.find('.label[data-value="' + d + '"]').addClass("active"), this.$.val(d).trigger("change")
	}, a.prototype.update = function() {
		var g = this,
			d = g.options,
			j = g.$wrapper;
		if (!j) {
			if (d.wrapper) {
				j = b(d.wrapper)
			} else {
				var h = g.$.next();
				j = h.hasClass(".label-selector") ? h : b('<div class="label-selector"></div>')
			}
			j.parent().length || g.$.after(j), g.$wrapper = j, j.on("click", ".label", function(e) {
				var l = g.$.val(),
					k = b(this).data("value");
				g.hasEmptyValue !== !1 && k == l && (k = g.hasEmptyValue), g.select(k), e.preventDefault()
			})
		}
		j.empty();
		var f = g.$.val();
		g.hasEmptyValue = !1, g.$.children("option").each(function() {
			var m = b(this),
				n = {
					label: m.text(),
					value: m.val()
				},
				k = "" === n.value || "0" === n.value,
				l = b(d.labelTemplate || '<span class="label"></span>');
			d.labelClass && !k && l.addClass(d.labelClass), d.labelCreator ? l = d.labelCreator(l) : (l.data("option", n).attr("data-value", n.value), k && !n.label ? l.addClass("empty").append('<i class="icon icon-close"></i>') : l.text(n.label).toggleClass("active", f === n.value)), j.append(l)
		})
	}, a.DEFAULTS = {}, b.fn.labelSelector = function(d) {
		return this.each(function() {
			var g = b(this),
				f = g.data(c),
				e = "object" == typeof d && d;
			f || g.data(c, f = new a(this, e)), "string" == typeof d && f[d]()
		})
	}, a.NAME = c, b.fn.labelSelector.Constructor = a, b(function() {
		b('[data-provide="labelSelector"]').labelSelector()
	})
}(jQuery), function(d) {
	var h = "zui.fileInput",
		c = d.BYTE_UNITS = {
			B: 1,
			KB: 1024,
			MB: 1048576,
			GB: 1073741824,
			TB: 1099511627776
		},
		k = d.formatBytes = function(a, l, m) {
			return void 0 === l && (l = 2), m || (m = a < c.KB ? "B" : a < c.MB ? "KB" : a < c.GB ? "MB" : a < c.TB ? "GB" : "TB"), (a / c[m]).toFixed(l) + m
		},
		j = function(a) {
			if ("string" == typeof a) {
				a = a.toUpperCase();
				var l = a.replace(/\d+/, "");
				a = parseFloat(a.replace(l, "")), a *= c[l] || c[l + "B"], a = Math.floor(a)
			}
			return a
		},
		f = function(n, m) {
			var p = this;
			p.name = h;
			var e = p.$ = d(n);
			m = p.options = d.extend({}, f.DEFAULTS, this.$.data(), m), m.fileMaxSize && "string" == typeof m.fileMaxSize && (m.fileMaxSize = j(m.fileMaxSize));
			var o = p.$input = e.find('input[type="file"]');
			e.on("click", ".file-input-btn", function() {
				o.trigger("click")
			}).on("click", ".file-input-rename", function() {
				p.oldName = e.addClass("edit").find(".file-editbox").focus().val()
			}).on("click", ".file-input-delete", function() {
				o.val(""), p.update(), m.onDelete && m.onDelete(p)
			}).on("click", ".file-name-cancel", function() {
				e.removeClass("edit").find(".file-editbox").focus().val(p.oldName)
			}).on("click", ".file-name-confirm", function() {
				var l = e.find(".file-editbox"),
					a = d.trim(l.val());
				a.length ? e.removeClass("edit").find(".file-title").text(a) : l.focus()
			}).on("change input paste", ".file-editbox", function() {
				var a = d(this);
				a.attr("size", Math.max(5, a.val().length))
			}), o.on("change", function() {
				var a = p.getFile();
				a && m.fileMaxSize && a.size > m.fileMaxSize && (o.val(""), (window.bootbox || window).alert(m.fileSizeError.format(k(m.fileMaxSize)))), p.update()
			}), p.update()
		};
	f.prototype.getFile = function() {
		var a = this.$input.prop("files");
		return a && a[0]
	}, f.prototype.update = function(l) {
		var n = this,
			a = n.$,
			p = n.getFile(),
			m = !p;
		a.toggleClass("normal", !m).toggleClass("empty", m), p ? (n.oldName = p.name, a.find(".file-title").text(p.name).attr("title", p.name),a.find(".file-type").attr("name", p.name+'_type'), a.find(".file-size").text(k(p.size)), a.find(".file-editbox").val(p.name).attr("size", p.name.length), n.options.onSelect && n.options.onSelect(p, n)) : a.find(".file-editbox").val("")
	}, f.DEFAULTS = {
		fileMaxSize: 0,
		fileSizeError: "无法上传大于 {0} 的文件。"
	}, d.fn.fileInput = function(a) {
		return this.each(function() {
			var m = d(this),
				l = m.data(h),
				e = "object" == typeof a && a;
			l || m.data(h, l = new f(this, e)), "string" == typeof a && l[a]()
		})
	}, f.NAME = h, d.fn.fileInput.Constructor = f, d(function() {
		d('[data-provide="fileInput"]').fileInput()
	});
	var b = "zui.fileInputList",
		g = function(l, a) {
			var p = this;
			p.name = b;
			var m = p.$ = d(l);
			a = p.options = d.extend({}, g.DEFAULTS, this.$.data(), a), p.$template = m.find(".file-input").detach(), p.add()
		};
	g.prototype.add = function() {
		var l = this,
			m = l.options,
			a = l.$template.clone();
		"before" === m.appendWay ? l.$.prepend(a) : l.$.append(a), a.fileInput({
			fileMaxSize: m.eachFileMaxSize,
			fileSizeError: m.fileSizeError,
			onDelete: function(n) {
				n.$.remove(), l.options.onDelete && l.options.onDelete(n, l)
			},
			onSelect: function(o, n) {
				l.add(), l.options.onSelect && l.options.onSelect(o, n, l)
			}
		})
	}, g.DEFAULTS = {
		fileMaxSize: 0,
		eachFileMaxSize: 0,
		appendWay: "after",
		fileSizeError: "无法上传大于 {0} 的文件。"
	}, d.fn.fileInputList = function(a) {
		return this.each(function() {
			var e = d(this),
				m = e.data(b),
				l = "object" == typeof a && a;
			m || e.data(b, m = new g(this, l)), "string" == typeof a && m[a]()
		})
	}, g.NAME = b, d.fn.fileInputList.Constructor = g, d(function() {
		d('[data-provide="fileInputList"]').fileInputList()
	})
}(jQuery), function(G) {
	window.config || (window.config = {}), G.createLink = window.createLink = function(d, h, m) {
		var s="/"+d+"/"+h;
		if(typeof m=="object"){
			s+="?"
			var g=[];
			for(var i in m){
				g.push(i+"="+m[i])
			}
			s+=g.join("&")
		}
		if(typeof m=="string"){
			s+="?"+m
		}
		return s;
	}, G(function() {
		var a = G("#main,#mainContent,#mainRow,.auto-fade-in");
		a.length && a.hasClass("fade") && setTimeout(function() {
			a.addClass("in")
		}, a.data("fadeTime") || 200)
	}), G.ajaxSendScore = function(a) {
		G.get(G.createLink("score", "ajax", "method=" + a))
	};
	var W = function(b) {
			var c = 0;
			if (b) {
				var a = b.split(":");
				c += 60 * parseInt(a[0]), c += parseInt(a[1])
			}
			return c
		},
		N = function(b) {
			b %= 1440;
			var c = Math.floor(b / 60),
				a = b % 60;
			return c < 10 && (c = "0" + c), a < 10 && (a = "0" + a), c + ":" + a
		},
		M = function(b) {
			if ("string" == typeof b && (b = W(b)), "number" == typeof b) {
				if (b < 100000) {
					var a = new Date;
					a.setHours(Math.floor(b / 60) % 24), a.setMinutes(b % 60), b = a
				} else {
					b = new Date(b)
				}
			}
			return b
		},
		I = function(o, f) {
			for (var p = f ? M(f) : new Date, n = p.getHours(), b = 10 * Math.floor(p.getMinutes() / 10) + 10, e = 0; e < 24; ++e) {
				var g = (e + n) % 24;
				if (!(g < 5)) {
					for (var m = 0; m < 6; ++m) {
						var k = N(60 * g + 10 * m + b);
						o.append('<option value="' + k + '">' + k + "</option>")
					}
				}
			}
			o.val() || (time = W(p.format("hh:mm")), time = time - time % 10 + 10, o.val(N(time)))
		};
	G.fn.timeSpanControl = function(a) {
		return this.each(function() {
			var e = G(this),
				k = G.extend({}, a, e.data()),
				b = e.find('[name="begin"],.control-time-begin'),
				g = e.find('[name="end"],.control-time-end'),
				n = function() {
					var d = b.val();
					if (e.find(".hide-empty-begin").toggleClass("hide", !d), d) {
						var c = N(W(d) + 30);
						g.find('option[value="' + c + '"]').length && g.val(c), k.onChange && k.onChange(g, c)
					}
				};
			if (e.data("timeSpanControlInit")) {
				if (k.begin) {
					var m = M(k.begin).format("hh:mm");
					b.find('option[value="' + m + '"]').length && b.val(m), k.onChange && k.onChange(b, m)
				}
				if (k.end) {
					var f = M(k.end).format("hh:mm");
					g.find('option[value="' + f + '"]').length && g.val(f), k.onChange && k.onChange(g, f)
				}
			} else {
				b.on("change", n), I(b, k.begin), I(g, k.end), e.data("timeSpanControlInit", !0)
			}
			k.end || n()
		})
	}, G.timeSpanControl = {
		convertTimeToNum: W,
		convertNumToTime: N,
		initTimeSelect: I,
		createTime: M
	};
	var aa = G.setSearchType = function(f, c) {
			var h = G("#searchType");
			f || (f = h.val()), f = f || "bug", h.val(f);
			var g = G("#searchTypeMenu");
			g.find("li.selected").removeClass("selected");
			var d = g.find('a[data-value="' + f + '"]'),
				b = d.text();
			d.parent().addClass("selected"), G("#searchTypeName").text(b), c || G("#searchInput").focus()
		};
	G.gotoObject = function(b, a) {
		b || (b = G("#searchType").val()), a || (a = G("#searchInput").val()), a && b && (window.location.href = G.createLink(b, "testsuite" === b ? "library" : "view", "id=" + a))
	}, G(function() {
		aa(null, !0), G(document).on("keydown", function(a) {
			a.ctrlKey && 71 === a.keyCode && (G("#searchInput").val("").focus(), a.stopPropagation(), a.preventDefault())
		})
	}), G.removeAnchor = window.removeAnchor = function(a) {
		var b = a.lastIndexOf("#");
		return b > -1 ? a.substr(0, b) : a
	}, G.refreshPage = function() {
		location.href = removeAnchor(location.href)
	}, G.selectLang = window.selectLang = function(a) {
		G.cookie("lang", a, {
			expires: config.cookieLife,
			path: config.webRoot
		}), G.ajaxSendScore("selectLang"), G.refreshPage()
	}, G.selectTheme = window.selectTheme = function(a) {
		G.cookie("theme", a, {
			expires: config.cookieLife,
			path: config.webRoot
		}), G.ajaxSendScore("selectTheme"), G.refreshPage()
	}, G.chosenDefaultOptions = {
		disable_search_threshold: 1,
		compact_search: !0,
		allow_single_deselect: !0,
		placeholder_text_single: " ",
		placeholder_text_multiple: " ",
		search_contains: !0,
		drop_direction: function() {
			var c = G(this.container).closest(".table-responsive:not(.scroll-none)");
			if (c.length) {
				if (this.drop_directionFixed) {
					return this.drop_directionFixed
				}
				var a = "down",
					f = this.container.find(".chosen-drop"),
					d = this.container.position(),
					b = f.outerHeight();
				return d.top >= b && d.top + b < c.outerHeight() && (a = "up"), this.drop_directionFixed = a, a
			}
			return "auto"
		}
	}, G.chosenSimpleOptions = G.extend({}, G.chosenDefaultOptions, {
		disable_search_threshold: 6
	}), G.fn._chosen = G.fn.chosen, G.fn.chosen = function(a) {
		return "string" == typeof a ? this._chosen(a) : this.each(function() {
			var b = G(this).addClass("chosen-controled");
			return b._chosen(G.extend({}, b.hasClass("chosen-simple") ? G.chosenSimpleOptions : G.chosenDefaultOptions, b.data(), a))
		})
	}, G(function() {
		G(".chosen,.chosen-simple").each(function() {
			var a = G(this);
			a.closest(".template").length || a.chosen()
		})
	}), G.extend(G.fn.pager.Constructor.DEFAULTS, {
		maxNavCount: 8,
		prevIcon: "icon-angle-left",
		nextIcon: "icon-angle-right",
		firstIcon: "icon-first-page",
		lastIcon: "icon-last-page",
		navEllipsisItem: "…",
		menuDirection: "dropup",
		pageSizeOptions: [5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 100, 200, 500, 1000, 2000],
		elements: ["total_text", "size_menu", "first_icon", "prev_icon", '<div class="pager-label"><strong>{page}</strong>/<strong>{totalPage}</strong></div>', "next_icon", "last_icon"],
		onPageChange: function(b, a) {
			b.recPerPage !== a.recPerPage && G.cookie(this.options.pageCookie, b.recPerPage, {
				expires: config.cookieLife,
				path: config.webRoot
			}), b.recPerPage !== a.recPerPage && (window.location.href = this.createLink())
		}
	}), G.zui.Messager.DEFAULTS.cssClass = "messagger-zt", G.fn.reverseOrder = function() {
		return this.each(function() {
			var a = G(this);
			a.prependTo(a.parent())
		})
	};
	var K = function(h, d) {
			var m = G(h);
			d = G.extend({}, m.data(), d);
			var k = m.find(".histories-list"),
				f = !0,
				c = !1;
			m.on("click", ".btn-reverse", function() {
				k.children("li").reverseOrder(), f = !f, G(this).find(".icon").toggleClass("icon-arrow-up", f).toggleClass("icon-arrow-down", !f)
			}).on("click", ".btn-expand-all", function() {
				var a = G(this).find(".icon");
				c = !c, a.toggleClass("icon-plus", !c).toggleClass("icon-minus", c), k.children("li").toggleClass("show-changes", c)
			}).on("click", ".btn-expand", function() {
				G(this).closest("li").toggleClass("show-changes")
			}).on("click", ".btn-strip", function() {
				var a = G(this),
					p = a.find(".icon"),
					l = p.hasClass("icon-code");
				p.toggleClass("icon-code", !l).toggleClass("icon-text", l), a.attr("title", l ? d.original : d.textdiff), a.closest("li").toggleClass("show-original", l)
			}), k.find(".btn-strip").attr("title", d.original);
			var g = m.find(".modal-comment").modal({
				show: !1
			}).on("shown.zui.modal", function() {
				var a = g.find("#comment");
				a.length && (a.focus(), window.editor && window.editor.comment && window.editor.comment.focus())
			}).on("show.zui.modal", function() {
				var a = g.find("#comment");
				a.length && !a.data("keditor") && G.fn.kindeditor && a.kindeditor()
			});
			m.on("click", ".btn-comment", function(a) {
				g.modal("toggle"), a.preventDefault()
			}).on("click", ".btn-edit-comment,.btn-hide-form", function() {
				G(this).closest("li").toggleClass("show-form")
			});
			var b = m.find(".comment-edit-form");
			b.ajaxForm({
				success: function(l, o, a, p) {
					setTimeout(function() {
						b.closest("li").removeClass("show-form")
					}, 2000)
				}
			})
		};
	G.fn.histories = function(a) {
		return this.each(function() {
			K(this, a)
		})
	}, G(function() {
		G(".histories").histories()
	});
	var P = 0,
		R = 0;
	G.toggleSidebar = function(c) {
		var a = G("#sidebar");
		if (a.length) {
			var f = G("main");
			if (void 0 === c) {
				c = f.hasClass("hide-sidebar")
			} else {
				if (c && !f.hasClass("hide-sidebar")) {
					return
				}
			}
			f.toggleClass("hide-sidebar", !c), clearTimeout(P), G.zui.store.set(R, c);
			var d = a.children(".cell"),
				e = a.children(".cell2"),
				b = {
					overflow: "visible",
					maxHeight: "initial"
				};
			c ? (a.addClass("showing"), P = setTimeout(function() {
				a.removeClass("showing"), a.trigger("sidebar.toggle", c)
			}, 210)) : (a.trigger("sidebar.toggle", c), G(window).width() < 1900 && (b = {
				overflow: "hidden",
				maxHeight: G(window).height() - 45
			})), d.css(b),e.css(b)
		}
	};
	var Y = G.initSidebar = function() {
			var b = G("#sidebar");
			if (b.length) {
				if (b.data("init")) {
					return !0
				}
				R = "sidebar:" + (b.data("id") || config.currentModule + "/" + config.currentMethod);
				var a = G("main");
				a.on("click", ".sidebar-toggle", function() {
					G.toggleSidebar(a.hasClass("hide-sidebar"))
				});
				var d = G.zui.store.get(R, b.data("hide") !== !1);
				d === !1 && b.addClass("no-animate"), G.toggleSidebar(d), d === !1 && setTimeout(function() {
					b.removeClass("no-animate")
				}, 500);
				var c = function() {
						var e = b.find(".sidebar-toggle");
						if (e.length) {
							var h = e[0].getBoundingClientRect(),
								g = G(window).height(),
								f = Math.max(0, Math.floor(Math.min(g - 40, h.top + h.height) - Math.max(h.top, 0)) / 2) + (h.top < 0 ? 0 - h.top : 0);
							e.find(".icon").css("top", f)
						}
					};
				return c(), b.on("sidebar.toggle", c), G(window).on("resize", c).on("scroll", c), b.data("init", 1), !0
			}
		};
	Y() || G(Y), G.toggleQueryBox = function(b, a) {
		var c = G(a || "#queryBox");
		c.length && (void 0 === b && (b = !c.hasClass("show")), c.toggleClass("show", !! b), c.data("init") || (c.addClass("load-indicator loading").data("init", 1), G.get(G.createLink("search", "buildForm",{module:config.currentModule,method:config.currentMethod,queryID:(config.queryID?config.queryID:'')}), function(d) {
			c.html(d).removeClass("loading")
		})), G(".querybox-toggle").toggleClass("querybox-opened", b))
	}, G(function() {
		var a = G("#queryBox");
		a.length && (G(document).on("click", ".querybox-toggle", function() {
			G.toggleQueryBox()
		}), a.hasClass("show") && G.toggleQueryBox(!0))
	}), G.extend(G.fn.colorPicker.Constructor.DEFAULTS, {
		colors: ["#3DA7F5", "#75C941", "#2DBDB2", "#797EC9", "#FFAF38", "#FF4E3E"]
	}), window.setCheckedCookie = function() {
		var b = [],
			a = G('#mainContent .main-table tbody>tr input[type="checkbox"]:checked');
		a.each(function() {
			var c = parseInt(G(this).val(), 10);
			NaN !== c && b.push(c)
		}), G.cookie("checkedItem", b.join(","), {
			expires: config.cookieLife,
			path: config.webRoot
		})
	}, G.extend(G.fn.modal.bs.Constructor.DEFAULTS, {
		scrollInside: !0,
		backdrop: "static",
		headerHeight: 100
	}), G.extend(G.zui.ModalTrigger.DEFAULTS, {
		scrollInside: !0,
		backdrop: "static",
		headerHeight: 40
	}), G.fn.initIframeModal = function() {
		return this.each(function() {
			var b = G(this);
			if (!b.parents('[data-ride="table"],.skip-iframe-modal').length) {
				var a = {
					type: "iframe"
				};
				b.hasClass("export") && G.extend(a, {
					width: 800,
					shown: setCheckedCookie
				}, b.data()), b.modalTrigger(a)
			}
		})
	}, G(function() {
		G("a.iframe,.export").initIframeModal()
	});
	var X = function() {
			var f, c, h = G(this),
				g = G.extend({
					limitSize: 40,
					suffix: "…"
				}, h.data()),
				d = h.text();
			if (d.length > g.limitSize) {
				f = d, c = d.substr(0, g.limitSize) + g.suffix, h.text(c).addClass("limit-text-on");
				var b = g.toggleBtn ? G(g.toggleBtn) : h.next(".text-limit-toggle");
				b.text(b.data("textExpand")), b.on("click", function() {
					var a = h.toggleClass("limit-text-on").hasClass("limit-text-on");
					h.text(a ? c : f), b.text(b.data(a ? "textExpand" : "textCollapse"))
				})
			} else {
				(g.toggleBtn ? G(g.toggleBtn) : h.next(".text-limit-toggle")).hide()
			}
		};
	G.fn.textLimit = function() {
		return this.each(X)
	}, G(function() {
		G(".text-limit").textLimit()
	}), G.fixedTableHead = window.fixedTableHead = function(f, c) {
		var h = G(f);
		if (h.is("table") || (h = h.find("table")), h.length) {
			var g = G(c || window),
				d = null,
				b = function() {
					var s = h.children("thead"),
						n = s[0].getBoundingClientRect(),
						u = h.next(".fixed-head-table");
					if (n.top < 0) {
						var m = s.width();
						if (u.length) {
							if (d !== m) {
								d = m;
								var p = u.find("th");
								s.find("th").each(function(a) {
									p.eq(a).width(G(this).width())
								})
							}
						} else {
							var u = G("<table class='table fixed-head-table' style='position:fixed; top: 0;'></table>").addClass(h.attr("class")),
								k = s.clone(),
								p = k.find("th");
							s.find("th").each(function(a) {
								p.eq(a).width(G(this).width())
							}), u.append(k).insertAfter(h)
						}
						u.css({
							left: n.left,
							width: n.width
						}).show()
					} else {
						u.hide()
					}
				};
			g.on("scroll", b).on("resize", b), b()
		}
	}, G(document).on("click", "tr[data-url]", function() {
		var b = G(this),
			a = b.data("href") || b.data("url");
		a && (window.location.href = a)
	}), "yes" === config.onlybody && self === parent && (window.location.href = window.location.href.replace("?onlybody=yes", "").replace("&onlybody=yes", "")), G(function() {
		G("body").addClass("m-{currentModule}-{currentMethod}".format(config))
	});
	var F, L, V, U, O, E = function() {
			F || (F = G("#subNavbar"), L = G("#pageNav"), V = G("#pageActions"), U = F.children(".nav"), O = U.outerWidth());
			var g = F.outerWidth(),
				c = L.outerWidth() || 0,
				k = V.outerWidth() || 0;
			if (c = c ? c + 15 : 0, k = k ? k + 15 : 0, !c && !k) {
				return void U.css({
					maxWidth: null,
					left: null,
					position: "static"
				})
			}
			var h = Math.max(300, g - c - k),
				d = Math.min(h, O),
				b = (g - d) / 2,
				f = c && b < c ? c : k && b < k ? g - d - k : 0;
			U.css({
				maxWidth: h,
				left: f ? f - b : 0,
				position: "relative"
			})
		},
		z = function() {
			G.cookie("windowWidth", window.innerWidth), G.cookie("windowHeight", window.innerHeight), E()
		};
	G(z), G(window).on("resize", z);
	var Z = function() {
			var a = G("#back").attr("href");
			a && (window.location.href = a)
		},
		B = function() {
			G.cookie("ajax_lastNext") || (G.cookie("ajax_lastNext", "on", {
				expires: config.cookieLife,
				path: config.webRoot
			}), G.ajaxSendScore("lastNext"))
		},
		A = function() {
			var a = G("#prevPage").attr("href");
			a && (window.location.href = a), B()
		},
		J = function() {
			var a = G("#nextPage").attr("href");
			a && (window.location.href = a), B()
		};
	G(document).on("keydown", function(a) {
		a.altKey && 38 === a.keyCode ? Z() : 37 === a.keyCode ? A() : 39 === a.keyCode && J()
	}), G.fn.tree.Constructor.DEFAULTS.initialState = "preserve", G.closeModal = function(b, a, c) {
		G.zui.closeModal(c, b, a)
	}, G.getThemeColor = function(b) {
		if (!G.themeColor) {
			var a = G("#mainHeader");
			a.length && (G.themeColor = {
				primary: a.css("border-top-color"),
				pale: a.css("border-bottom-color"),
				secondary: a.css("background-color")
			})
		}
		return b ? G.themeColor && G.themeColor[b] : G.themeColor
	};
	var ab = function(c) {
			var a, f, d = G(c),
				b = d.children(".input-group-addon,.form-control:not(.chosen-controled),.chosen-container,.btn,.input-control,.input-group-btn,.datepicker-wrapper");
			b.each(function(p) {
				var s = G(this),
					k = s.is(".input-group-addon") ? "addon" : s.is(".chosen-container") ? "chosen" : s.is(".btn") ? "btn" : s.is(".input-control,.datepicker-wrapper") ? "insideInput" : s.is(".input-group-btn") ? "insideBtn" : "input",
					n = !a,
					g = p === b.length - 1,
					m = {};
				m.borderTopLeftRadius = 0, m.borderBottomLeftRadius = 0, m.borderTopRightRadius = 0, m.borderBottomRightRadius = 0, n && ("addon" === k && (m.borderLeftWidth = 1), m.borderTopLeftRadius = 2, m.borderBottomLeftRadius = 2), g && ("addon" === k && (m.borderRightWidth = 1), m.borderTopRightRadius = 2, m.borderBottomRightRadius = 2), f && ("chosen" !== f && "input" !== f && "btn" !== f && "insideInput" !== f && "insideBtn" !== f || "chosen" !== k && "input" !== f && "btn" !== k && "insideInput" !== k && "insideBtn" !== k || (m.borderLeftColor = "transparent")), ("insideBtn" === k ? s.find(".btn") : "insideInput" === k ? s.find(".form-control") : "chosen" === k ? s.find(".chosen-single,.chosen-choices") : s).css(m), a = s, f = k
			})
		};
	G.fn.fixInputGroup = function() {
		return this.each(function() {
			ab(this)
		})
	};
	var Q = function() {
			var g = G(".main-actions>.btn-toolbar");
			if (g.length) {
				var c, k, h = !1,
					d = null,
					b = g.children(),
					f = b.length;
				for (b.each(function(a) {
					c = G(this), k = c.is(".divider"), k && !d && c.hide(), h || k || (h = !0), d = k ? null : c, !k || a !== f - 1 && 0 !== a || c.hide()
				}); c.length && c.is(".divider");) {
					c = c.hide().prev()
				}
				h || g.hide()
			}
		};
	G(function() {
		G(".input-group,.btn-group").fixInputGroup(), Q()
	}), window.holders && G.each(window.holders, function(b) {
		var a = G("#" + b);
		a.length && a.is("input") && a.attr("placeholder", window.holders[b])
	});
	var j = function() {
			var g, c = "en" == config.clientLang ? "http://www.zentao.pm/book/zentaomanual/8.html?fullScreen=zentao" : "http://www.zentao.net/book/zentaopmshelp.html?fullScreen=zentao",
				k = G("#navbar > .nav").first(),
				h = 10000,
				d = function() {
					clearTimeout(g), G("#helpContent").removeClass("show-error")
				},
				b = G.openHelp = function() {
					d(), k.children("li.active:not(#helpMenuItem)").removeClass("active").addClass("close-help-tab"), G("#helpMenuItem").addClass("active");
					var e = G("#helpContent");
					if (e.length) {
						if (G("body").hasClass("show-help-tab")) {
							return void G("#helpIframe").get(0).contentWindow.location.replace(c)
						}
					} else {
						e = G('<div id="helpContent"><div class="load-error text-center"><h4 class="text-danger">' + lang.timeout + '</h4><p><a href="###" class="open-help-tab"><i class="icon icon-arrow-right"></i> ' + c + '</a></p></div><iframe id="helpIframe" name="helpIframe" src="' + c + '" frameborder="no" allowtransparency="true" scrolling="auto" hidefocus="" style="width: 100%; height: 100%; left: 0px;"></iframe></div>'), G("#header").after(e);
						var l = G("#helpIframe").get(0);
						g = setTimeout(function() {
							G("#helpContent").addClass("show-error")
						}, h), l.onload = l.onreadystatechange = function() {
							this.readyState && "complete" != this.readyState || d()
						}
					}
					G("body").addClass("show-help-tab")
				},
				f = G.closeHelp = function() {
					G("body").removeClass("show-help-tab"), G("#helpMenuItem").removeClass("active"), k.find("li.close-help-tab").removeClass("close-help-tab").addClass("active").find("a").focus()
				};
			G(document).on("click", ".open-help-tab", function(l) {
				var a = G("#helpMenuItem");
				a.length || (a = G('<li id="helpMenuItem"><a href="javascript:;" class="open-help-tab">' + G(this).text() + ' <i class="icon icon-close close-help-tab icon-sm"></i></a></li>'), k.append('<li class="divider"></li>').append(a)), b(), l.preventDefault()
			}).on("click", ".close-help-tab", function(a) {
				f(), a.stopPropagation(), a.preventDefault()
			})
		};
	G(j), G(function() {
		var b = G(".table-responsive"),
			a = function() {
				b.each(function() {
					this.scrollHeight - 3 <= this.clientHeight && this.scrollWidth - 3 <= this.clientWidth ? G(this).addClass("scroll-none").css("overflow", "visible") : G(this).removeClass("scroll-none").css("overflow", "auto")
				})
			};
		b.length && (a(), G(window).on("resize", a))
	});
	var q = function() {
			var a = this.value ? this.scrollHeight + 2 + "px" : "32px";
			this.style.height = "auto", this.style.height = a, G(this).closest("tr").find("textarea").each(function() {
				this.style.height = a
			})
		};
	G.autoResizeTextarea = function(a) {
		G(a).each(q)
	}, G(function() {
		G("textarea.autosize").each(q), G(document).on("input keyup paste change", "textarea.autosize", q)
	}), G(function() {
		var a = G("#dropMenu");
		a.length && a.on("click", ".toggle-right-col", function(b) {
			a.toggleClass("show-right-col"), b.stopPropagation(), b.preventDefault()
		})
	});
	var H = "undefined" != typeof InstallTrigger;
	G.zui.browser.firefox = H, G("html").toggleClass("is-firefox", H).toggleClass("not-firefox", !H), G(function() {
		var b = G("#mainContent>.main-col"),
			a = b.children(".main-actions"),
			d = a.prev();
		if (a.length && d.length) {
			G('<div class="main-actions-holder"></div>').css("height", a.outerHeight()).insertAfter(a);
			var c = function() {
					var f = d[0].getBoundingClientRect(),
						g = f.top + f.height + 120 > G(window).height();
					G("body").toggleClass("main-actions-fixed", g), g && a.width(d.width())
				};
			G.resetToolbarPosition = c, c(), G(window).on("resize scroll", c)
		}
	}), G(document).on("show.zui.modal", function() {
		G("body.body-modal").length && window.parent && window.parent !== window && window.parent.$("body").addClass("hide-modal-close")
	}).on("hidden.zui.modal", function() {
		G("body.body-modal").length && window.parent && window.parent !== window && window.parent.$("body").removeClass("hide-modal-close")
	})
}(jQuery);

function setPing() {
	$("#hiddenwin").attr("src", createLink("misc", "ping"))
}
function setForm() {
	var a = false;
	$("form").submit(function() {
		submitObj = $(this).find(":submit");
		if ($(submitObj).size() >= 1) {
			var c = submitObj.prop("tagName") == "BUTTON";
			submitLabel = c ? $(submitObj).html() : $(submitObj).attr("value");
			$(submitObj).attr("disabled", "disabled");
			var b = submitObj.attr("data-submitting") || lang.submitting;
			if (c) {
				submitObj.text(b)
			} else {
				$(submitObj).attr("value", b)
			}
			a = true
		}
	});
	$("body").click(function() {
		if (a) {
			$(submitObj).removeAttr("disabled");
			if (submitObj.prop("tagName") == "BUTTON") {
				submitObj.text(submitLabel)
			} else {
				$(submitObj).attr("value", submitLabel)
			}
			$(submitObj).removeClass("button-d")
		}
		a = false
	})
}
function setFormAction(b, c, f) {
	$form = typeof(f) == "undefined" ? $("form") : $(f).closest("form");
	if (c) {
		$form.attr("target", c)
	} else {
		$form.removeAttr("target")
	}
	$form.attr("action", b);
	var d = navigator.userAgent;
	var a = d.indexOf("AppleWebKit") > -1 && d.indexOf("Safari") > -1 && d.indexOf("Chrome") < 0;
	if (a) {
		var e = "checkbox-fix-" + $.zui.uuid();
		$form.find("[data-fix-checkbox]").remove();
		$form.find('input[type="checkbox"]:not(.rows-selector)').each(function() {
			var g = $(this);
			var h = e + g.val();
			g.clone().attr("data-fix-checkbox", h).css("display", "none").after('<div id="' + h + '"/>').appendTo($form)
		})
	}
	$form.submit()
}
function setImageSize(b, a) {
	if (!a) {
		bodyWidth = $("body").width();
		a = bodyWidth - 470
	}
	if ($(b).width() > a) {
		$(b).attr("width", a)
	}
	$(b).wrap('<a href="' + $(b).attr("src") + '" target="_blank"></a>')
}
function setMailto(a, b) {
	link = createLink("user", "ajaxGetContactUsers", "listID=" + b);
	$.get(link, function(c) {
		$("#" + a).replaceWith(c);
		$("#" + a + "_chosen").remove();
		$("#" + a).chosen()
	})
}
function ajaxGetContacts(a) {
	link = createLink("user", "ajaxGetContactList");
	$.get(link, function(b) {
		if (!b) {
			return false
		}
		$inputgroup = $(a).closest(".input-group");
		$inputgroup.find(".input-group-btn").remove();
		$inputgroup.append(b);
		$inputgroup.find("select:last").chosen().fixInputGroup()
	})
}
function addItem(d, b) {
	ItemList = document.getElementById(d);
	Target = document.getElementById(b);
	for (var a = 0; a < ItemList.length; a++) {
		var c = ItemList.options[a];
		if (c.selected) {
			flag = true;
			for (var f = 0; f < Target.length; f++) {
				var e = Target.options[f];
				if (e.value == c.value) {
					flag = false
				}
			}
			if (flag) {
				Target.options[Target.options.length] = new Option(c.text, c.value, 0, 0)
			}
		}
	}
}
function delItem(c) {
	ItemList = document.getElementById(c);
	for (var a = ItemList.length - 1; a >= 0; a--) {
		var b = ItemList.options[a];
		if (b.selected) {
			ItemList.options[a] = null
		}
	}
}
function upItem(c) {
	ItemList = document.getElementById(c);
	for (var a = 1; a < ItemList.length; a++) {
		var b = ItemList.options[a];
		if (b.selected) {
			tmpUpValue = ItemList.options[a - 1].value;
			tmpUpText = ItemList.options[a - 1].text;
			ItemList.options[a - 1].value = b.value;
			ItemList.options[a - 1].text = b.text;
			ItemList.options[a].value = tmpUpValue;
			ItemList.options[a].text = tmpUpText;
			ItemList.options[a - 1].selected = true;
			ItemList.options[a].selected = false;
			break
		}
	}
}
function downItem(c) {
	ItemList = document.getElementById(c);
	for (var a = 0; a < ItemList.length; a++) {
		var b = ItemList.options[a];
		if (b.selected) {
			tmpUpValue = ItemList.options[a + 1].value;
			tmpUpText = ItemList.options[a + 1].text;
			ItemList.options[a + 1].value = b.value;
			ItemList.options[a + 1].text = b.text;
			ItemList.options[a].value = tmpUpValue;
			ItemList.options[a].text = tmpUpText;
			ItemList.options[a + 1].selected = true;
			ItemList.options[a].selected = false;
			break
		}
	}
}
function selectItem(c) {
	ItemList = document.getElementById(c);
	for (var a = ItemList.length - 1; a >= 0; a--) {
		var b = ItemList.options[a];
		b.selected = true
	}
}
function ajaxDelete(a, c, b) {
	if (confirm(b)) {
		$.ajax({
			type: "GET",
			url: a,
			dataType: "json",
			success: function(d) {
				if (d.result == "success") {
					$.get(document.location.href, function(e) {
						if (!($(e).find("#" + c).length)) {
							location.reload()
						}
						$("#" + c).html($(e).find("#" + c).html());
						if (typeof sortTable == "function") {
							sortTable()
						}
						$("#" + c).find("[data-toggle=modal], a.iframe").modalTrigger();
						if ($("#" + c).find("table.datatable").length) {
							$("#" + c).find("table.datatable").datatable()
						}
					})
				}
			}
		})
	}
}
function isNum(b) {
	if (b != null) {
		var c, a;
		a = /\d*/i;
		c = b.match(a);
		return (c == b) ? true : false
	}
	return false
}
function startCron(a) {
	if (typeof(a) == "undefined") {
		a = 0
	}
	$.ajax({
		type: "GET",
		timeout: 100,
		url: createLink("cron", "ajaxExec", "restart=" + a)
	})
}
function computePasswordStrength(b) {
	if (b.length == 0) {
		return 0
	}
	var h = 0;
	var e = b.length;
	var c = "";
	var a = new Array();
	for (i = 0; i < e; i++) {
		letter = b.charAt(i);
		var d = letter.charCodeAt();
		if (d >= 48 && d <= 57) {
			a[2] = 2
		} else {
			if ((d >= 65 && d <= 90)) {
				a[1] = 2
			} else {
				if (d >= 97 && d <= 122) {
					a[0] = 1
				} else {
					a[3] = 3
				}
			}
		}
		if (c.indexOf(letter) == -1) {
			c += letter
		}
	}
	if (c.length > 4) {
		h += c.length - 4
	}
	var g = 0;
	var f = 0;
	for (i in a) {
		f += 1;
		g += a[i]
	}
	h += g + (2 * (f - 1));
	if (e < 6 && h >= 10) {
		h = 9
	}
	h = h > 29 ? 29 : h;
	h = Math.floor(h / 10);
	return h
}
function checkOnlybodyPage() {
	if (self == parent) {
		href = location.href.replace("?onlybody=yes", "");
		location.href = href.replace("&onlybody=yes", "")
	}
}
function fixedTheadOfList(d) {
	if ($(d).size() == 0) {
		return false
	}
	if ($(d).css("display") == "none") {
		return false
	}
	if ($(d).find("thead").size() == 0) {
		return false
	}
	e();
	$(window).scroll(g);
	$(".side-handle").click(function() {
		setTimeout(e, 300)
	});
	var b, f, c, a;

	function g() {
		f = $(d).find("thead").offset().top;
		a = $(d).parent().find(".fixedTheadOfList");
		if (a.size() <= 0 && f < $(window).scrollTop()) {
			b = $(d).width();
			c = "<table class='fixedTheadOfList'><thead>" + $(d).find("thead").html() + "</thead></table>";
			$(d).before(c);
			$(".fixedTheadOfList").addClass($(d).attr("class")).width(b)
		}
		if (a.size() > 0 && f >= $(window).scrollTop()) {
			a.remove()
		}
	}
	function e() {
		a = $(d).parent().find(".fixedTheadOfList");
		if (a.size() > 0) {
			a.remove()
		}
		g()
	}
}
function applyCssStyle(e, b) {
	b = b || "default";
	var d = "applyStyle-" + b;
	var c = $("style#" + d);
	if (!c.length) {
		c = $('<style id="' + d + '">').appendTo("body")
	}
	var a = c.get(0);
	if (a.styleSheet) {
		a.styleSheet.cssText = e
	} else {
		a.innerHTML = e
	}
}
function showBrowserNotice() {
	userAgent = navigator.userAgent.toLowerCase();
	$browser = new Object();
	$browser.msie = /msie/.test(userAgent);
	$browser.chrome = /chrome/.test(userAgent);
	var a = false;
	if (navigator.userAgent.indexOf("MetaSr") >= 0) {
		a = true
	} else {
		if (navigator.userAgent.indexOf("LBBROWSER") >= 0) {
			a = true
		} else {
			if (navigator.userAgent.indexOf("QQBrowser") >= 0) {
				a = true
			} else {
				if (navigator.userAgent.indexOf("TheWorld") >= 0) {
					a = true
				} else {
					if (navigator.userAgent.indexOf("BIDUBrowser") >= 0) {
						a = true
					} else {
						if (navigator.userAgent.indexOf("Maxthon") >= 0) {
							a = true
						}
					}
				}
			}
		}
	}
	if (a) {
		$("body").prepend('<div class="alert alert-info alert-dismissable" style="margin:0px;"><button type=button" onclick="ajaxIgnoreBrowser()" class="close" data-dismiss="alert" aria-hidden="true"><i class="icon-remove"></i></button><p>' + browserNotice + "</p></div>")
	}
}
function removeCookieByKey(a) {
	$.cookie(a, "", {
		expires: config.cookieLife,
		path: config.webRoot
	});
	location.href = location.href
}
function setHomepage(a, b) {
	$.get(createLink("custom", "ajaxSetHomepage", "module=" + a + "&page=" + b), function() {
		location.reload(true)
	})
}
function checkTutorial() {
	if (config.currentModule != "tutorial" && window.TUTORIAL && (!frameElement || frameElement.tagName != "IFRAME")) {
		if (confirm(window.TUTORIAL.tip)) {
			$.getJSON(createLink("tutorial", "ajaxQuit"), function() {
				window.location.reload()
			}).error(function() {
				alert(lang.timeout)
			})
		}
	}
}
function removeDitto() {
	$firstTr = $(".table-form").find("tbody tr:first");
	$firstTr.find("td select").each(function() {
		$(this).find("option[value='ditto']").remove();
		$(this).trigger("chosen:updated")
	})
}
function revertModuleCookie() {
	if ($('#mainmenu .nav li[data-id="project"]').hasClass("active")) {
		$('#modulemenu .nav li[data-id="task"] a').click(function() {
			$.cookie("moduleBrowseParam", 0, {
				expires: config.cookieLife,
				path: config.webRoot
			})
		})
	}
	if ($('#mainmenu .nav li[data-id="product"]').hasClass("active")) {
		$('#modulemenu .nav li[data-id="story"] a').click(function() {
			$.cookie("storyModule", 0, {
				expires: config.cookieLife,
				path: config.webRoot
			})
		})
	}
	if ($('#mainmenu .nav li[data-id="qa"]').hasClass("active")) {
		$('#modulemenu .nav li[data-id="bug"] a').click(function() {
			$.cookie("bugModule", 0, {
				expires: config.cookieLife,
				path: config.webRoot
			})
		});
		$('#modulemenu .nav li[data-id="testcase"] a').click(function() {
			$.cookie("caseModule", 0, {
				expires: config.cookieLife,
				path: config.webRoot
			})
		})
	}
}
function inputFocusJump(d) {
	var f = $("input").is(":focus");
	if (f) {
		var e = $("input:focus").attr("name").replace(/\[\d]/g, "");
		var g = $(":input[name^=" + e + "]:text:not(:disabled):not([name*='%'])");
		var c = g.length;
		var b = parseInt($("input:focus").attr("name").replace(/[^0-9]/g, ""));
		var a = d == "down" ? b + 1 : b - 1;
		if (a < c && a >= 0) {
			g[a].focus()
		}
	}
}
function selectFocusJump(e) {
	var g = $("select").is(":focus");
	if (g) {
		var f = $("select:focus").attr("name").replace(/\[\d]/g, "");
		var d = $("select[name^=" + f + "]:not([name*='%'])");
		var c = d.length;
		var b = parseInt($("select:focus").attr("name").replace(/[^0-9]/g, ""));
		var a = e == "down" ? b + 1 : b - 1;
		if (a < c && a >= 0) {
			d[a].focus()
		}
	}
}
function adjustNoticePosition() {
	var a = 25;
	$("#noticeBox").find(".alert").each(function() {
		$(this).css("bottom", a + "px");
		a += $(this).outerHeight(true) - 10
	})
}
function notifyMessage(a) {
	if (window.Notification) {
		if (Notification.permission == "granted") {
			new Notification("", {
				body: a
			})
		} else {
			if (Notification.permission != "denied") {
				Notification.requestPermission(function(b) {
					new Notification("", {
						body: a
					})
				})
			}
		}
	}
}
function getFingerprint() {
	if (typeof(Fingerprint) == "function") {
		return new Fingerprint().get()
	}
	fingerprint = "";
	$.each(navigator, function(a, b) {
		if (typeof(b) == "string") {
			fingerprint += b.length
		}
	});
	return fingerprint
}
needPing = true;
$(document).ready(function() {
	if (needPing) {
		setTimeout("setPing()", 1000 * 60 * 10)
	}
	checkTutorial();
	revertModuleCookie();
	$(document).on("click", "#helpMenuItem .close-help-tab", function() {
		$("#helpMenuItem").prev().remove();
		$("#helpMenuItem").remove()
	})
});
/*!
 * jQuery Cookie Plugin v1.4.1
 * https://github.com/carhartl/jquery-cookie
 * Copyright 2006, 2014 Klaus Hartl
 * Released under the MIT license
 */

!function(a) {
	"function" == typeof define && define.amd ? define(["jquery"], a) : a("object" == typeof exports ? require("jquery") : jQuery)
}(function(d) {
	function h(a) {
		return g.raw ? a : encodeURIComponent(a)
	}
	function c(a) {
		return g.raw ? a : decodeURIComponent(a)
	}
	function k(a) {
		return h(g.json ? JSON.stringify(a) : String(a))
	}
	function j(a) {
		0 === a.indexOf('"') && (a = a.slice(1, -1).replace(/\\"/g, '"').replace(/\\\\/g, "\\"));
		try {
			return a = decodeURIComponent(a.replace(b, " ")), g.json ? JSON.parse(a) : a
		} catch (l) {}
	}
	function f(l, a) {
		var m = g.raw ? l : j(l);
		return d.isFunction(a) ? a(m) : m
	}
	var b = /\+/g,
		g = d.cookie = function(n, A, r) {
			if (void 0 !== A && !d.isFunction(A)) {
				if (r = d.extend({}, g.defaults, r), "number" == typeof r.expires) {
					var s = r.expires,
						z = r.expires = new Date;
					z.setTime(+z + 86400000 * s)
				}
				return document.cookie = [h(n), "=", k(A), r.expires ? "; expires=" + r.expires.toUTCString() : "", r.path ? "; path=" + r.path : "", r.domain ? "; domain=" + r.domain : "", r.secure ? "; secure" : ""].join("")
			}
			for (var y = n ? void 0 : {}, C = document.cookie ? document.cookie.split("; ") : [], e = 0, x = C.length; e < x; e++) {
				var w = C[e].split("="),
					q = c(w.shift()),
					B = w.join("=");
				if (n && n === q) {
					y = f(B, A);
					break
				}
				n || void 0 === (B = f(B)) || (y[q] = B)
			}
			return y
		};
	g.defaults = {}, d.removeCookie = function(l, a) {
		return void 0 !== d.cookie(l) && (d.cookie(l, "", d.extend({}, a, {
			expires: -1
		})), !d.cookie(l))
	}
}), function(d, g) {
	var c, j, h = "localStorage",
		f = "page_" + d.location.pathname + d.location.search,
		b = function() {
			this.slience = !0;
			try {
				h in d && d[h] && d[h].setItem && (this.enable = !0, c = d[h])
			} catch (e) {}
			this.enable || (j = {}, c = {
				getLength: function() {
					var a = 0;
					return g.each(j, function() {
						a++
					}), a
				},
				key: function(k) {
					var a, l = 0;
					return g.each(j, function(m) {
						return l === k ? (a = m, !1) : void l++
					}), a
				},
				removeItem: function(a) {
					delete j[a]
				},
				getItem: function(a) {
					return j[a]
				},
				setItem: function(a, k) {
					j[a] = k
				},
				clear: function() {
					j = {}
				}
			}), this.storage = c, this.page = this.get(f, {})
		};
	b.prototype.pageSave = function() {
		if (g.isEmptyObject(this.page)) {
			this.remove(f)
		} else {
			var e, a = [];
			for (e in this.page) {
				var k = this.page[e];
				null === k && a.push(e)
			}
			for (e = a.length - 1; e >= 0; e--) {
				delete this.page[a[e]]
			}
			this.set(f, this.page)
		}
	}, b.prototype.pageRemove = function(a) {
		"undefined" != typeof this.page[a] && (this.page[a] = null, this.pageSave())
	}, b.prototype.pageClear = function() {
		this.page = {}, this.pageSave()
	}, b.prototype.pageGet = function(k, l) {
		var a = this.page[k];
		return void 0 === l || null !== a && void 0 !== a ? a : l
	}, b.prototype.pageSet = function(e, a) {
		g.isPlainObject(e) ? g.extend(!0, this.page, e) : this.page[this.serialize(e)] = a, this.pageSave()
	}, b.prototype.check = function() {
		if (!this.enable && !this.slience) {
			throw new Error("Browser not support localStorage or enable status been set true.")
		}
		return this.enable
	}, b.prototype.length = function() {
		return this.check() ? c.getLength ? c.getLength() : c.length : 0
	}, b.prototype.removeItem = function(a) {
		return c.removeItem(a), this
	}, b.prototype.remove = function(a) {
		return this.removeItem(a)
	}, b.prototype.getItem = function(a) {
		return c.getItem(a)
	}, b.prototype.get = function(k, l) {
		var a = this.deserialize(this.getItem(k));
		return "undefined" != typeof a && null !== a || "undefined" == typeof l ? a : l
	}, b.prototype.key = function(a) {
		return c.key(a)
	}, b.prototype.setItem = function(a, k) {
		return c.setItem(a, k), this
	}, b.prototype.set = function(a, k) {
		return void 0 === k ? this.remove(a) : (this.setItem(a, this.serialize(k)), this)
	}, b.prototype.clear = function() {
		return c.clear(), this
	}, b.prototype.forEach = function(a) {
		for (var k = this.length(), m = k - 1; m >= 0; m--) {
			var l = c.key(m);
			a(l, this.get(l))
		}
		return this
	}, b.prototype.getAll = function() {
		var a = {};
		return this.forEach(function(l, k) {
			a[l] = k
		}), a
	}, b.prototype.serialize = function(a) {
		return "string" == typeof a ? a : JSON.stringify(a)
	}, b.prototype.deserialize = function(a) {
		if ("string" == typeof a) {
			try {
				return JSON.parse(a)
			} catch (k) {
				return a || void 0
			}
		}
	}, g.zui({
		store: new b
	})
}(window, jQuery);