# WBA BE Project - 23

### 기술 스택

go version : go1.19.4 darwin/arm64

# 실행 방법

```bash
git clone https://github.com/codestates/WBABEProject-23.git
cd WBABEProject-23
go mod tidy
```

설정파일은 빠져있으므로 config 폴더에 config.toml 파일을 만들고 다음 내용을 넣는다.

```bash
[server] ##nomal type
mode = "dev"
host = "localhost:8080"

[db] #data access object
[db.admin]
host = "mongodb://127.0.0.1:27017"
user = #user name
pass = #password

[log]
level = "debug"
fpath = "./logs/go-loger"
msize = 2000
mage = 7
mbackup = 5
```

이후 실행한다

```bash
go run main.go
```

# **기능**

### 메뉴 신규 등록  - 피주문자

- **API | POST /menu/admin/new**
    - 사업장에서 신규 메뉴 관련 정보를 등록하는 과정(ex. 메뉴 이름, 주문가능여부, 한정수량,  원산지, 가격, 맵기정도, etc)
    - 성공 여부를 리턴
    - input form
        
        ```json
        {
          "category": "string",
          "name": "string",
          "origin": "string",
          "price": 0
        }
        ```
        

### 메뉴 수정 / 삭제 - 피주문자

- **API | PATCH /menu/admin/modify**
    - 사업장에서 기존의 메뉴 정보 변경기능(ex. 가격변경, 원산지 변경, soldout)
    - 메뉴 삭제시, 실제 데이터 백업이나 뷰플래그를 이용한 안보임 처리
    - 금일 추천 메뉴 설정 변경, 리스트 출력
    - 성공 여부를 리턴
    - input form
        
        ```json
        {
          "category": "string",
          "origin": "string",
          "price": 0,
          "state": 0,
          "toUpdate": "string"
        }
        ```
        

### 메뉴 리스트 출력 조회 - 주문자

- **API | GET /menu/list**
    - 각 카테고리별  sort 리스트 출력(ex. order by 추천, 평점, 재주문수, 최신)
    - 결과 5~10여개 임의 생성 출력, sorting 여부 확인
    - input form
        
        ```
        name :  query : 이름
        sort :  query : sort할 컬럼 이름
        order:  query : 1은 오름차순 그 외 내림자순
        ```
        

### 메뉴별 평점 및 리뷰 조회 - 주문자

- **API | GET /menu/list/review**
    - UI에서 메뉴 리스트에서 상기 리스트 출력에 따라 개별 메뉴를 선택했다고 가정
    - 해당 메뉴 선택시 메뉴에 따른 평점 및 리뷰 데이터 리턴
    - input form
        
        ```
        id   : query : 가게 사업체 id
        name : query : 메뉴 이름
        ```
        

### 메뉴별 평점 작성 - 주문자

- **API | POST /review**
    - 해당 주문내역을 기준, 평점 정보, 리뷰 스트링을 입력받아 과거 주문내역 업데이트 저장
    - 성공 여부 리턴
    - input form
        
        ```json
        {
          "businessID": "string",
          "content": "string",
          "menuName": "string",
          "orderID": "string",
          "orderer": "string",
          "score": 0
        }
        ```
        

### 주문 - 주문자

- **API | POST /order/make**
    - 주문정보를 입력받아 주문 저장(ex. 선택 메뉴 정보, 전화번호, 주소등 정보를 입력받아 DB 저장)
    - 주문 내역 초기상태 저장
    - 금일 주문 받은 일련번호-주문번호 리턴
    - input form : 주문자 이름, 주문 가게 이름, 메뉴 배열형태만 입력
        
        ```json
        {
          "businessName": "string",
          "menu": [
            {
              "menuName": "string",
              "number": 0
            }
          ],
          "orderer": "string"
        }
        ```
        

### 주문 변경 - 주문자

- **API | PATCH /order/modify**
    - 메뉴 추가시 상태조회 후 `배달중`일 경우 실패 알림
        - 성공 실패 알림, ~~실패시 신규주문으로 전환~~
    - 메뉴 변경시 상태가 `조리중`, `배달중`일 경우 확인
        - 성공 실패 알림
    - input form : 수정할 주문 번호, 변경한 주문 메뉴 [{메뉴이름, 수량}]
        
        ```json
        {
          "menu": [
            {
              "menuName": "string",
              "number": 0
            }
          ],
          "orderID": "string"
        }
        ```
        

### 주문 내역 조회 - 주문자

- **API | GET /order/list**
    - 현재 주문내역 리스트 및 상태 조회 - 하기 **주문 상태 조회**에서도 사용
        - ex. 접수중/조리중/배달중 etc
        - 없으면 null 리턴
    - 과거 주문내역 리스트 최신순으로 출력
        - 없으면 null 리턴
    - 과거 주문 : 배달 완료된 주문, 현재 주문: 그 외
    - input form
        
        ```
        name : query : 유저이름
        cur : query : 1은 현재, 그외 과거
        ```
        

### 주문 상태 조회 - 피주문자

- **API | GET /order/admin/list**
    - input form
        
        ```
        businessname : query : 사업체이름
        ```
        
- **API | PATCH /order/admin/update**
    - ~~메뉴별로 상태 저장~~  → 주문 별로 저장
    - ex. 상태 : 접수중/접수취소/추가접수/접수-조리중/배달중/배달완료 등을 이용 상태 저장
    - 각 단계별 사업장에서 상태 업데이트
        - **접수중 → 접수** or **접수취소 → 조리중** or **추가주문 → 배달중**
        - 성공여부 리턴
    - input form
        
        ```json
        {
          "orderId": "string",
          "state": 0
        }
        ```
        
    
- Swagger 참조
    
    ![swagger화면](README/Untitled.png)
    

# DB

![DB컬렉션](README/Untitled%201.png)

위와 같이 4개의 컬렉션을 만들었다.

### business

사업체 정보를 저장하는 컬렉션. 이번 구현한 api에서 write하지 않기 때문에 직접 데이터를 만들어 넣어준다.

name, admin, menu 필드를 갖고있고 각각 사업체 이름, 관리자 유저 id, 배열 형태의 메뉴를 나타낸다. 설계 당시에 메뉴를 컬렉션으로 만들지 않아 배열로 저장하게 되었다. 

ex)

![business](README/Untitled%202.png)

메뉴는 name, state, score, is_deleted를 필수적으로 갖고 각각 메뉴이름, 준비 상태, 평점, 삭제 여부 뷰 플레그를 나타낸다. 그 외에 가격, 원산지, 종류 등이 추가로 있을 수 있다.

실제로 사용한 필드는 name, price, origin, score, is_deleted, category 다

### order

주문 정보를 저장하는 컬렉션 api를 통해 추가하고 상태를 변경할 수 있다.

필드로 orderid, orderer, businessname, menu, createdat, state를 갖고 각각 

주문번호, 주문자 이름, 사업체 이름, {주문메뉴, 수량, 리뷰}의 배열, 주문 시간, 상태를 나타낸다.

주문번호는, 그 날에 몇 번째로 주문했는지를 나타낸다. 다만 mongodb에서 Date가 UTC로만 저장 돼서 UTC기준으로 카운팅 되게 된다.

ex)

![order](README/Untitled%203.png)

메뉴는 자세히 보면 다음과 같이 생겼다.

![order.menu](README/Untitled%204.png)

### review

주문 메뉴에대한 리뷰를 저장하는 컬렉션

필드로 orderid, businessid, orderer, menuname, content, score가 있고 각각 작성한 리뷰가 들어간 order의 id, 사업체의 id, 주문자, 메뉴이름, 리뷰 내용, 평점이다.

![review](README/Untitled%205.png)

설계 단계에서 NoSQL인데 rdb처럼 설계하면 안되는게 아닌가 싶어서 많이 부족한 구성이다. 메뉴를 컬렉션으로 분리하고 order 컬렉션에서 business 컬렉션을 참조하는 컬럼도 통일해야할 듯 하다.