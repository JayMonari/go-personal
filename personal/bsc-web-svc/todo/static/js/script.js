for (const item of document.getElementsByTagName("li")) {
  item.addEventListener("click", () => item.classList.toggle("done"));
}
