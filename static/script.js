var listBTN = document.getElementById("list-btn");
var submitBTN = document.getElementById("submit-btn");
var movieForm =  document.getElementById("movie-form")
submitBTN.addEventListener("click", submitMovies)
listBTN.addEventListener("click", listMovies);

const address = "http://localhost:8000"

function listMovies() {
    getMovies(address+"/movies")
};

function submitMovies(event) {
    event.preventDefault()
    addMovie(address+"/movies")
};
// Add an event listener to the form's submit event To prevent the default form submission behavior .
movieForm.addEventListener("submit", (event) => {
    if (event.target.id == "submit-btn") {
        event.preventDefault()
        submitMovies(event)
    }
})

async function checkURLStatus(url) {
    const response = await fetch(url);

    if (!response.ok) {
        throw new Error(`Network response was not okay: ${response.statusText}`)
    };

    return response;
}

async function addMovie(url) {
    try {
        const response = checkURLStatus(url)

        //  const formData = new FormData(movieForm)
        //  formData.forEach((val, key) => {
        //     console.log(val,key)
        //  })
        //  console.log("FormData", formData)

        // var id = document.getElementById("id").value;
        var isbn = document.getElementById("isbn").value;
        var title = document.getElementById("title").value;
        var firstname = document.getElementById("dir-firstname").value;
        var lastname = document.getElementById("dir-lastname").value;

        var formData = {
            // id,
            isbn,
            title,
            director: {
                firstname,
                lastname
            }
        }
        
        await fetch(url, {
            method: "POST",
            body: JSON.stringify(formData),
        })

        console.log("Success sending request", response)
        
    } catch (error) {
        console.error(`Error fetching data: ${error}`)
    }

    
};

async function getMovies(url) {
    try {

       const response = await fetch(url,{
        method: "GET"
       })

       if (!response.ok){
        throw new Error(`Network response was not okay: ${response.statusText}`)
       }

       const responseBody = await response.json()
       var container = document.querySelector(".movie-container")
       container.innerHTML = ""

       responseBody.forEach((item) => {
           const itemContainer = document.createElement("div")


        const itemHTML = `
        <li class="title" >Title: ${item.title}</li>
        <ul class="other-info">
            <li>ID: ${item.id}</li>
            <li>ISBN: ${item.isbn}</li>
            <li>Director: ${item.director.firstname} ${item.director.lastname}</li>
        </ul>
        `;
        itemContainer.innerHTML = itemHTML
        container.appendChild(itemContainer)
       })
        
    } catch (error) {
        console.error(`Error fetching data ${error}`)
    }
};