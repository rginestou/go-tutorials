window.addEventListener('scroll', function (e) {
	let pos = window.scrollY;
	let start = 280
	let end = 460
	let opacity = pos / (end - start) - start / (end - start)

	let toc = document.getElementById("toc")

	if (opacity > 1) opacity = 1
	if (opacity < 0) {
		opacity = 0
		toc.style.display = "none"
	} else {
		toc.style.display = "block"
		toc.style.opacity = opacity;
	}
})
