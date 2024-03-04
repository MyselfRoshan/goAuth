const email = document.getElementById("email");
const password = document.getElementById("password");
const submit = document.getElementById("submit");
const loginForm = document.getElementById("loginForm");
// console.log(email, password, submit);

loginForm.addEventListener("submit", async e => {
  e.preventDefault();
  const formData = new FormData(e.target);
  const data = Object.fromEntries(formData);
  console.log(data);

  const { message } = await ajax("/login", "post", data);
  message === "success" ? window.location.assign("/") : alert(message);
});
