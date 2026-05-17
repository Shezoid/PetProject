package infrastructure

import (
	"PetProject/pkg"
)

type PetRepository struct {
	repository *pkg.Repository
}

func NewPetRepository(repository *pkg.Repository) *PetRepository {
	return &PetRepository{repository: repository}
}

func (petRepository *PetRepository) SavePet(pet *Pet) {
	row := petRepository.repository.Db.QueryRow(`select (pet.id) from pets`)
	var id int
	err := row.Scan(id)
	if err != nil {
		petRepository.repository.Db.Exec(`delete from pets where id = $1`, id)
		petRepository.repository.Db.Exec(`delete from pet_owners where pet_id = $1 or friend_id = $1`, id)
		petRepository.repository.Db.Exec(`delete from pet_friends where pet_id = $1`, id)
	}

	petRepository.repository.Db.QueryRow(`
			insert into pets 
		    (id, name, birth_date, breed, color, owner_id)
			values ($1, $2, $3, $4, $5))`,
		pet.Id, pet.Name, pet.BirthDate, pet.Breed, pet.Color)

	petRepository.repository.Db.QueryRow(`
			insert into pet_owners 
    		(pet_id, owner_id) 
			values ($1, $2)`,
		pet.Id, pet.OwnerId)

	for _, id := range pet.FriendIds {
		petRepository.repository.Db.QueryRow(`
			insert into pet_friends 
    		(pet_id, friend_id) 
			values ($1, $2)`,
			pet.Id, id)
	}
}

func (petRepository *PetRepository) FindPetById(petId int) (pet *Pet) {
	row := petRepository.repository.Db.QueryRow(`
			select (id, name, birth_date, breed, color, owner_id) 
			from pets
			where id = $1`,
		petId)
	pet = &Pet{}
	err := row.Scan(&pet.Id, &pet.Name, &pet.BirthDate, &pet.Breed, &pet.Color, &pet.OwnerId)
	if err != nil {
		println(err.Error())
	}

	return pet
}

func (petRepository *PetRepository) FindAllPet() (pets []Pet) {
	rows, err := petRepository.repository.Db.Query(`
			select (id, name, birth_date, breed, color, owner_id) 
			from pets`)
	if err != nil {
		println(err.Error())
		return
	}

	pets = []Pet{}
	for rows.Next() {
		pet := &Pet{}
		err := rows.Scan(&pet.Id, &pet.Name, &pet.BirthDate, &pet.Breed, &pet.Color, &pet.OwnerId)
		if err != nil {
			println(err.Error())
		}
		pets = append(pets, *pet)
	}
	return pets
}

func (petRepository *PetRepository) Delete(id int) {
	petRepository.repository.Db.Exec(`delete from pets where id = $1`, id)
}
