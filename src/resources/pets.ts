// File generated from our OpenAPI spec by Stainless.

import * as Core from 'audioreworkvisions/core';
import { APIResource } from 'audioreworkvisions/resource';
import { isRequestOptions } from 'audioreworkvisions/core';
import * as PetsAPI from 'audioreworkvisions/resources/pets';

export class Pets extends APIResource {
  /**
   * Create a pet
   */
  create(options?: Core.RequestOptions): Core.APIPromise<void> {
    return this._client.post('/pets', { ...options, headers: { Accept: '*/*', ...options?.headers } });
  }

  /**
   * Info for a specific pet
   */
  retrieve(petId: string, options?: Core.RequestOptions): Core.APIPromise<Pet> {
    return this._client.get(`/pets/${petId}`, options);
  }

  /**
   * List all pets
   */
  list(query?: PetListParams, options?: Core.RequestOptions): Core.APIPromise<Pets>;
  list(options?: Core.RequestOptions): Core.APIPromise<Pets>;
  list(
    query: PetListParams | Core.RequestOptions = {},
    options?: Core.RequestOptions,
  ): Core.APIPromise<Pets> {
    if (isRequestOptions(query)) {
      return this.list({}, query);
    }
    return this._client.get('/pets', { query, ...options });
  }
}

export interface Pet {
  id: number;

  name: string;

  tag?: string;
}

export type Pets = Array<Pet>;

export interface PetListParams {
  /**
   * How many items to return at one time (max 100)
   */
  limit?: number;
}

export namespace Pets {
  export import Pet = PetsAPI.Pet;
  export import Pets = PetsAPI.Pets;
  export import PetListParams = PetsAPI.PetListParams;
}
