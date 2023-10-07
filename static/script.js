var listBTN = document.getElementById("list-btn");
var submitBTN = document.getElementById("submit-btn");
var movieForm =  document.getElementById("movie-form")
submitBTN.addEventListener("click", submitMovies)
listBTN.addEventListener("click", listMovies);

function listMovies() {
    getMovies()
};

function submitMovies(event) {
    event.preventDefault()
    addMovie()
};
// Add an event listener to the form's submit event To prevent the default form submission behavior .
movieForm.addEventListener("submit", (event) => {
    if (event.target.id == "submit-btn") {
        event.preventDefault()
        submitMovies(event)
    }
})

async function addMovie() {
    var url = "http://localhost:8000/movies"

    try {
        const response = await fetch(url);

        if (!response.ok) {
            throw new Error(`Network response was not okay: ${response.statusText}`)
        }

        //  const formData = new FormData(movieForm)
        //  formData.forEach((val, key) => {
        //     console.log(val,key)
        //  })
        //  console.log("FormData", formData)

        var id = document.getElementById("id").value;
        var isbn = document.getElementById("isbn").value;
        var title = document.getElementById("title").value;
        var firstname = document.getElementById("dir-firstname").value;
        var lastname = document.getElementById("dir-lastname").value;

        var formData = {
            id,
            isbn,
            title,
            director: {
                firstname,
                lastname
            }
        }
        
        await fetch(url, {
            method: "POST",
            body: formData
        })

        console.log("Success sending request")
        
    } catch (error) {
        console.error(`Error fetching data: ${error}`)
    }

    
};

function getMovies() {
    
};