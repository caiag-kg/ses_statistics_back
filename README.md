# DataReadAPI  ![Golang](./statics/golang.svg)

## Project Overview

This project aims to provide a backend solution for managing and analyzing seismic event data. It allows users to retrieve, filter, and sort earthquake event information stored in a PostgreSQL database. The application is built using Go and utilizes the Gin framework for handling HTTP requests. Its primary goal is to enable users to easily access and analyze seismic data for research and educational purposes.

### Features

* Retrieve a list of seismic events from the database.
* Filter events based on various criteria, such as epicenter and magnitude.
* Sort events by date or other attributes.

### Getting Started

1. Clone the repository: `git clone https://github.com/caiag-kg/ses_statistics_back.git`
2. Copy the **xconfig.yaml** file to **config.yaml** and fill in the necessary configuration settings: 
    - Linux, MacOS: `cp xconfig.yaml config.yaml`
    - Windows: `copy xconfig.yaml config.yaml`
3. Build the Docker image: 
    - Windows, MacOS: `docker build -t ses_statistics_back .`
    - Linux: `make build`
4. Run the Docker container: 
    - Windows or MacOS: `docker run -d --name ses_back -p 8000:8000 ses_back`
    - Linux: `make run`
5. To test the API, you can send a `GET` request to your IP address and port to retrieve all data from the database. Alternatively, you can send a `POST` request with a JSON body structured like the example in the **POST_request_model.json** file to filter data based on specific criteria.

### Contributing

Contributions are welcome! If you'd like to contribute to this project, please follow these steps:

1. Fork the repository
2. Create a new branch: `git checkout -b [branch-name]`
3. Make your changes and commit them: `git commit -m '[commit-message]'`
4. Push your changes: `git push origin [branch-name]`
5. Open a pull request

### License

This project is licensed under the [MIT License](https://tlo.mit.edu/understand-ip/exploring-mit-open-source-license-comprehensive-guide).

### Author

* **Kairat Kubanychbek uulu**:
    - GitHub: [https://github.com/ImKairat](https://github.com/ImKairat)
    - Mail: k.kubanychbek@caiag.kg

### Organization: 

* **Central-Asian Institute for Applied Geosciences (CAIAG)**:
    - Address: *73/2 Timur Frunze Street, Bishkek, 720027, Kyrgyz Republic.*
    - Website: [https://caiag.kg/](https://caiag.kg/)
    - GitHub: [https://github.com/caiag-kg](https://github.com/caiag-kg)