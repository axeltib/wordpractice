
const words = [];
words.push("하다", "있다", "되다", "수", "것")

var curW = 0

function updatePromptWord(text) {
	document.getElementById("promptWord").textContent = text;
}
updatePromptWord(words[0])


document.getElementById("senditButton").onclick = function() {
	var input = document.getElementById("answerBox").value;
	console.log(input);
	curW = curW + 1;
	updatePromptWord(words[curW]);
}

