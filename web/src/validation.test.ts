import { describe, expect, test } from "vitest";
import {
  IConfirmPasswordValidation,
  IEmailValidation,
  IPasswordValidation,
  isValid,
  validateConfirmPassword,
  validateEmail,
  validatePassword,
} from "./validation";

// Robust testing saves lives
describe("Test the isValid function", () => {
  test("Success on valid password validation", () => {
    const validation: IPasswordValidation = {
      hasMinLength: true,
      hasLowerCase: true,
      hasUpperCase: true,
      hasNumber: true,
      hasSpecialCharacter: true,
    };
    expect(isValid(validation)).toBe(true);
  });

  test("Failure on invalid password validation", () => {
    const validation: IPasswordValidation = {
      hasMinLength: false,
      hasLowerCase: false,
      hasUpperCase: false,
      hasNumber: false,
      hasSpecialCharacter: false,
    };
    expect(isValid(validation)).toBe(false);
  });

  test("Failure on invalid password validation #2", () => {
    const validation: IPasswordValidation = {
      hasMinLength: true,
      hasLowerCase: false,
      hasUpperCase: false,
      hasNumber: false,
      hasSpecialCharacter: false,
    };
    expect(isValid(validation)).toBe(false);
  });

  test("Failure on invalid password validation #3", () => {
    const validation: IPasswordValidation = {
      hasMinLength: true,
      hasLowerCase: true,
      hasUpperCase: true,
      hasNumber: false,
      hasSpecialCharacter: false,
    };
    expect(isValid(validation)).toBe(false);
  });

  test("Failure on invalid password validation #4", () => {
    const validation: IPasswordValidation = {
      hasMinLength: true,
      hasLowerCase: true,
      hasUpperCase: true,
      hasNumber: true,
      hasSpecialCharacter: false,
    };
    expect(isValid(validation)).toBe(false);
  });

  test("Failure on invalid password validation #5", () => {
    const validation: IPasswordValidation = {
      hasMinLength: false,
      hasLowerCase: true,
      hasUpperCase: true,
      hasNumber: true,
      hasSpecialCharacter: true,
    };
    expect(isValid(validation)).toBe(false);
  });

  test("Success on valid email validation", () => {
    const validation: IEmailValidation = {
      isNotEmpty: true,
      isValidEmail: true,
    };
    expect(isValid(validation)).toBe(true);
  });

  test("Failure on invalid email validation", () => {
    const validation: IEmailValidation = {
      isNotEmpty: false,
      isValidEmail: false,
    };
    expect(isValid(validation)).toBe(false);
  });

  test("Failure on invalid email validation #2", () => {
    const validation: IEmailValidation = {
      isNotEmpty: false,
      isValidEmail: true,
    };
    expect(isValid(validation)).toBe(false);
  });

  test("Failure on invalid email validation #3", () => {
    const validation: IEmailValidation = {
      isNotEmpty: true,
      isValidEmail: false,
    };
    expect(isValid(validation)).toBe(false);
  });

  test("Success on valid confirmPassword validation", () => {
    const validation: IConfirmPasswordValidation = {
      isMatch: true,
    };
    expect(isValid(validation)).toBe(true);
  });

  test("Failure on invalid confirmPassword validation", () => {
    const validation: IConfirmPasswordValidation = {
      isMatch: false,
    };
    expect(isValid(validation)).toBe(false);
  });
});

describe("Test password validation", () => {
  const lower = "password";
  const upper = "PASSWORD";
  const number = "12345";
  const special = "!£$%^&*()_+=-#';.,:`][{}/@";
  const tooShort = "1!pP";
  const valid = "1!pP34j@£nBBsae98";
  const valid2 = "b78mba11*^FES2";
  const empty = "";

  test("Lower true, others false", () => {
    const validation = validatePassword(lower);
    expect(validation.hasLowerCase).toBe(true);
    expect(validation.hasUpperCase).toBe(false);
    expect(validation.hasNumber).toBe(false);
    expect(validation.hasSpecialCharacter).toBe(false);
  });

  test("Upper true, others false", () => {
    const validation = validatePassword(upper);
    expect(validation.hasUpperCase).toBe(true);
    expect(validation.hasLowerCase).toBe(false);
    expect(validation.hasNumber).toBe(false);
    expect(validation.hasSpecialCharacter).toBe(false);
  });

  test("Number true, others false", () => {
    const validation = validatePassword(number);
    expect(validation.hasNumber).toBe(true);
    expect(validation.hasLowerCase).toBe(false);
    expect(validation.hasUpperCase).toBe(false);
    expect(validation.hasSpecialCharacter).toBe(false);
  });

  test("Special true, others false", () => {
    const validation = validatePassword(special);
    expect(validation.hasSpecialCharacter).toBe(true);
    expect(validation.hasLowerCase).toBe(false);
    expect(validation.hasUpperCase).toBe(false);
    expect(validation.hasNumber).toBe(false);
  });

  test("Valid but too short", () => {
    const validation = validatePassword(tooShort);
    expect(validation.hasLowerCase).toBe(true);
    expect(validation.hasUpperCase).toBe(true);
    expect(validation.hasNumber).toBe(true);
    expect(validation.hasSpecialCharacter).toBe(true);
    expect(validation.hasMinLength).toBe(false);
  });

  test("Valid unconditionally", () => {
    const validation = validatePassword(valid);
    expect(validation.hasLowerCase).toBe(true);
    expect(validation.hasUpperCase).toBe(true);
    expect(validation.hasNumber).toBe(true);
    expect(validation.hasSpecialCharacter).toBe(true);
    expect(validation.hasMinLength).toBe(true);
  });

  test("Valid unconditionally #2", () => {
    const validation = validatePassword(valid2);
    expect(validation.hasLowerCase).toBe(true);
    expect(validation.hasUpperCase).toBe(true);
    expect(validation.hasNumber).toBe(true);
    expect(validation.hasSpecialCharacter).toBe(true);
    expect(validation.hasMinLength).toBe(true);
  });

  test("Empty password", () => {
    const validation = validatePassword(empty);
    expect(validation.hasLowerCase).toBe(false);
    expect(validation.hasUpperCase).toBe(false);
    expect(validation.hasNumber).toBe(false);
    expect(validation.hasSpecialCharacter).toBe(false);
    expect(validation.hasMinLength).toBe(false);
  });
});

describe("Test email validation", () => {
  const standard = "user@example.com";
  const alphanumeric = "user123@email.co.uk";
  const localDot = "john.doe@company.org";
  const hostDash = "user_name1234@email-provider.net";
  const subdomain = "info@sub.domain.com";
  const ipInBrackets = "user123@[192.168.1.1]";
  const noAt = "user#domain.com";
  const space = "spaced user@domain.info";
  const space2 = "user@domain with space.com";
  const doubleDots = "double..dots@email.org";
  const doubleDots2 = "user@domain..com";
  const noLocal = "@.com";
  const empty = "";

  test("Standard email format is valid", () => {
    const validation = validateEmail(standard);
    expect(validation.isValidEmail).toBe(true);
  });

  test("Emails can be alphanumeric", () => {
    const validation = validateEmail(alphanumeric);
    expect(validation.isValidEmail).toBe(true);
  });

  test("Emails can contain '.' in username", () => {
    const validation = validateEmail(localDot);
    expect(validation.isValidEmail).toBe(true);
  });

  test("Emails can contain '-' in hostname", () => {
    const validation = validateEmail(hostDash);
    expect(validation.isValidEmail).toBe(true);
  });

  test("Emails can contain a subdomain", () => {
    const validation = validateEmail(subdomain);
    expect(validation.isValidEmail).toBe(true);
  });

  test("Hostname can be an IP address in square brackets", () => {
    const validation = validateEmail(ipInBrackets);
    expect(validation.isValidEmail).toBe(true);
  });

  // TLD validation sounds like a job for the server
  test("No '@' in email", () => {
    const validation = validateEmail(noAt);
    expect(validation.isValidEmail).toBe(false);
  });

  test("Spaces are not allowed in email addresses", () => {
    const validation = validateEmail(space);
    expect(validation.isValidEmail).toBe(false);
  });

  test("Spaces are not allowed in email addresses #2", () => {
    const validation = validateEmail(space2);
    expect(validation.isValidEmail).toBe(false);
  });

  test("Consecutive dots are not allowed in email addresses", () => {
    const validation = validateEmail(doubleDots);
    expect(validation.isValidEmail).toBe(false);
  });

  test("Consecutive dots are not allowed in email addresses #2", () => {
    const validation = validateEmail(doubleDots2);
    expect(validation.isValidEmail).toBe(false);
  });

  test("Email address needs a username", () => {
    const validation = validateEmail(noLocal);
    expect(validation.isValidEmail).toBe(false);
  });

  test("Empty email", () => {
    const validation = validateEmail(empty);
    expect(validation.isNotEmpty).toBe(false);
  });
});

describe("Test confirmPassword validation", () => {
  const password1 = "mypassword";
  const password2 = "drowssapym"; // it's 'mypassword' backwards
  const empty = "";

  test("Valid if passwords match", () => {
    const validation = validateConfirmPassword(password1, password1);
    expect(validation.isMatch).toBe(true);
  });

  test("Valid if passwords match #2", () => {
    const validation = validateConfirmPassword(password2, password2);
    expect(validation.isMatch).toBe(true);
  });

  test("Valid if passwords match #3", () => {
    const validation = validateConfirmPassword(empty, empty);
    expect(validation.isMatch).toBe(true);
  });

  test("Invalid if passwords do not match", () => {
    const validation = validateConfirmPassword(password1, password2);
    expect(validation.isMatch).toBe(false);
  });

  test("Invalid if passwords do not match #2", () => {
    const validation = validateConfirmPassword(password1, empty);
    expect(validation.isMatch).toBe(false);
  });

  test("Invalid if passwords do not match #3", () => {
    const validation = validateConfirmPassword(password2, empty);
    expect(validation.isMatch).toBe(false);
  });
});
