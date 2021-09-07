package option

import "dam/config"

type DefaultRepo struct{}

func (o *DefaultRepo) GetUnknownRepoName() string {
	return config.UNKNOWN_REPO_NAME
}

func (o *DefaultRepo) GetUnknownRepoID() int {
	return config.UNKNOWN_REPO_ID
}

func (o *DefaultRepo) GetNewRepoPrefix() string {
	return config.NEW_REPO_PREFIX
}

func (o *DefaultRepo) GetNewRepoPostfixLimit() int {
	return config.NEW_REPO_POSTFIX_LIMIT
}

func (o *DefaultRepo) GetLabelReposSeparator() string {
	return config.LABEL_REPOS_SEPARATOR
}
