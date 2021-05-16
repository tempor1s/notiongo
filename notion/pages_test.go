package notion

import "testing"

// TestPage_marshal will test marshalling functionality with mock Notion Page data
func TestPage_marshal(t *testing.T) {
	// test empty body
	testJSONMarshal(t, &Page{}, "{}")

	// mock
	// TODO(imthaghost): create a function that will generate random member values
	p := &Page{
		Object: "page",
		ID:     "0692b915742242bb988172b9705873d6",
	}

	// expected values
	want := `
		{
			"object": "page",
			"id": "0692b915742242bb988172b9705873d6"
		}
	`

	testJSONMarshal(t, p, want)

}

//func TestOrganization_marshal(t *testing.T) {
//	testJSONMarshal(t, &Organization{}, "{}")
//
//	o := &Organization{
//		BillingEmail:                         String("support@github.com"),
//		Blog:                                 String("https://github.com/blog"),
//		Company:                              String("GitHub"),
//		Email:                                String("support@github.com"),
//		TwitterUsername:                      String("github"),
//		Location:                             String("San Francisco"),
//		Name:                                 String("github"),
//		Description:                          String("GitHub, the company."),
//		IsVerified:                           Bool(true),
//		HasOrganizationProjects:              Bool(true),
//		HasRepositoryProjects:                Bool(true),
//		DefaultRepoPermission:                String("read"),
//		MembersCanCreateRepos:                Bool(true),
//		MembersCanCreateInternalRepos:        Bool(true),
//		MembersCanCreatePrivateRepos:         Bool(true),
//		MembersCanCreatePublicRepos:          Bool(false),
//		MembersAllowedRepositoryCreationType: String("all"),
//		MembersCanCreatePages:                Bool(true),
//		MembersCanCreatePublicPages:          Bool(false),
//		MembersCanCreatePrivatePages:         Bool(true),
//	}
//	want := `
//		{
//			"billing_email": "support@github.com",
//			"blog": "https://github.com/blog",
//			"company": "GitHub",
//			"email": "support@github.com",
//			"twitter_username": "github",
//			"location": "San Francisco",
//			"name": "github",
//			"description": "GitHub, the company.",
//			"is_verified": true,
//			"has_organization_projects": true,
//			"has_repository_projects": true,
//			"default_repository_permission": "read",
//			"members_can_create_repositories": true,
//			"members_can_create_public_repositories": false,
//			"members_can_create_private_repositories": true,
//			"members_can_create_internal_repositories": true,
//			"members_allowed_repository_creation_type": "all",
//			"members_can_create_pages": true,
//			"members_can_create_public_pages": false,
//			"members_can_create_private_pages": true
//		}
//	`
//	testJSONMarshal(t, o, want)
//}
