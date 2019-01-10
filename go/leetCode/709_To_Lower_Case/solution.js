/*
* @Author: wcy
* @Date:   2018-07-22 15:14:18
* @Last Modified by:   wcy
* @Last Modified time: 2018-07-22 15:22:39
*/
let toLowerCase_1 = function(str) {
	let result = "";
	for(let item of str) {
		if (item.charCodeAt() >= 65 && item.charCodeAt() <= 90) {
			result += String.fromCharCode(item.charCodeAt() + 32);
			continue;
		}
		result += item;
	}
	return result;
}

console.log(toLowerCase_1("asdaA"))