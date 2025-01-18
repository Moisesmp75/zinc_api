const main_url = "http://localhost:3000/"
const url = `${main_url}api/v1/email`

export class MailService {
  constructor() {
    this.requestOption = {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      }
    }
  }
  async search_data(size, from, sort = "") {
    try {
      const sizeParam = size !== undefined && size !== null ? String(size) : "10";
      const fromParam = from !== undefined && from !== null ? String(from) : "0";
      const sortParam = sort ? String(sort) : "";

      const queryParams = new URLSearchParams();
      queryParams.append("size", sizeParam);
      queryParams.append("from", fromParam);
      if (sortParam) {
        queryParams.append("sort", sortParam);
      }
      const requestUrl = `${url}/search?${queryParams.toString()}`;

      const response = await fetch(requestUrl, this.requestOption);
      if(!response.ok)
        throw new Error(response.status)
      
      return await response.json()
    } catch (error) {
      console.log(error)
      throw error
    }
  }
  async new_mail(request) {
    try {
      this.requestOption["body"] = JSON.stringify(request)
      const response = await fetch(`${url}`, this.requestOption)
      if(!response.ok)
        throw new Error(response.status)
      const { success, message, data, meta } = await response.json()
      return {
        success,
        message,
        data,
        meta
      }
    } catch (error) {
      console.log(error)
      throw error
    }
  }
}