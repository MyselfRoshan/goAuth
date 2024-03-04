async function ajax(url, method = "get", body = {}) {
  method = method.toUpperCase();

  let options = {
    method: method,
    headers: {
      "Content-Type": "application/json",
      "X-Requested-With": "XMLHttpRequest",
    },
  };

  const csrfMethods = new Set(["POST", "PUT", "DELETE", "PATCH"]);

  if (csrfMethods.has(method)) {
    // options.body = body;
    options.body = JSON.stringify(body);
  } else if (method === "GET") {
    url += `?${new URLSearchParams(body).toString()}`;
  }
  //   return await fetch(url, options);
  //   try {
  const response = await fetch(url, options);
  const data = await response.json();
  console.log(data);
  return data;
  //   } catch (error) {
  //     console.error(`Error setting data:`, error);
  //     return null;
  //   }
}
