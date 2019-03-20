# Elimity Developer Intern Interview Exercise

## Welcome

Hi, and thanks for applying for the __Elimity Developer Intern__ position with the Elimity Engineering team.

This exercise is designed to give you an opportunity to show off programming skills that would be relevant to work on [Elimity](https://www.elimity.com/).

## Exercise instructions
Please use this repository to develop a simple application that allows users to view and filter a list of users.
Depending on the position you applied for, read the instructions in the applicable section below (e.g., if you want to work on the frontend, read the instructions below [Frontend](#frontend)).

The installation instructions for the backend and the frontend can been seen in their respective README's: `./backend/README.md` and `./frontend/README.md`.

### Frontend
- Use the backend REST API provided in `./backend` to fetch all users instead of using the mock data in `./frontend/MOCK_DATA.json`.
- Add a search field in the material UI to filter users client-side.
- Style data table to use the full width of the screen.
- Include a git repo in `frontend/`.

Useful resources:
- [angular material components](https://material.angular.io/components/categories)
- [`mat-table`](https://material.angular.io/components/table/overview)
- [Angular HttpClient](https://angular.io/guide/http)

### Backend
- Add a parameter to the provided REST call to allow case-insensitive filtering of users by name.
- Add integration tests at HTTP level to verify this functionality.
- Include a git repo in `backend/`.

Useful resources:

- [Testify](http://godoc.org/github.com/stretchr/testify) (the [suite package](http://godoc.org/github.com/stretchr/testify/suite) is something we use at Elimity for example)
- [Resty](https://github.com/go-resty/resty)


### Full stack
- Use the backend REST API provided in `./backend` to fetch all users instead of using the mock data in `./frontend/MOCK_DATA.json`.
- Add a parameter to the provided backend REST call to allow sorting (ascending/descending) by name.
- Add a checkbox to the material UI for ascending/descending of user results and use this in the REST call.
- Include a git repo in the root folder: `elimity-intern-exercise/`.

Useful resources:
- [angular material components](https://material.angular.io/components/categories)
- [`mat-table`](https://material.angular.io/components/table/overview)
- [Angular HttpClient](https://angular.io/guide/http)


## Evaluation

Your submission will be evaluated against the following criteria:

* Meets the stated requirements
* Code is high quality and well organized
* Best practices are followed including:
  * Version control
  * Iterative development
  * ...

## Submission

To submit your work, email us a zip of your code. Share your zip file via wetransfer (google will block your mail if you attach the zip directly). Make sure to include the `.git` folder so we can look at your git history.
