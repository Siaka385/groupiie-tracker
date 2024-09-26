
const searchWrapper = document.querySelector(".search-form");
const inputBox = searchWrapper.querySelector("#search");
const suggBox = searchWrapper.querySelector(".suggestions");

// If the user presses any key and releases
inputBox.onkeyup = (e) => {
       
 
  let userData = e.target.value.trim(); // Get the input value and trim whitespace

  // POST request to the server with the user's search key
  fetch("http://localhost:8089/searchy", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ Key: userData }),
  })
    .then((response) => response.json())
    .then((data) => {
      if (userData !="" && data != null) {
      
        if (data.length !== 0) {
          // Map filtered suggestions to HTML
          data= data.map((item) => {
            return item["contents"] === "artist/band" || item["contents"] === "member"
              ? `<li><a class="suggestion-item" href="/artist?id=${item["id"][0]}">${item["name"]} <span class="badge">${item["contents"]}</span></a></li>`
              : `<li><a class="suggestion-item" href="/serch?search=${item["name"]}">${item["name"]} <span class="badge">${item["contents"]}</span></a></li>`;
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


};

// Function to display suggestions in the UI
function showSuggestions(list) {
  let listData = list.join(""); // Join list items into a single string
  suggBox.innerHTML = listData; // Update suggestions box with list data
}
