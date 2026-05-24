package infrastructure

import (
	"PetProject/pkg"
)

type OwnerRepository struct {
	repository *pkg.Repository
}

func NewOwner(repository *pkg.Repository) *OwnerRepository {
	return &OwnerRepository{repository: repository}
}

func (r *OwnerRepository) Create(owner *Owner) error {
	err := r.repository.Db.QueryRow(`
		INSERT INTO owners
		(name, birth_date)
		VALUES ($1, $2)
		RETURNING id`,
		owner.Name,
		owner.BirthDate,
	).Scan(&owner.ID)

	if err != nil {
		return err
	}

	for _, petID := range owner.PetIDs {
		_, err := r.repository.Db.Exec(`
			UPDATE pets
			SET owner_id = $1
			WHERE id = $2`,
			owner.ID,
			petID,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *OwnerRepository) Update(owner *Owner) error {
	_, err := r.repository.Db.Exec(`
		UPDATE owners
		SET name = $1,
		    birth_date = $2
		WHERE id = $3`,
		owner.Name,
		owner.BirthDate,
		owner.ID,
	)

	if err != nil {
		return err
	}

	_, err = r.repository.Db.Exec(`
		UPDATE pets
		SET owner_id = NULL
		WHERE owner_id = $1`,
		owner.ID,
	)

	if err != nil {
		return err
	}

	for _, petID := range owner.PetIDs {
		_, err := r.repository.Db.Exec(`
			UPDATE pets
			SET owner_id = $1
			WHERE id = $2`,
			owner.ID,
			petID,
		)

		if err != nil {
			return err
		}
	}

	return nil
}

func (r *OwnerRepository) FindByID(OwnerID int) (owner *Owner, err error) {
	row := r.repository.Db.QueryRow(`
			select
		    (id, name, birth_date, pet_id)
		    from owners
			where id = $1`,
		OwnerID)
	owner = &Owner{}
	err = row.Scan(&owner.ID, &owner.Name, &owner.BirthDate, &owner.PetIDs)

	return owner, err
}

func (r *OwnerRepository) FindAll() (Owners []Owner, err error) {
	rows, err := r.repository.Db.Query(`
			select 
    			o.id, o.name, o.birth_date, array_agg(p.id) as pet_ids
			from owners o
			left join pets p on p.owner_id = o.id
			group by o.id`)
	if err != nil {
		return nil, err
	}

	Owners = []Owner{}
	for rows.Next() {
		owner := &Owner{}
		err := rows.Scan(&owner.ID, &owner.Name, &owner.BirthDate, &owner.PetIDs)
		if err != nil {
			return nil, err
		}
		Owners = append(Owners, *owner)
	}
	return Owners, rows.Err()
}

func (r *OwnerRepository) Delete(ownerID int) (err error) {
	_, err = r.repository.Db.Exec(`delete from owners where id = $1`, ownerID)
	return err
}
