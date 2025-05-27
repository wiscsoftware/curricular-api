# Design and Implementation of a REST API for Curricular data in Higher Education

The design and implementation of a REST API for student and course data for a Higher Education institution.This shows 
how to start with users' needs (user stories) and use that to design
the API specification, and finally the implementation. This API design focus on some of the key resources such as
`/students`, `/teachers`, `/classes`
and can be extended to include other resources such as `/universities` (if the institution consists of multiple
universities such as [Universities of Wisconsin](https://www.wisconsin.edu/), that
has [13 universities](https://www.wisconsin.edu/campuses/) or [University of California](https://www.universityofcalifornia.edu/) 
that consist of [10 campuses](https://www.universityofcalifornia.edu/campuses-locations)), `/enrollments`, `/schools`, `/courses` etc..

Feel free to reach us at wiscsoftware@gmail.com to see how we can collaborate in your API design and implementation
effort.

## Tech Stack 

Technologies used in the implementation. `Go` language was used considering its support for building low cost and maintainable
Cloud Native apps in AWS. `SQLite` was used because this is a high transaction system (500K requests/second) 
with mostly read-only data (write-once a day and read-many times a day). Note that the implementation here uses a mock 
datastore([store.go](src/api/store.go)) 


* ![Docker](https://img.shields.io/badge/-Docker-eee?style=flat-square&logo=Docker&logoColor=2496ED)
* ![Go](https://img.shields.io/badge/-Go-eee?style=flat-square&logo=Go&logoColor=00ADD8)
* ![SQLite](https://img.shields.io/badge/-SQLite-eee?style=flat-square&logo=SQLite&logoColor=003B57)
* ![JSON](https://img.shields.io/badge/-JSON-eee?style=flat-square&logo=JSON&logoColor=000000)
* ![Terraform](https://img.shields.io/badge/-Terraform-eee?style=flat-square&logo=Terraform&logoColor=844FBA)
* ![AWSS3](https://img.shields.io/badge/-AWS%20S3-eee?style=flat-square&logo=amazon-s3&logoColor=569A31)
* ![AWSDynamodDB](https://img.shields.io/badge/-AWS%20DynamoDB-eee?style=flat-square&logo=amazon-dynamodb&logoColor=4053D6)
* ![AWSECS](https://img.shields.io/badge/-AWS%20ECS-eee?style=flat-square&logo=amazon-ECS&logoColor=FF9900)
* ![AWSCW](https://img.shields.io/badge/-AWS%20CloudWatch-eee?style=flat-square&logo=amazon-CloudWatch&logoColor=FF4F8B)
* ![AWSAPIGW](https://img.shields.io/badge/-AWS%20APIGateway-eee?style=flat-square&logo=amazon-API-Gateway&logoColor=FF4F8B)
* <img src="https://www.vectorlogo.zone/logos/apigee/apigee-ar21.svg" align="center" alt="Apigee" width="100">


## System Design

[Design of a REST API for Curricular data](https://github.com/baranasoftware/system-design/blob/main/edu-api.md)

## Running Locally
```
go run main.go  --local=true --port=8080
Curricular API server listing on port: 8080
```

Registering a client. Returns `clientId` and `secret` for subsequent requests:
```
$ curl -X GET localhost:8000/register
{"clientId":"db72caf4-5002-4c54-967e-55dba9a2704e","secret":"hztUn9rROthCb8gmdtt4gshSlgdaXTuN"}
```
    
Authorize the client:
```
$ curl -X GET 'http://localhost:8080/authorize?client_id=db72caf4-5002-4c54-967e-55dba9a2704e&response_type=code'

```

Get the token:
```
$ curl -X GET 'http://localhost:8080/oauth/token?grant_type=client_credentials&client_id=db72caf4-5002-4c54-967e-55dba9a2704e&client_secret=hztUn9rROthCb8gmdtt4gshSlgdaXTuN&scope=read'
{"access_token":"ZJU3NDNMOWUTOGE3YY0ZNTM5LWFJNZYTZDFJNTDMYJQWZJM0","expires_in":7200,"scope":"read","token_type":"Bearer"}
```

Access the resource `/students`:
```
$ curl -X GET -H 'Authorization: Bearer ZJU3NDNMOWUTOGE3YY0ZNTM5LWFJNZYTZDFJNTDMYJQWZJM0' localhost:8000/students
[
  {
    "identities": [
      {
        "name": "emplId",
        "value": "cu6n5r8tnuc5dvfiniug"
      },
      {
        "name": "libraryId",
        "value": "cu6n5r8tnuc5dvfiniv0"
      },
      {
        "name": "campusId",
        "value": "cu6n5r8tnuc5dvfinivg"
      }
    ],
    "firstName": "Tierra",
    "lastName": "Rice",
    "addresses": [
      {
        "addressLine1": "89876 Mountburgh",
        "addressLine2": "",
        "city": "St. Louis",
        "state": "Arkansas",
        "country": "Saint Pierre and Miquelon",
        "zipCode": "75060"
      },
      {
        "addressLine1": "6835 Locksmouth",
        "addressLine2": "",
        "city": "St. Louis",
        "state": "New Mexico",
        "country": "Bhutan",
        "zipCode": "44661"
      }
    ],
    "birthdate": "2011-03-29T10:26:47.533935787Z",
    "ageInYears": 14,
    "residencyStatus": 2
  },
  {
    "identities": [
      {
        "name": "campusId",
        "value": "cu6n5r8tnuc5dvfinj20"
      },
      {
        "name": "emplId",
        "value": "cu6n5r8tnuc5dvfinj2g"
      },
      {
        "name": "libraryId",
        "value": "cu6n5r8tnuc5dvfinj30"
      }
    ],
    "firstName": "Abner",
    "lastName": "Hansen",
    "addresses": [
      {
        "addressLine1": "225 Villagetown",
        "addressLine2": "",
        "city": "Oklahoma",
        "state": "Tennessee",
        "country": "Paraguay",
        "zipCode": "76888"
      },
      {
        "addressLine1": "1841 Lake Fieldmouth",
        "addressLine2": "",
        "city": "Dallas",
        "state": "Vermont",
        "country": "Liechtenstein",
        "zipCode": "67213"
      }
    ],
    "birthdate": "1999-10-16T09:29:20.221982237Z",
    "ageInYears": 26,
    "residencyStatus": 2
  },
  {
    "identities": [
      {
        "name": "campusId",
        "value": "cu6n5r8tnuc5dvfinj5g"
      },
      {
        "name": "emplId",
        "value": "cu6n5r8tnuc5dvfinj60"
      },
      {
        "name": "libraryId",
        "value": "cu6n5r8tnuc5dvfinj6g"
      }
    ],
    "firstName": "Carolyne",
    "lastName": "Eichmann",
    "addresses": [
      {
        "addressLine1": "26219 Cornersport",
        "addressLine2": "",
        "city": "Orlando",
        "state": "Alabama",
        "country": "Slovakia",
        "zipCode": "84322"
      },
      {
        "addressLine1": "457 Passchester",
        "addressLine2": "",
        "city": "San Jose",
        "state": "Iowa",
        "country": "Somalia",
        "zipCode": "16710"
      }
    ],
    "birthdate": "1951-12-13T05:20:22.497091342Z",
    "ageInYears": 74,
    "residencyStatus": 2
  }
]
```
             
## AWS Resources 
AWS resources were created with [Terraform](./terraform) [using](./Makefile):
```
cd src; GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bootstrap main.go; \
cd ../terraform; \
AWS_PROFILE=<aws-profile> AWS_REGION=us-east-1 terraform init  -var aws_account_ids=[aws-account-ids] -var sec_group=<security-group-id>;\
AWS_PROFILE=<aws-profile> AWS_REGION=us-east-1 terraform plan  -var aws_account_ids=[aws-account-ids] -var sec_group=<security-group-id>;\
AWS_PROFILE=<aws-profile> AWS_REGION=us-east-1 terraform apply -var aws_account_ids=[aws-account-ids] -var sec_group=<security-group-id>;
```

## Invoking API from APIGateway
```
curl -X GET -H 'X-API-Key: <API-Key>' https://<api-id>.execute-api.us-east-1.amazonaws.com/curricular-and-academic-api/students/ 
```

## API Design

API design was done
using [Align-Define-Design Process](https://blog.stoplight.io/aligning-on-your-api-design-using-jobs-to-be-done).

### User stories
| Story ID | When... (Triggering Situation)              | I want to...    (Digital Capability)                   | So I can...    (Outcome)                                    |
|----------|---------------------------------------------|--------------------------------------------------------|-------------------------------------------------------------|
| 1        | I want to find students                     | Search students by student ID, first name and lastname | Confirm their details and set up an appointment             |
| 2        | I want to find teachers                     | Search teachers by emplID, first name and lastname     | Confirm their details                                       |
| 3        | I want to find students for a teacher       | View number of students for a teacher                  | Determine if the class size is too big                      |
| 5        | I want to find more details about the class | Book an appointment                                    | So I can set up an appointment to discuss course assignment |
| 4        | I want to enroll in a class                 | Search for a class                                     | Confirm that's the class I need to enroll in                |

### Activities
| Digital Capability                                     | Activity                    | Participants        | Description                                            |
|--------------------------------------------------------|-----------------------------|---------------------|--------------------------------------------------------|
| Search Students by student ID, first name and lastname | Search Students             | Teacher, Admin User | Search for students by student Id, firstname, lastname |
| Search teachers by emplID, first name and lastname     | Search Teachers             | Teacher, Admin User | Search for teachers by empl Id, firstname, lastname    |
| View number of students for a teacher                  | View Students for Teacher   | Teacher, Admin User | View students for the teacher                          |
| Book an appointment                                    | View Classes                | Student             | Search classes by class number, name                   |                        |
| Book an appointment                                    | View Teachers for the Class | Student             | Search for a teacher by class                          |

### Activity Steps
| Digital Capability                                     | Activity                    | Activity Step              | Participants        | Description                                            |
|--------------------------------------------------------|-----------------------------|----------------------------|---------------------|--------------------------------------------------------|
| Search Students by student ID, first name and lastname | Search Students             | Search Students            | Teacher, Admin User | Search for students by student Id, firstname, lastname |
| Search teachers by emplID, first name and lastname     | Search Teachers             | Search Teachers            | Teacher, Admin User | Search for teachers by empl Id, firstname, lastname    |
| View number of students for a teacher                  | View Students for Teacher   | View Teachers              | Teacher, Admin User | Search for teachers by empl Id, firstname, lastname    |
| View number of students for a teacher                  | Search Students for Teacher | View Students for Teacher  | Teacher, Admin User | View students for the teacher                          |
| Book an appointment                                    | View Classes                | View Classes               | Student             | Search classes by class number, name                   |
| Book an appointment                                    | View Teachers for the Class | View teacher for the class | Student             | Search class by teacher                                |

### API Resources and Models

Provide access to students, teachers, classes, courses and appointment data

#### API Resources
| Operation Name          | Description                                            | Participants                 | Resource(s) | Emitted Events    | Operation Details                                                 | Traits               |
|-------------------------|--------------------------------------------------------|------------------------------|-------------|-------------------|-------------------------------------------------------------------|----------------------|
| searchStudents()        | Search Students by student ID, first name and lastname | Teacher, Admin User          | Student     | Students.Searched | __Request Parameters:__ searchQuery,    __Returns:__   Student[ ] | safe   / synchronous |
| getTeachers()           | View available teachers                                | Teacher, Admin User          | Student     | Teacher.Viewed    | __Request Parameters:__,     __Returns:__   Teacher[ ]            | safe   / synchronous |
| searchTeachers()        | Search Teachers by empl ID, first name and lastname    | Teacher, Admin User          | Teacher     | Teachers.Searched | __Request Parameters:__ searchQuery,    __Returns:__   Teacher[ ] | safe   / synchronous |
| getClasses()            | View Classes by class number, name                     | Student                      | Class       | Classes.Searched  | __Request Parameters:__,     __Returns:__   Claas[ ]              | safe   / synchronous |
| getCourses()            | View available courses                                  | Teacher, Admin User, Student | Course      | Course.Viewed     | __Request Parameters:__ ,   __Returns:__   Course[ ]              | safe   / synchronous |
| getTeachersForClass()   | View Classes by teachers                               | Student                      | Teacher     | Teacher.Viewed    | __Request Parameters:__ classId,    __Returns:__   Teacher[ ]     | safe   / synchronous |
| getStudentsForTeacher() | View Students for teacher                              | Teacher, Admin User          | Student     | Students.Viewed   | __Request Parameters:__ teacherId,    __Returns:__   Student[ ]   | safe   / synchronous |

#### Modeled Resources   

##### Address 
| Property Name  | Description                    |
|----------------|--------------------------------|
| addressType    | If this is the primary address |
| streetAddress1 | Street 1                       |
| streetAddress2 | Street 2                       |
| city           | City                           |
| state          | State                          |
| zipCode        | ZipCode                        |
| country        | County                         |

##### Student
| Property Name | Description                                       |
|---------------|---------------------------------------------------|
| studentId     | Unique identifier identifying the student         |
| firstName     | Student first name                                |
| lastName      | Student last name                                 |
| address[]     | Student addresses (list of `Address`)                                |
| birthDate     | Student birthdate                                 |
| ageInYears    | Student's age in years (directly consumable data) |
| residency     | Student residency status                          |
 
##### Teacher          
| Property Name | Description                                          |
|---------------|------------------------------------------------------|
| emplId        | Unique employment identifier identifying the teacher |
| firstName     | Teacher first name                                   |
| lastName      | Teacher last name                                    |
| address[]     | Teacher addresses (list of `Address`)                |
| birthDate     | Teacher birthdate                                    |
| ageInYears    | Teacher's age in years (directly consumable data)    |

##### Class
| Property Name | Description                           |
|---------------|---------------------------------------|
| classId       | Unique identifier to identify a class |
| className     | Name of the class                     |
| credit        | Credits for this class                |
| location      | Address of the class location         |
| dayAndTime    | Day and the time of the class         |

##### Course
| Property Name | Description                           |
|---------------|---------------------------------------|
| courseId      | Unique identifier to identify a class |
| courseName    | Name of the class                     |
| termCode      | Term code for this course             |
| credit        | Total credits for this course         |
| teachers[]    | List of teachers (of  type `Teacher`) |


### Curricular API Design

| Resource Path                  | Operation Name          | HTTP Method | Description                                                                  | Request Details | Response Details | Response Code(s) |
|--------------------------------|-------------------------|-------------|------------------------------------------------------------------------------|-----------------|------------------|------------------|
| /students                      | getStudents()           | GET         | View students                                                                |                 | Students[]       | 200              |
| /students/search               | searchStudents()        | POST        | Search for students by student id, first/last name                           | searchQuery     | Students[]       | 200              |
| /teachers                      | getTeachers()           | GET         | View teachers                                                                |                 | Teacher[]        | 200              |
| /teachers/search               | searchTeachers()        | POST        | Search for teachers by empl id, first/last name                              | searchQuery     | Teacher[]        | 200              |
| /classes                       | getClasses()            | GET         | View classes(class is an instance of a course)                               |                 | Class[]          | 200              |
| /courses                       | getCourses()            | GET         | View courses(a course consist of multiple classes such as lab, lecture etc.) |                 | Class[]          | 200              |
| /classes/{classId}/teachers    | getTeachersForClass()   | GET         | View teachers for a class                                                    | classId         | Teacher[]        | 200              |
| /teachers/{teacherId}/students | getStudentsForTeacher() | GET         | View students for a teacher                                                  | teacherId       | Students[]       | 200              |

## Roadmap
- [x] Complete API design
- [x] Include tech stack
- [x] Complete system design. Include steps to build SQLite data for static data and build a Docker image with static data
- [x] Add the ability to turn locally. Include instructions to test the flow
- [x] Add Terraform for AWS deployment
- [x] Add OAuth2 API for local set up: https://github.com/go-oauth2/oauth2. Document how to use OAuth2.
- [ ] Implement pagination https://www.jsonapi.net/usage/reading/pagination.html
- [ ] Implement filtering(searching): document filter/search query langauge
  - Do some design around sorting and filtering through body vs query parameters 
  - https://help.smartsuite.com/en/articles/6963760-sorting-and-filtering-records-in-the-rest-api
  - https://www.jsonapi.net/usage/reading/filtering.html
  - Implement the parser for filtering
- [ ] Implement sorting https://www.jsonapi.net/usage/reading/sorting.html
- [ ] Implement sparse filed selection https://www.jsonapi.net/usage/reading/sparse-fieldset-selection.html
- [ ] Implement batch APIs
- [ ] Add OAuth2 via JWT: https://github.com/golang-jwt/jwt?tab=readme-ov-file
