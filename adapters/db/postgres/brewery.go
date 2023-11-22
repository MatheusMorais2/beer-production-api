package postgres

import (
	"beer-production-api/entities/brewery"
	"beer-production-api/entities/recipe"
	"database/sql"
)

type BreweryRepository struct {
	db *sql.DB
}

func NewBreweryRepository(db *sql.DB) *BreweryRepository {
	return &BreweryRepository{db: db}
}

func (t *BreweryRepository) Insert(breweryToInsert *brewery.CreateBreweryDBInputDTO) (*brewery.Brewery, error) {
	createdBrewery := &brewery.Brewery{}

    err := t.db.QueryRow(
        `INSERT INTO "brewery" (name, email)
        VALUES ($1, $2)
        RETURNING id
        `,
		breweryToInsert.Name, breweryToInsert.Email,
	).Scan(&createdBrewery.ID)

    if err != nil {
		return nil, err
	}

    _, err = t.db.Exec(`
    INSERT INTO brewery_user (user_id, brewery_id, role)
    VALUES ($1, $2, 'creator')
    RETURNING id
    `, breweryToInsert.UserId, createdBrewery.ID)

	if err != nil {
		return nil, err
	}

	return createdBrewery, nil
}

func (t *BreweryRepository) GetByID(id string) (*brewery.Brewery, error) {
    brewery := &brewery.Brewery{} // Create a brewery object to store the retrieved data

    err := t.db.QueryRow(
        `SELECT id, name, email FROM "brewery" WHERE id = $1`,
        id,
    ).Scan(&brewery.ID, &brewery.Name, &brewery.Email)

    if err != nil {
        return nil, err
    }

    return brewery, nil
}

func (t *BreweryRepository) GetBreweriesByUserId(userID string) ([]*brewery.GetUserBreweriesOutputDTO, error) {
    breweries := []*brewery.GetUserBreweriesOutputDTO{}

    rows, err := t.db.Query(
        `SELECT
        b.id AS brewery_id, b.name AS brewery_name, b.email AS brewery_email, bu.role as user_role
        FROM "brewery" b
        JOIN "brewery_user" bu 
        ON b.id = bu.brewery_id
        WHERE bu.user_id = $1;`,
        userID,
    )
    if err != nil {
        return nil, err
    }
    defer rows.Close()


    for rows.Next() {
        breweryForUser := &brewery.GetUserBreweriesOutputDTO{}
        err := rows.Scan(&breweryForUser.Id, &breweryForUser.Name, &breweryForUser.Email, &breweryForUser.Role)
        if err != nil {
            return nil, err
        }
        breweries = append(breweries, breweryForUser)
    }
    return breweries, nil
}

func (t *BreweryRepository) InviteUserToBrewery(invite brewery.InviteUserInputDTO) (*brewery.Invite, error) {
    createdInvite := &brewery.Invite{}

    err := t.db.QueryRow(`
        INSERT INTO "brewery_invite" (inviting_user_id, invited_user_id, brewery_id, role)
        VALUES ($1, $2, $3, $4)
        RETURNING id, inviting_user_id, invited_user_id, brewery_id, role;`,
    invite.InvitingUserId, invite.InvitedUserId, invite.BreweryId, invite.Role).
    Scan(&createdInvite.Id, &createdInvite.InvitingUserId, &createdInvite.InvitingUserId, &createdInvite.BreweryId, &createdInvite.Role)
    
    if err != nil {
        return nil, err
    }

    return createdInvite, nil
}

func (t *BreweryRepository) GetUserRole(userId string, breweryId string) (string, error) {
    var role string

    err := t.db.QueryRow(`
        SELECT role FROM brewery_user WHERE user_id = $1 AND brewery_id = $2;
    `, userId, breweryId).Scan(&role)
     
    if err != nil {
        return "", err
    }

    return role, nil
}

//TODO: essa funcao deve ser em recipes como Get Recipes by Brewery ID
func (t *BreweryRepository) GetBreweryRecipes(breweryID string) ([]*recipe.Recipe, error) {
    recipes := []*recipe.Recipe{}

    rows, err := t.db.Query(
        `SELECT id, name FROM "recipe" WHERE brewery_id = $1`,
        breweryID,
    )
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        recipeForBrewery := &recipe.Recipe{}
        err := rows.Scan(&recipeForBrewery.ID, &recipeForBrewery.Name)
        if err != nil {
            return nil, err
        }
        recipes = append(recipes, recipeForBrewery)
    }

    return recipes, nil
}

//TODO: essa funcao deve ficar em batches como Get batches by brewery id
/* func (t *BreweryRepository) GetBreweryBatches(breweryID uuid.UUID) ([]*batch.BatchWithRecipe, error) {
    batches := []*batch.BatchWithRecipe{}

    rows, err := t.db.Query(
        `SELECT b.id AS batch_id, b.name AS batch_name, b.start_date, b.finish_date, r.id AS recipe_id, r.name AS recipe_name
         FROM "batch" b
         JOIN "recipe" r ON b.recipe_id = r.id
         WHERE r.brewery_id = $1`,
        breweryID,
    )
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        batchWithRecipe := &batch.BatchWithRecipe{}
        err := rows.Scan(&batchWithRecipe.BatchID, &batchWithRecipe.BatchName, &batchWithRecipe.StartDate, &batchWithRecipe.FinishDate, &batchWithRecipe.RecipeID, &batchWithRecipe.RecipeName)
        if err != nil {
            return nil, err
        }
        batches = append(batches, batchWithRecipe)
    }

    return batches, nil
} */