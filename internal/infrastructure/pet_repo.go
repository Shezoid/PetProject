package infrastructure

import (
	"PetProject/pkg"
)

type PetRepository struct {
	repository *pkg.Repository
}

func NewPet(repository *pkg.Repository) *PetRepository {
	return &PetRepository{repository: repository}
}

func (r *PetRepository) Create(pet *Pet) error {
	err := r.repository.Db.QueryRow(`
		INSERT INTO pets
		(name, birth_date, breed, color, owner_id)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`,
		pet.Name,
		pet.BirthDate,
		pet.Breed,
		pet.Color,
		pet.OwnerID,
	).Scan(&pet.ID)

	if err != nil {
		return err
	}

	for _, friendID := range pet.FriendIDs {
		_, err := r.repository.Db.Exec(`
			INSERT INTO pet_friends
			(pet_id, friend_id)
			VALUES ($1, $2)`,
			pet.ID,
			friendID,
		)

		if err != nil {
			return err
		}
	}

	return nil
}

func (r *PetRepository) Update(pet *Pet) error {
	_, err := r.repository.Db.Exec(`
		UPDATE pets
		SET
			name = $1,
			birth_date = $2,
			breed = $3,
			color = $4,
			owner_id = $5
		WHERE id = $6`,
		pet.Name,
		pet.BirthDate,
		pet.Breed,
		pet.Color,
		pet.OwnerID,
		pet.ID,
	)

	if err != nil {
		return err
	}

	_, err = r.repository.Db.Exec(`
		DELETE FROM pet_friends
		WHERE pet_id = $1`,
		pet.ID,
	)

	if err != nil {
		return err
	}

	for _, friendID := range pet.FriendIDs {
		_, err = r.repository.Db.Exec(`
			INSERT INTO pet_friends
			(pet_id, friend_id)
			VALUES ($1, $2)`,
			pet.ID,
			friendID,
		)

		if err != nil {
			return err
		}
	}

	return nil
}

func (r *PetRepository) FindByID(petID int) (pet *Pet, err error) {
	row := r.repository.Db.QueryRow(`
			select (id, name, birth_date, breed, color, owner_id) 
			from pets
			where id = $1`,
		petID)
	pet = &Pet{}
	err = row.Scan(&pet.ID, &pet.Name, &pet.BirthDate, &pet.Breed, &pet.Color, &pet.OwnerID)

	return pet, err
}

func (r *PetRepository) FindAllPet() (pets []Pet, err error) {
	rows, err := r.repository.Db.Query(`
			select (id, name, birth_date, breed, color, owner_id) 
			from pets`)
	if err != nil {
		println(err.Error())
		return
	}

	pets = []Pet{}
	for rows.Next() {
		pet := &Pet{}
		err := rows.Scan(&pet.ID, &pet.Name, &pet.BirthDate, &pet.Breed, &pet.Color, &pet.OwnerID)
		if err != nil {
			println(err.Error())
		}
		pets = append(pets, *pet)
	}
	return pets, rows.Err()
}

func (r *PetRepository) Delete(id int) (err error) {
	_, err = r.repository.Db.Exec(`delete from pets where id = $1`, id)
	return err
}
