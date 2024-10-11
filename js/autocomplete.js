const searchWrapper = document.querySelector(".search-form");
const inputBox = searchWrapper.querySelector("#search");
const suggBox = searchWrapper.querySelector(".suggestions");

let timer;

function debouncing(keypressed) {
  if (timer) clearTimeout(timer);

  timer = setTimeout(() => {
    // POST request to the server with the user's search key
    fetch("http://localhost:8089/searchy", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ Key: keypressed }),
    })
      .then((response) => response.json())
      .then((data) => {
        if (keypressed != "" && data != null) {
          if (data.length !== 0) {
            // Map filtered suggestions to HTML
            data = data.map((item) => {
              return item["contents"] === "artist/band"
                ? `<li><a class="suggestion-item" href="/artist?id=${item["id"][0]}">${item["name"]} - <span class="badge">${item["contents"]}</span></a></li>`
                : item["contents"] === "member"
                ? `<li><a class="suggestion-item" href="/serch?search=${item["id"][0]}-bandmember-${item["name"]}">${item["name"]} - <span class="badge">${item["contents"]}</span></a></li>`
                : `<li><a class="suggestion-item" href="/serch?search=${item["name"]}">${item["name"]} - <span class="badge">${item["contents"]}</span></a></li>`;
            });
            showSuggestions(data); // Show filtered suggestions
          } else {
            suggBox.innerHTML = ""; // Clear suggestions if none match
          }
        } else {
          suggBox.innerHTML = ""; // Clear suggestions if input is empty
        }
      })
      .catch((error) => console.log("Error fetching data:", error));
  }, 300);
}

// Function to display suggestions in the UI
function showSuggestions(list) {
  let listData = list.join(""); // Join list items into a single string
  suggBox.innerHTML = listData; // Update suggestions box with list data
}

// If the user presses any key and releases
inputBox.onkeyup = (e) => {
  let userData = e.target.value.trim();
  if (userData === "") {
    suggBox.innerHTML = "";
  }
  debouncing(userData);
};
