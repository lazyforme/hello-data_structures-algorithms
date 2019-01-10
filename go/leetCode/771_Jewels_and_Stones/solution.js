/*
 * @Author: wcy
 * @Date:   2018-07-21 21:15:25
 * @Last Modified by:   wcy
 * @Last Modified time: 2018-07-21 22:28:56
 */

let numJewelsInStones_1 = function(J, S) {
	let result = 0;
	for (let j of J) {
		for (let s of S) {
			if (j !== s) {
				continue;
			}
			result++;
		}
	}
	return result;
}

let numJewelsInStones_2 = function(J, S) {
	let result = 0;
	for (let s of S) {
		if (J.indexOf(s) < 0) {
			continue;
		}
		result++;
	}
	return result;
}

let J = 'Aads';
let S = 'aAAsdfsgfsg';

console.log(numJewelsInStones_3(J, S))