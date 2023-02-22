package messages

const (
	AlreadyExist               = "un mot de passe existe déjà avec le libellé %s"
	PasswordNotFound           = "aucun mot de passe trouvé avec le libellé %s"
	AddLabelAndValueMissing    = "Pour ajouter un nouveau mot de passe,\nrenseignez le libellé et la valeur de ce dernier"
	AddExample                 = "-add -label=\"foo\" -value=\"bar\""
	UpdateLabelAndValueMissing = "Pour mettre à jour un mot de passe,\nrenseignez son libellé et sa nouvelle valeur valeur"
	UpdateExample              = "-update -label=\"foo\" -value=\"bar\""
	ShowLabelMissing           = "Pour visualiser un mot de passe,\nrenseignez son libellé"
	ShowExample                = "-show -label=\"foo\""
	DeleteLabelMissing         = "Pour supprimer un mot de passe,\nrenseignez son libellé"
	DeleteExample              = "-delete -label=\"foo\""
	AddSuccess                 = "Mot de passe %s ajouté avec succès"
	UpdateSuccess              = "Mot de passe %s mis à jour avec succès"
	DeleteSuccess              = "Mot de passe %s supprimé avec succès"
)
