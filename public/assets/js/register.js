const name = document.getElementById("name");
const email = document.getElementById("email");
const password = document.getElementById("password");
const registerForm = document.getElementById("registerForm");

registerForm.addEventListener("submit", async e => {
  e.preventDefault();
  const formData = new FormData(e.target);
  const data = Object.fromEntries(formData);

  const { id, name, email } = await ajax("/register", "post", data);
  if (id === 0) {
    alert("User with this email already exists");
  }
  if (id > 0) {
    alert(
      `User with\nEmail: ${email}\nName: ${name}\nregistered successfully\nRedirecting to login page... in 5 seconds...`,
    );
    window.location.assign("/login");
  }
});
