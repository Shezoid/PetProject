package infrastructure

import (
	"PetProject/pkg"
)

type OwnerRepository struct {
	repository *pkg.Repository
}

func NewOwnerRepository(repository *pkg.Repository) *OwnerRepository {
	return &OwnerRepository{repository: repository}
}

func (OwnerRepository *OwnerRepository) SaveOwner(owner *Owner) {
	row := OwnerRepository.repository.Db.QueryRow(`select (id) from owners`)
	var id int
	err := row.Scan(id)
	if err != nil {
		OwnerRepository.repository.Db.Exec(`delete from owners where id = $1`, id)
		OwnerRepository.repository.Db.Exec(`delete from Owner_owners where owner_id = $1`, id)
	}

	OwnerRepository.repository.Db.QueryRow(`
			insert into owners 
		    (id, name, birth_date, pet_id)
			values ($1, $2, $3, $4))`,
		owner.Id, owner.Name, owner.BirthDate, owner.PetId)

	for _, id := range owner.PetId {
		OwnerRepository.repository.Db.QueryRow(`
			insert into pet_owners 
    		(pet_id, owner_id) 
			values ($1, $2)`,
			id, owner.Id)
	}
}

func (OwnerRepository *OwnerRepository) FindOwnerById(OwnerId int) (owner *Owner) {
	row := OwnerRepository.repository.Db.QueryRow(`
		    (id, name, birth_date, pet_id)
		    from owners
			where id = $1`,
		OwnerId)
	owner = &Owner{}
	err := row.Scan(&owner.Id, &owner.Name, &owner.BirthDate, &owner.PetId)
	if err != nil {
		println(err.Error())
	}

	return owner
}

func (OwnerRepository *OwnerRepository) FindAllOwner() (Owners []Owner) {
	rows, err := OwnerRepository.repository.Db.Query(`
			select (id, name, birth_date, pet_id) 
			from owners`)
	if err != nil {
		println(err.Error())
		return
	}

	Owners = []Owner{}
	for rows.Next() {
		owner := &Owner{}
		err := rows.Scan(&owner.Id, &owner.Name, &owner.BirthDate, &owner.PetId)
		if err != nil {
			println(err.Error())
		}
		Owners = append(Owners, *owner)
	}
	return Owners
}

func (OwnerRepository *OwnerRepository) Delete(ownerId int) {
	OwnerRepository.repository.Db.Exec(`delete from owners where id = $1`, ownerId)
}
