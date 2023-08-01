# teler Proxy (Demo)

Welcome to the teler Proxy demo! This directory showcases how to utilize teler Proxy to secure the Damn Vulnerable Web Application (DVWA).

## Configuration

The `config/` directory houses essential configurations for both DVWA and teler-proxy. Below are the changes made to each component:

### DVWA

* Authentication: **Disabled**
* Default Security Level: Set to **Low**

These adjustments ensure that the DVWA is more vulnerable and allows for better testing and demonstration of the capabilities of teler-proxy.

### teler-waf

* Whitelisted URIs: We have added a whitelist pattern that matches URIs with the regex `"^/(index|about)\.php"`. This action protects against the **DirectoryBruteforce** threat, ensuring these specific URIs are secure.

## Run

**Prerequisites:**

Before starting, ensure you have the following prerequisites installed on your system:

* [docker](https://docs.docker.com/engine/install/)
* [docker-compose](https://docs.docker.com/compose/install/)

To run the teler Proxy demo with the pre-configured settings, simply execute `docker-compose up` command in this directory.

The demo will now start running, utilizing the teler Proxy along with the specified configurations for the Damn Vulnerable Web Application (DVWA) and the teler-waf. Access the DVWA application using your preferred web browser in http://localhost:8080.

Now you are all set! With the teler Proxy in place, you can explore the various functionalities and features it offers for protecting DVWA from potential threats. Test the security of DVWA by attempting various attacks, and observe how teler-waf handles and mitigates them.

Happy testing and stay secure!

---

> ⚠️ **Warning**: Please note that this is a simplified demonstration of the teler Proxy's capabilities and is intended for testing purposes only. In a real-world scenario, you would integrate the teler Proxy with your web application to enhance security and protect it from potential attacks. For detailed instructions on integrating the teler Proxy with your application, refer to the full documentation available in the main repository.