document.addEventListener('DOMContentLoaded', function() {
	TableOfContents();
}
);

function TableOfContents(container, output) {
	var toc = "";
	var level = 0;
	var container = document.querySelector(container) || document.querySelector('#contents');
	var output = output || '#toc';


	container.innerHTML =
	container.innerHTML.replace(
			/<h([\d])>([^<]+)<\/h([\d])>/gi,
			function (str, openLevel, titleText, closeLevel) {
					if (openLevel != closeLevel || openLevel >= 4) {
							return str;
					}

					if (openLevel > level) {
							toc += (new Array(openLevel - level + 1)).join('<ul>');
					} else if (openLevel < level) {
							toc += (new Array(level - openLevel + 1)).join('</li></ul>');
					} else {
							toc += (new Array(level+ 1)).join('</li>');
					}

					level = parseInt(openLevel);

					var anchor = titleText.replace(/ /g, "_");
					toc += '<li><a href="#' + anchor + '">' + titleText
							+ '</a>';

					return '<h' + openLevel + '><a href="#' + anchor + '" id="' + anchor + '">'
							+ titleText + '</a></h' + closeLevel + '>';
			}
	);




	// var reg = /<h([\d]) id\=\"(.+)\">([^<]+)<\/h([\d])>/gi;
	// var result;
	// while((result = reg.exec(container.innerHTML)) !== null) {
	// 	openLevel = result[1]
	// 	idLevel = result[2]
	// 	titleText = result[3]
	// 	closeLevel = result[4]

	// 	if (openLevel != closeLevel || openLevel >= 4) {
	// 		continue;
	// 	}

	// 	if (openLevel > level) {
	// 		toc += (new Array(openLevel - level + 1)).join('<ul>');
	// 	} else if (openLevel < level) {
	// 		toc += (new Array(level - openLevel + 1)).join('</li></ul>');
	// 	} else {
	// 		toc += (new Array(level+ 1)).join('</li>');
	// 	}

	// 	level = parseInt(openLevel);

	// 	toc += '<li><a href="#' + idLevel + '">' + titleText
	// 	+ '</a>';
	// }

	if (level) {
		toc += (new Array(level + 1)).join('</ul>');
	}
	document.querySelector(output).innerHTML += toc;
};
