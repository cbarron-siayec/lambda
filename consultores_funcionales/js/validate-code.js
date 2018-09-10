$(document).ready(function() {
  function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
  }

  async function demo() {
    await sleep(7000);
    window.location.href = "../index.html";
  }

  demo();
});
