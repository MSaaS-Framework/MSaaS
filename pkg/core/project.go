package core

/*

project.go 파일은 프로젝트와 관련된 기능을 제공합니다. 단순 데이터베이스의 CRUD 기능을 제공하는 것이 아닌, 물리적 로직 레벨에서 프로젝트를 생성하고 관리하는 기능을 제공합니다.

1. 로컬 호스트 경로에 모든 마이크로서비스를 관장하는 비즈니스 레벨 개념인 프로젝트를 생성합니다.
호스트 경로는 사용자가 지정할 수 있으며, 프로젝트의 실제 코어 파일은 유지 관리는 프로젝트 경로에 저장된 파일을 통해 이루어지지만 메타데이터는 디비에 저장 됩니다.
사용자가 github 계정을 연동하였다면, 프로젝트 생성시 git organization 저장소를 생성합니다.

2. 사용자가 호스트 경로 설정을 바꾸면 프로젝트 경로를 변경합니다.

3. 프로젝트를 삭제하면 프로젝트 경로에 저장된 파일을 삭제하고, 디비에 저장된 메타데이터를 삭제합니다.

4. 이외의 필요한 로직들을 처리합니다 ... 향후 추가

*/

import (
	"MSaaS-Framework/MSaaS/pkg/object"
	"context"
	"fmt"
	"os"

	// 권한 확인용 패키지
	// 여기에 필요한 패키지 추가

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// 프로젝트를 호스트에 생성하는 함수
func CreateProjectInHostPath(project *object.Project) error {

	// 호스트에 생성할 프로젝트 경로 (폴더 + 프로젝트 이름)
	path := project.HostPath + "/" + project.General.Name

	// project.HostPath 경로 권한 확인 및 폴더 생성. 만약 쓰기 권한이 없으면 에러 반환
	if err := ensureDirectoryExists(path); err != nil {
		return err
	}

	// project.GitHubRepo가 존재하면, GitHub API를 통해 저장소 생성 (추후 구현)
	// if project.GitHubRepo != "" {
	//     // GitHub API 호출 코드
	// }

	return nil
}

// 경로가 존재하는지 확인하고, 존재하지 않으면 생성하는 함수
func ensureDirectoryExists(path string) error {
	// 경로가 존재하는지 확인
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// 경로가 없으면 폴더 생성
		if err := os.MkdirAll(path, 0755); err != nil {
			return fmt.Errorf("failed to create directory at %s: %v", path, err)
		}
	} else if err != nil {
		// 다른 에러 발생 시 에러 반환
		return fmt.Errorf("failed to stat path %s: %v", path, err)
	}

	return nil
}

// GitHub 저장소를 생성하는 함수
func CreateGitHubRepo(project *object.Project) error {

	// GitHub Personal Access Token을 통해 인증
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: project.User.GitHub.AccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// 새로운 저장소 생성 요청 설정
	repo := &github.Repository{
		Name:    github.String(project.GitHubRepo),
		Private: github.Bool(true), // true일 경우 비공개 저장소 생성, false는 공개
	}

	// GitHub API를 통해 조직에 저장소 생성
	createdRepo, _, err := client.Repositories.Create(ctx, project.General.Name, repo)
	if err != nil {
		return fmt.Errorf("failed to create GitHub repository in organization '%s': %v", project.General.Name, err)
	}

	// 저장소가 성공적으로 생성되면 URL 저장
	project.GitHubRepo = createdRepo.GetHTMLURL()

	fmt.Printf("GitHub repository '%s' created successfully! URL: %s\n", project.GitHubRepo, project.GitHubRepo)
	return nil
}

func UpdateProjectInHostPath(project *object.Project) error {

	// 호스트에 프로젝트 경로 변경

	return nil

}

func DeleteProjectInHostPath(project *object.Project) error {

	// 호스트에 프로젝트 경로 삭제

	return nil

}
