### MSaaS Backend 를 소개합니다.

MSaaS Backend 는 MSaaS 서비스의 Backend 서버를 구성하는 모듈입니다.

MSaaS Backend 는 다음과 같은 컴포넌트로 구성되어 있습니다.

1. WIZCRAFT (**W**orking **I**n the microservice **Z**one: **CRAFT**ing Scalable Architectures) 
   - WIZCRAFT는 MSaaS 서비스의 API 서버를 구성하는 모듈입니다. 
   - Kubernetes 프로젝트의 apiserever 의 개념을 차용했습니다.
   - MSA의 구성요소를 Project, Service, Database, APIspec, ... 으로 정의하고, 이를 REST API로 관리합니다.
   - Team 단위로 테넌시를 보장하며, Team 별로 Project를 관리할 수 있습니다.
   - 생성한 Project 와 GitHub 연동 기능을 제공합니다.
2. WAND (CLI tool for MSaaS : **W**ith **A** **N**imble **D**evice)
   - WAND 는 MSaaS 서비스의 Backend 서버를 관리하는 CLI 도구입니다.
   - Kubernetes 프로젝트의 kubectl 의 개념을 차용했습니다.
   - 별도의 프론트엔드 서버 없이, Backend 서버를 관리할 수 있습니다.
3. SPELLBOOK (**M**igration tool for MSaaS : **S**witching **P**latform **E**asily from **L**ocal to **L**andscape)
   - SPELLBOOK은 MSaaS 서비스의 배포와 관리를 위한 마이그레이션 도구입니다. 
   - docker-compose 으로의 배포와 Kubernetes 으로의 배포, 그리고 두 환경간의 마이그레이션을 지원합니다.
   - yaml 파일 작성에 대한 스키마 툴을 제공합니다.
   - 프로덕션, 테스트, 개발 환경을 구분하여 관리할 수 있습니다.