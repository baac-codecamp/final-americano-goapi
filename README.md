# Check Salak GO API

Check Salak GO API คือ REST API สำหรับเว็บไซต์ตรวจสลากออมทรัพย์ ธ.ก.ส. โดยใช้ภาษา Golang

## 📦 Built With

- [x] Golang
- [x] Gorilla/mux
- [x] MongoDB Go Driver

## 📋 Feature

## 🏷 Version

## 📘 API Document

**URL SIT** : `http://localhost:8000`

**URL UAT** : `https://americano-salak-goapi.topwork.asia`

1. **[getNews](#)** : `ดึงข้อมูลข่าวสารทั้งหมด โดยเรียงจากวันที่ล่าสุด`
   <details>
   <summary>View Details</summary>

   - **Method** : `GET`
   - **Url** : `http://{URL}:{PORT}/go/getNews`
   - **Request** :
     - header : `-`
     - body : `-`
     - param : `-`
     - queryString : `-`
   - **Response** :

     ```json
     {
       "response_status": "success",
       "response_message": "Get News Success!",
       "response_data": {
         "ListNews": [
           {
             "_id": "5f339227cb71c124ba2f417e",
             "title": "ฝาก-โอนเงินที่ตู้บุญเติม ลุ้นรับทองคำหนัก 2 สลึง",
             "imgUrl": "https://www.baac.or.th/file-upload/015299-1-OK%20AD_A-Mobile%20-01.jpg",
             "desc": "ฝาก-โอนเงินที่ตู้บุญเติม ลุ้นรับทองคำหนัก 2 สลึง",
             "createdAt": "2020-08-12T06:54:31.542Z",
             "updatedAt": "2020-08-12T06:54:31.542Z",
             "__v": 0,
             "id": "5f339227cb71c124ba2f417e"
           },
           {
             "_id": "5f339227cb71c124ba2f417f",
             "title": "งานมหกรรมการเงินโคราช ครั้งที่ 14 Money Expo Korat 2020",
             "imgUrl": "https://www.baac.or.th/file-upload/015295-1-leaflet%20on%20web-1.jpg",
             "desc": "งานมหกรรมการเงินโคราช ครั้งที่ 14 Money Expo Korat 2020",
             "createdAt": "2020-08-12T06:54:31.542Z",
             "updatedAt": "2020-08-12T06:54:31.542Z",
             "__v": 0,
             "id": "5f339227cb71c124ba2f417f"
           }
         ]
       }
     }
     ```

     </details>

2. **[getNewsById](#)** : `ดึงข้อมูลข่าวสาร จาก id ข้อมูล`
   <details>
   <summary>View Details</summary>

   - **Method** : `GET`
   - **Url** : `http://{URL}:{PORT}/go/getNewsById/:id`
   - **Request** :
     - header : `-`
     - body : `-`
     - param : `id`
     - queryString : `-`
   - **Response** :
     ```json
     {
       "response_status": "success",
       "response_message": "Get news by id success!",
       "response_data": {
         "currNews": {
           "_id": "5f339227cb71c124ba2f417e",
           "title": "ฝาก-โอนเงินที่ตู้บุญเติม ลุ้นรับทองคำหนัก 2 สลึง",
           "imgUrl": "https://www.baac.or.th/file-upload/015299-1-OK%20AD_A-Mobile%20-01.jpg",
           "desc": "ฝาก-โอนเงินที่ตู้บุญเติม ลุ้นรับทองคำหนัก 2 สลึง",
           "createdAt": "2020-08-12T06:54:31.542Z",
           "updatedAt": "2020-08-12T06:54:31.542Z",
           "__v": 0,
           "id": "5f339227cb71c124ba2f417e"
         }
       }
     }
     ```

    </details>

## ⚙ How to use

1. Clone project

   ```bash
   git clone https://github.com/baac-codecamp/final-americano-goapi.git
   ```

2. Go to project folder

   ```bash
   cd final-americano-goapi
   ```

3. Run project by command

   ```bash
   go run main.go
   ```

4. Runs the app in the development mode.
   Open [http://localhost:8000](http://localhost:8000) to view it in the browser.
