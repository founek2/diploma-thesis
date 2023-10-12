export function flatten(arr){return arr.flat(10)}
export function buildTag(tagName,children,options={}){const tag=document.createElement(tagName)
Object.entries(options).forEach(([key,val])=>{if(key==="actions"){Object.entries(val).forEach(([event,fn])=>{tag.addEventListener(event,fn)})}else if(val!==null)tag.setAttribute(key,val)})
if(Array.isArray(children))
flatten(children).map(el=>typeof el==="string"?document.createTextNode(el):el).forEach(tag.appendChild.bind(tag))
else tag.appendChild(document.createTextNode(children))
return tag}
export function isVersionGreater(v1,v2){const[major1,minor1,patch1]=v1.replace('v','').split('.').map(Number);const[major2,minor2,patch2]=v2.replace('v','').split('.').map(Number);if(major1>major2){return true;}
if(major1<major2){return false;}
if(minor1>minor2){return true;}
if(minor1<minor2){return false;}
if(patch1>patch2){return true;}
if(patch1<=patch2){return false;}}
export function toDate(date){const d=new Date(date);return d.toISOString().split('T')[0];}
export function insertAfter(newNode,existingNode){if(Array.isArray(newNode)){newNode.forEach(el=>existingNode.parentNode.insertBefore(el,existingNode.nextSibling))}else existingNode.parentNode.insertBefore(newNode,existingNode.nextSibling)}