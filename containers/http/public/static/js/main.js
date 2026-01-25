// global htmx event handlers
document.addEventListener("DOMContentLoaded", (even) => {
  document.body.addEventListener("htmx:beforeSwap", function (evt) {
    if (
      evt.detail.xhr.status === 400 ||
      evt.detail.xhr.status === 401 ||
      evt.detail.xhr.status === 422
    ) {
      evt.detail.shouldSwap = true;
      evt.detail.isError = false;
    }
  });
});
